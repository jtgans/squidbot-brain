package main;

import (
	"flag"
	"log"

	"github.com/jtgans/squidbot-brain/brain"
)

var port = flag.String("port", "", "the port which the squidbot brain server must listen on.")
var configFile = flag.String("config", "config.toml", "the file containing the brain's config.")

func main() {
	flag.Parse()

	if flag.Lookup("port").DefValue == *port {
		log.Fatalf("No port specified to listen on.")
	}

	brainServer := brain.NewServer(*port, *configFile)
	brainServer.Run()
}
