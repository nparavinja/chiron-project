package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	configFilename := "config.yaml"
	configFile, err := ioutil.ReadFile(configFilename)
	if err != nil {
		log.Println("Cannot read configuration - aborting.", err)
		os.Exit(1)
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Println("Cannot parse configuration - aborting.", err)
		os.Exit(1)
	}
	server, err := StartServer(config)
	if err != nil {
		log.Println("Cannot start server - aborting.", err)
		os.Exit(1)
	}

	go func() {
		err = http.ListenAndServe(":8000", server.router)
		if err != nil {
			log.Println("Cannot start http.ListenAndServe - aborting.", err)
			os.Exit(1)
		}
	}()

	log.Println("Server is up and running!")
	select {}
}
