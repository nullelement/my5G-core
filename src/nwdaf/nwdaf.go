package main

import (
	"free5gc/src/app"
	"free5gc/src/nwdaf/logger"
	nwdaf_service "free5gc/src/nwdaf/service"
	"free5gc/src/nwdaf/version"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var NWDAF = &nwdaf_service.NWDAF{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "nwdaf"
	appLog.Infoln(app.Name)
	appLog.Infoln("NWDAF version: ", version.GetVersion())
	app.Usage = "-free5gccfg common configuration file -nwdafcfg nwdaf configuration file"
	app.Action = action
	app.Flags = NWDAF.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		logger.AppLog.Warnf("Error args: %v", err)
	}
}

func action(c *cli.Context) {
	app.AppInitializeWillInitialize(c.String("free5gccfg"))
	NWDAF.Initialize(c)
	NWDAF.Start()
}
