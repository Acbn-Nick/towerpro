package client

import (
	"github.com/getlantern/systray"
	"github.com/sirupsen/logrus"
)

type Performer struct {
	config *configuration
	log    *logrus.Logger
	Done   chan interface{}
}

func New(d chan interface{}, l *logrus.Logger) *Performer {
	c := NewConfiguration()

	return &Performer{
		config: c,
		log:    l,
		Done:   d,
	}
}

func (p *Performer) Start() {
	go systray.Run(p.onReady, p.onExit)
	p.loadConfig()
	<-p.Done
}
