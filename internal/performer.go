package client

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/sirupsen/logrus"
)

type Platform interface {
	startHooks([]string)
	Play()
}

type Performer struct {
	config *configuration
	log    *logrus.Logger
	music  *os.File
	OS     Platform
	Done   chan interface{}
}

func New(d chan interface{}, l *logrus.Logger) *Performer {
	c := NewConfiguration()

	//Create corresponding Os object for platform
	//For development purposes, Windows only right now
	o := &OsWin{
		Log: l,
	}

	return &Performer{
		config: c,
		log:    l,
		OS:     o,
		Done:   d,
	}
}

func (p *Performer) Start() {
	go systray.Run(p.onReady, p.onExit)
	p.loadConfig()
	<-p.Done
}

func (p *Performer) LoadMusic() {

}

func (p *Performer) Play() {}
