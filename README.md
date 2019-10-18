# go-proto-micro

This is a reference setup for microservices written in `go`.

## befor you start
Before you start, make sure you have a working go environment, the `make` command and Google's
`protoc` compiler installed.

### go
You need a working `go` development environment with the environment variables
 `$GOPATH` and `$GOBIN` set up.
See the official [Getting started](https://golang.org/doc/install) and
[How to write go code](https://golang.org/doc/code.html#Workspaces)
for details

### protoc
Install the `protoc` compiler binary from [here](https://github.com/protocolbuffers/protobuf/releases).
On Ubuntu you can also just do `sudo apt install protobuf-compiler`.


## quick start

```
git clone https://github.com/senexi/go-proto-micro.git
cd go-proto-micro
chmod +x scripts/*
make install-requirements
make generate &&
go get &&
make run
```

If something goes wrong, run  `make check-requirements`
to see if any of the required dependencies are missing.


## in detail

### binaries for code generation
Once `go` is installed, you can use `make install-requirements`
to get all the required `go` binaries that are needed for code generation. You can check
the script `scripts/installRequirements` to see what they are.

### make
You need `make` to run the various goals in the `Makefile`.

Run `make list` to see all targets.

## protobuf support

All required protobuf tools must be installed. We are using `gogo/protobuf` extension `gogoproto.moretags`
to add additional ORM tags to the generated models.
### grpc support
### clients
- java
- python
### grpc-gateway
### swagger support


## database and ORM support



## metrics

## logger
