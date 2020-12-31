package client

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/sirupsen/logrus"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/player"
	"gitlab.com/gomidi/rtmididrv"
)

type Platform interface {
	startHooks([]string, *Performer)
	PlayNote()
}

type Performer struct {
	config *configuration
	log    *logrus.Logger
	OS     Platform
	driver *rtmididrv.Driver
	player *player.Player
	Done   chan interface{}
	Kill   chan interface{}
	stop   chan bool
}

func New(d chan interface{}, l *logrus.Logger) *Performer {
	c := NewConfiguration()

	//Create corresponding Os object for platform
	//For development purposes, Windows only right now
	o := &OsWin{
		Log: l,
	}

	drv, err := rtmididrv.New()
	if err != nil {
		l.Fatalf("failed to open rtmididrv: %+v", err)
	}

	return &Performer{
		config: c,
		log:    l,
		OS:     o,
		driver: drv,
		Done:   d,
		stop:   make(chan bool),
	}
}

func (p *Performer) Start() {
	go systray.Run(p.onReady, p.onExit)
	p.loadConfig()
	<-p.Kill
	p.stop <- true
	p.Done <- nil
}

func (p *Performer) Exit() {
	p.driver.Close()
}

func (p *Performer) LoadMusic(fs string) {
	p.log.Infof("file to open: %s", fs)

	pl, err := player.SMF(fs)
	if err != nil {
		p.log.Warnf("error opening player on file: %+v", err)
		return
	}

	p.player = pl
	p.Play()
}

func (p *Performer) Play() {
	outs, err := p.driver.Outs()
	if err != nil {
		p.log.Warnf("error retrieving output channels: %+v", err)
		return
	}

	out := outs[0]

	err = out.Open()
	if err != nil {
		p.log.Warnf("error opening output channel: %+v", err)
		return
	}

	defer out.Close()

	err = printMIDIPorts(p.driver)
	if err != nil {
		p.log.Warnf("error: %+v", err)
	}

	p.player.PlayAll(out, p.stop)
}

func printMIDIPorts(drv midi.Driver) error {
	outs, err := drv.Outs()

	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "MIDI outputs")

	for _, out := range outs {
		fmt.Fprintf(os.Stdout, "[%v] %s\n", out.Number(), out.String())
	}

	return nil
}
