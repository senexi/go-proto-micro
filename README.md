# go-proto-micro

This is a reference setup for microservices written in `go`.

## requirements

Run `scripts/checkRequirements.sh` to check if any requirement is missing.

### go
You need a working `go` development environment with `GOPATH` and `GOBIN` setup.
See the official [Getting started](https://golang.org/doc/install) and
[How to write go code](https://golang.org/doc/code.html#Workspaces)
for details

Once `go` is installed, you can use `scripts/installRequirements.sh`
to get all the required binaries.

### make
You need `make` to run the various goals in the `Makefile`.

Run `make list` to see all targets.

## protobuf support

Install the `protoc` compiler binary from [here](https://github.com/protocolbuffers/protobuf/releases).
On Ubuntu you can also just do `sudo apt install protobuf-compiler`.

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
