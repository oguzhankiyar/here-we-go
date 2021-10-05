# Installation

## Ubuntu
1. Update and upgrade apt
```bash
sudo apt-get update
sudo apt-get -y upgrade
```
2. Install git
```bash
sudo apt-get install git
```
3. Fetch, unzip and move go sources (please check the latest version)
```bash
wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz
sudo tar -xvf go1.16.4.linux-amd64.tar.gz
sudo mv go /usr/local
```
4. Set the environment variables
```bash
vi ~/.bashrc
export GOROOT=/usr/local/go
export GOPATH=$HOME/work
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
5. Check the installation
```bash
go version
```

## Windows
1. Install git from this address
```
https://git-scm.com/download/win
```
2. Install golang from this address
```
https://git-scm.com/download/win
```
3. Set the environment variables
```cmd
setx GOROOT=C:\Go
setx GOPATH=%USERPROFILE%\go
setx PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin
```
4. Check the installation
```cmd
go version
```


## MacOS
1. Update brew
```bash
brew update
```
2. Install git
```bash
brew install git
```
3. Install golang
```bash
brew install golang
```
4. Set the environment variables
```bash
vi ~/.bashrc
export GOROOT=/usr/local/go
export GOPATH=$HOME/work
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
5. Check the installation
```bash
go version
```