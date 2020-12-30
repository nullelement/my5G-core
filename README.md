# my5G-core


![GitHub](https://img.shields.io/github/license/LABORA-INF-UFG/my5G-core?color=blue) 
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/LABORA-INF-UFG/my5G-core) ![GitHub commit activity](https://img.shields.io/github/commit-activity/y/LABORA-INF-UFG/my5G-core) 
![GitHub last commit](https://img.shields.io/github/last-commit/LABORA-INF-UFG/my5G-core)
![GitHub contributors](https://img.shields.io/github/contributors/LABORA-INF-UFG/my5G-core)

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
git clone https://github.com/my5G/my5Gcore.git
cd my5Gcore
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



## Recommended Environment
- Software
    - OS: Ubuntu 18.04
    - gcc 7.3.0
    - Go 1.14.4 linux/amd64
    - kernel version 5.0.0-23-generic (MUST for UPF)
    
**Note: Please use Ubuntu 18.04 and kernel version 5.0.0-23-generic** 


You can use `go version` to check your current Go version.
```bash
- Hardware
    - CPU: Intel i5 processor
    - RAM: 4GB
    - Hard drive: 160G
    - NIC card: 1Gbps ethernet card

- Hardware recommended
    - CPU: Intel i7 processor
    - RAM: 8GB
    - Hard drive: 160G
    - NIC card: 10Gbps ethernet card
```


## Installation
### A. Pre-requisite

0. Required kernel version `5.0.0-23-generic`. This request is from the module
   [gtp5g](https://github.com/PrinzOwO/gtp5g) that we has used. Any more details
   please check [here](https://github.com/PrinzOwO/gtp5g)
   ```bash
   # Check kernel version
   $ uname -r
   5.0.0-23-generic
   ```

1. Require go language
    * If another version of Go is installed
        - Please remove the previous Go version
            - ```sudo rm -rf /usr/local/go```
        - Install Go 1.14.4
            ```bash
            wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
            sudo tar -C /usr/local -zxvf go1.14.4.linux-amd64.tar.gz
            ```
    * Clean installation
        - Install Go 1.14.4
             ```bash
            wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
            sudo tar -C /usr/local -zxvf go1.14.4.linux-amd64.tar.gz
            mkdir -p ~/go/{bin,pkg,src}
            echo 'export GOPATH=$HOME/go' >> ~/.bashrc
            echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
            echo 'export PATH=$PATH:$GOPATH/bin:$GOROOT/bin' >> ~/.bashrc
            source ~/.bashrc
            ```

2. Required packages for control plane
    ```bash
    sudo apt -y update
    sudo apt -y install mongodb wget git
    sudo systemctl start mongodb
    ```

3. Required packages for user plane
    ```bash
    sudo apt -y update
    sudo apt -y install git gcc cmake autoconf libtool pkg-config libmnl-dev libyaml-dev
    go get -u github.com/sirupsen/logrus
    ```

4. Network Setting
    ```bash
    sudo sysctl -w net.ipv4.ip_forward=1
    sudo iptables -t nat -A POSTROUTING -o <dn_interface> -j MASQUERADE
    sudo systemctl stop ufw
    ```

### B. Install Control Plane Entities
    
1. Clone free5GC project
    ```bash
    cd ~
    git clone --recursive -b v3.0.3 -j `nproc` https://github.com/free5gc/free5gc.git
    cd free5gc
    ```

    (Optional) If you want to use the nightly version, runs:
    ```bash
    cd ~/free5gc
    git checkout master
    git submodule sync
    git submodule update --init --jobs `nproc`
    git submodule foreach git checkout master
    git submodule foreach git pull --jobs `nproc`
    ```

2. Run the script to install dependent packages
    ```bash
    cd ~/free5gc
    go mod download
    ```
    **In step 2, the folder name should remain free5gc. Please do not modify it or the compilation would fail.**

3. Compile network function services in `free5gc` individually, e.g. AMF (redo this step for each NF), or
    ```bash
    cd ~/free5gc
    go build -o bin/amf -x src/amf/amf.go
    ```
    **To build all network functions in one command**
    ```bash
    ./build.sh
    ```


### C. Install User Plane Function (UPF)
    
1. Please check Linux kernel version if it is `5.0.0-23-generic`
    ```bash
    uname -r
    ```


    Get Linux kernel module 5G GTP-U
    ```bash
    git clone -b v0.1.0 https://github.com/PrinzOwO/gtp5g.git
    cd gtp5g
    make
    sudo make install
    ```
2. Build from sources
    ```bash
    cd ~/free5gc/src/upf
    mkdir build
    cd build
    cmake ..
    make -j`nproc`
    ```
    
**Note: UPF's config is located at** `free5gc/src/upf/build/config/upfcfg.yaml
   `

## Run

### A. Run Core Network 
Option 1. Run network function service individually, e.g. AMF (redo this for each NF), or
```bash
cd ~/free5gc
./bin/amf
```

**Note: For N3IWF needs specific configuration in section B** 

Option 2. Run whole core network with command
```
./run.sh
```

### B. Run N3IWF (Individually)
To run N3IWF, make sure the machine is equipped with three network interfaces. (one is for connecting AMF, another is for connecting UPF, the other is for IKE daemon)

We need to configure each interface with a suitable IP address.

We have to create an interface for IPSec traffic:
```bash
# replace <...> to suitable value
sudo ip link add ipsec0 type vti local <IKEBindAddress> remote 0.0.0.0 key <IPSecInterfaceMark>
```
Assign an address to this interface, then bring it up:
```bash
# replace <...> to suitable value
sudo ip address add <IPSecInterfaceAddress/CIDRPrefix> dev ipsec0
sudo ip link set dev ipsec0 up
```

Run N3IWF (root privilege is required):
```bash
cd ~/free5gc/
sudo ./bin/n3iwf
```

### C. Run all in one with outside RAN

Reference to [sample config](./sample/ran_attach_config) if need to connect the
outside RAN with all in one free5GC core network.

### D. Deploy with container

Reference to [free5gc-compose](https://github.com/free5gc/free5gc-compose/) as
the sample for container deployment.

## Test
Start Wireshark to capture any interface with `pfcp||icmp||gtp` filter and run the tests below to simulate the procedures:
```bash
cd ~/free5gc
chmod +x ./test.sh
```
a. TestRegistration
```bash
(In directory: ~/free5gc)
./test.sh TestRegistration
```
b. TestServiceRequest
```bash
./test.sh TestServiceRequest
```
c. TestXnHandover
```bash
./test.sh TestXnHandover
```
d. TestDeregistration
```bash
./test.sh TestDeregistration
```
e. TestPDUSessionReleaseRequest
```bash
./test.sh TestPDUSessionReleaseRequest
```

f. TestPaging
```!
./test.sh TestPaging
```

g. TestN2Handover
```!
./test.sh TestN2Handover
```

h. TestNon3GPP
```bash
./test.sh TestNon3GPP
```

i. TestULCL
```bash
./test_ulcl.sh -om 3 TestRegistration
```

**For more details, you can reference to our [wiki](https://github.com/free5gc/free5gc/wiki)**

## Release Note
Detailed changes for each release are documented in the [release notes](https://github.com/free5gc/free5gc/releases).
