package main

import (
	log "github.com/sirupsen/logrus"
)

type application struct {
	config *configuration
	reload chan interface{}
	done   chan interface{}
}

func main() {
	log.Info("starting towerpro...")
	app := &application{
		reload: make(chan interface{}),
		done:   make(chan interface{}),
	}
	//call setup
	//start
	<-app.done
	return
}
