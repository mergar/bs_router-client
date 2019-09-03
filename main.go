package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	configFile	= flag.String("config", "config.json", "Path to config.json")
	body		= flag.String("body", "foobar", "Body of message")
)

func init() {
	flag.Parse()
}

func main() {
	config, err := LoadConfiguration(*configFile)

	if err != nil {
		fmt.Println("config load error")
		os.Exit(1)
	}

	beanstalkSend(config.BeanstalkConfig, *body)
}
