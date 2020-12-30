package main

import (
	"os"

	"github.com/sirupsen/logrus"

	client "github.com/Acbn-Nick/towerpro/internal"
)

func main() {
	log := logrus.New()

	logFile, err := os.OpenFile("log.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error creating log.txt: %+v", err)
	}

	defer logFile.Close()

	log.SetReportCaller(true)
	log.Out = logFile

	done := make(chan interface{})

	log.Info("starting towerpro...")

	app := client.New(done, log)
	app.Start()
	<-app.Done
	return
}
