# my5G-core


![GitHub](https://img.shields.io/github/license/my5G/my5G-core?color=blue) 
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/my5G/my5G-core) ![GitHub commit activity](https://img.shields.io/github/commit-activity/y/my5G/my5G-core) 
![GitHub last commit](https://img.shields.io/github/last-commit/my5G/my5G-core)
![GitHub contributors](https://img.shields.io/github/contributors/my5G/my5G-core)

<img width="20%" src="docs/media/img/my5g-logo.png" alt="my5g-core"/>

- [Description](#description)
- [Installation](#installation)
    - [Recommended environment](#recommended-environment)
    - [Requirements](#requirements)
    - [RAN Tester](#ran-tester)
- [Check](#check)
- [More information](#more-information)

----
# Description

Currently, my5G-core is a fork of the free5GC project, with some extensions to facilitate the deployment.

This 5G core is very flexible and can be installed and configured in many different ways. In this document, we present only the most simple deployment, in which all network functions are installed and configure to run inside the same system.

If you want to cite this tool, please use the following information:
```
@misc{???,
    title={???},
    author={???},
    year={2020},
    eprint={???},
    archivePrefix={arXiv},
    primaryClass={cs.NI}
}
```
If you have questions or comments, please email us: [my5G team](mailto:my5G.initiative@gmail.com). 

----
# Installation

## Recommended environment

* CPU: Intel i7 processor
* RAM: 8 GB
* Disk space: 160 GB
* Operating System (OS): Linux Ubuntu 18.04 or 20.04 LTS
* Kernel verstion: > 5.4 (important for UPF only).

The installation can be done directly in the host OS or inside a virtual machine (VM).

## Requirements

Install the necessary packages available in Ubuntu repositories:
```
sudo apt update && sudo apt -y install gcc cmake autoconf build-essential libtool pkg-config libmnl-dev libyaml-dev wget git net-tools mongodb
```

Make sure that MongoDB is running:
```
sudo systemctl start mongodb
```

Install Go (assuming there is no previous version installed):
```
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
sudo tar -C /usr/local -zxvf go1.14.4.linux-amd64.tar.gz
mkdir -p ~/go/{bin,pkg,src}
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin:$GOROOT/bin' >> ~/.bashrc
source ~/.bashrc
```

Install the proper version of Logrus and fatal:
```
go get -u github.com/sirupsen/logrus
go get -u github.com/calee0219/fatal
```

Install the proper version of the Linux kernel module 5G GTP-U:
```
git clone -b v0.1.0 https://github.com/PrinzOwO/gtp5g.git
cd gtp5g
make clean && make && sudo make install
```

Configure the Linux host to offer routing and NAT services. In the following, **<DN_INT>** must be substituted by the name of the interface used for Internet access, e.g., **eth0** or **enp0s3**. The firewall service (i.e., **ufw**) is disabled to assure the communication of the 5G core with the outside networks.
```
sudo sysctl -w net.ipv4.ip_forward=1
sudo iptables -t nat -A POSTROUTING -o <DN_INT> -j MASQUERADE
sudo systemctl stop ufw
```

## 5G core

Download the source code:
```
git clone https://github.com/my5G/my5G-core.git
cd my5G-core
git checkout master
git submodule sync
git submodule update --init --jobs `nproc`
git submodule foreach git checkout master
git submodule foreach git pull --jobs `nproc`
```

Install the dependencies:
```
go mod download
```

Compile the network functions:
```
make all
```

Done! The software is successfully installed.

----
# Check

Make Registration test:
```
./test.sh TestRegistration
```

You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-RegistrationTest-OK.png"/>
</p>

Make GUTI Registration test:
```
./test.sh TestGUTIRegistration
```

You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-GUTIRegistrationTest-OK.png"/>
</p>

Make ServiceRequest test:
```
./test.sh TestServiceRequest
```

You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-ServiceRequestTest-OK.png"/>
</p>

Make XnHandover test:
```
./test.sh TestXnHandover
```

You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-XnHandoverTest-OK.png"/>
</p>

Make Deregistration test:
```
./test.sh TestDeregistration
```

You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-DeregistrationTest-OK.png"/>
</p>

Make PDUSessionReleaseRequest test:
```bash
./test.sh TestPDUSessionReleaseRequest
```

You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-PDUSessionReleaseRequestTest-OK.png"/>
</p>

Make Paging test:
```
./test.sh TestPaging
```

You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-PagingTest-OK.png"/>
</p>


Make N2Handover test:
```
./test.sh TestN2Handover
```
You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-N2HandoverTest-OK.png"/>
</p>

Make Non3GPP test:
```
./test.sh TestNon3GPP
```

Make ULCL test:
```
./test_ulcl.sh -om 3 TestRegistration
```

You should see several INFO messages and the last lines should be similar to the following:
<p align="">
     <img src="docs/media/img/my5G-core-ULCLTest-OK.png"/>
</p>

# More information
   
If my5G-core was successfuly installed, you may find useful take a look in the [USAGE](https://github.com/my5G/my5G-core/blob/develop/USAGE.md) manual that also presents several examples.

Please review the [CONTRIBUTION](https://github.com/my5G/template/blob/main/CONTRIBUTING.md) guide for information on how to get started contributing to the project.

If you found problems during the installation or during the checking, you may find a solution in the [TROUBLESHOOTING](https://github.com/my5G/my5G-core/blob/develop/TROUBLESHOOTING.md) document.
