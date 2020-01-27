# Learning Go Lang


## Ressources

**[Golang Language Specs](https://golang.org/ref/spec)** Offical Specs

[Effective Go](https://golang.org/doc/effective_go.html)

[godoc](https://godoc.org/) includes 3 Party Packages

[Presentations](https://goo.gl/Tbz6Xf)

[Tod MacLeod Github/GolangTraining](https://github.com/GoesToEleven/GolangTraining)



## Installation and Setup

Download and Exttract

    $ sudo tar -C /usr/local/ -xzf Downloads go1.13.6.linux-amd64.tar.gz

Add /usr/local/go/bin to environment variable --> Add this line to your /etc/profile for System Wide Installation. or .bashrc

    export PATH=$PATH:/user/local/go/bin

GOBIN for installation


## Go workspace

Folder somewhere
 * bin
 * pkg
 * src
     * github.com
     * username --> get repo with go get
  

## Commands - $ go help

1. go fmt --> formats code
2. go run --> just runs file
3. go buit --> builds
4. go install --> installs in workspace/bin
5. flags
