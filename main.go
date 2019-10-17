package main

import (
	"github.com/senexi/go-proto-micro/cmd"
	log "github.com/sirupsen/logrus"
)

var VERSION string
var BUILD string
var NAME string

func main() {
	log.WithFields(log.Fields{"app": NAME, "version": VERSION, "build": BUILD}).Info("starting")
	cmd.Execute()
}
