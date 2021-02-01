package service

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"free5gc/lib/MongoDBLibrary"
	"free5gc/lib/http2_util"
	"free5gc/lib/logger_util"
	"free5gc/lib/path_util"
	"free5gc/src/app"
	"free5gc/src/nwdaf/consumer"
	nwdaf_context "free5gc/src/nwdaf/context"
	"free5gc/src/nwdaf/datarepository"
	"free5gc/src/nwdaf/factory"
	"free5gc/src/nwdaf/logger"
	"free5gc/src/nwdaf/util"
)

type NWDAF struct{}

type (
	// Config information.
	Config struct {
		nwdafcfg string
	}
)

var config Config

var nwdafCLi = []cli.Flag{
	cli.StringFlag{
		Name:  "free5gccfg",
		Usage: "common config file",
	},
	cli.StringFlag{
		Name:  "nwdafcfg",
		Usage: "config file",
	},
}

var initLog *logrus.Entry

func init() {
	initLog = logger.InitLog
}

func (*NWDAF) GetCliCmd() (flags []cli.Flag) {
	return nwdafCLi
}

func (*NWDAF) Initialize(c *cli.Context) {

	config = Config{
		nwdafcfg: c.String("nwdafcfg"),
	}

	if config.nwdafcfg != "" {
		factory.InitConfigFactory(config.nwdafcfg)
	} else {
		DefaultNwdafConfigPath := path_util.Gofree5gcPath("free5gc/config/nwdafcfg.conf")
		factory.InitConfigFactory(DefaultNwdafConfigPath)
	}

	if app.ContextSelf().Logger.NWDAF.DebugLevel != "" {
		level, err := logrus.ParseLevel(app.ContextSelf().Logger.NWDAF.DebugLevel)
		if err != nil {
			initLog.Warnf("Log level [%s] is not valid, set to [info] level", app.ContextSelf().Logger.NWDAF.DebugLevel)
			logger.SetLogLevel(logrus.InfoLevel)
		} else {
			logger.SetLogLevel(level)
			initLog.Infof("Log level is set to [%s] level", level)
		}
	} else {
		initLog.Infoln("Log level is default set to [info] level")
		logger.SetLogLevel(logrus.InfoLevel)
	}

	logger.SetReportCaller(app.ContextSelf().Logger.NWDAF.ReportCaller)

}

func (nwdaf *NWDAF) FilterCli(c *cli.Context) (args []string) {
	for _, flag := range nwdaf.GetCliCmd() {
		name := flag.GetName()
		value := fmt.Sprint(c.Generic(name))
		if value == "" {
			continue
		}

		args = append(args, "--"+name, value)
	}
	return args
}

func (nwdaf *NWDAF) Start() {
	// get config file info
	config := factory.NwdafConfig
	mongodb := config.Configuration.Mongodb

	initLog.Infof("NWDAF Config Info: Version[%s] Description[%s]", config.Info.Version, config.Info.Description)

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)

	initLog.Infoln("Server started")

	router := logger_util.NewGinWithLogrus(logger.GinLog)

	datarepository.AddService(router)

	nwdafLogPath := util.NwdafLogPath
	nwdafPemPath := util.NwdafPemPath
	nwdafKeyPath := util.NwdafKeyPath

	self := nwdaf_context.NWDAF_Self()
	util.InitNwdafContext(self)

	addr := fmt.Sprintf("%s:%d", self.BindingIPv4, self.SBIPort)
	profile := consumer.BuildNFInstance(self)
	var newNrfUri string
	var err error
	newNrfUri, self.NfId, err = consumer.SendRegisterNFInstance(self.NrfUri, profile.NfInstanceId, profile)
	if err == nil {
		self.NrfUri = newNrfUri
	} else {
		initLog.Errorf("Send Register NFInstance Error[%s]", err.Error())
	}

	server, err := http2_util.NewServer(addr, nwdafLogPath, router)
	if server == nil {
		initLog.Errorf("Initialize HTTP server failed: %+v", err)
		return
	}

	if err != nil {
		initLog.Warnf("Initialize HTTP server: %+v", err)
	}

	serverScheme := factory.NwdafConfig.Configuration.Sbi.Scheme
	if serverScheme == "http" {
		err = server.ListenAndServe()
	} else if serverScheme == "https" {
		err = server.ListenAndServeTLS(nwdafPemPath, nwdafKeyPath)
	}

	if err != nil {
		initLog.Fatalf("HTTP server setup failed: %+v", err)
	}
}

func (nwdaf *NWDAF) Exec(c *cli.Context) error {

	//NWDAF.Initialize(cfgPath, c)

	initLog.Traceln("args:", c.String("nwdafcfg"))
	args := nwdaf.FilterCli(c)
	initLog.Traceln("filter: ", args)
	command := exec.Command("./nwdaf", args...)

	nwdaf.Initialize(c)

	var stdout io.ReadCloser
	if readCloser, err := command.StdoutPipe(); err != nil {
		initLog.Fatalln(err)
	} else {
		stdout = readCloser
	}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		in := bufio.NewScanner(stdout)
		for in.Scan() {
			fmt.Println(in.Text())
		}
		wg.Done()
	}()

	var stderr io.ReadCloser
	if readCloser, err := command.StderrPipe(); err != nil {
		initLog.Fatalln(err)
	} else {
		stderr = readCloser
	}
	go func() {
		in := bufio.NewScanner(stderr)
		for in.Scan() {
			fmt.Println(in.Text())
		}
		wg.Done()
	}()

	var err error
	go func() {
		if errormessage := command.Start(); err != nil {
			fmt.Println("command.Start Fails!")
			err = errormessage
		}
		wg.Done()
	}()

	wg.Wait()
	return err
}
