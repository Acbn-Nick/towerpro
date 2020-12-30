package client

import (
	"io/ioutil"
	"time"

	"github.com/getlantern/systray"
)

func (p *Performer) onReady() {
	ico, err := ioutil.ReadFile("./assets/icon.ico")
	if err != nil {
		p.log.Fatal("error loading systray icon: " + err.Error())
	}

	time.Sleep(499 * time.Millisecond) // Add 500ms delay to fix issue with systray.AddMenuItem() in goroutines on Windows.

	systray.SetIcon(ico)
	systray.SetTitle("TowerPro")
	systray.SetTooltip("TowerPro Performer")

	drums := systray.AddMenuItem("Drums", "Use drumset bindings")
	drums.Check()

	//more instrument bindings later

	systray.AddSeparator()
	reload := systray.AddMenuItem("Reload config", "Reload config.toml")

	go func() {
		<-reload.ClickedCh
		p.log.Info("reloading config")
		p.loadConfig()
	}()

	quit := systray.AddMenuItem("Quit", "Exit TowerPro")

	go func() {
		<-quit.ClickedCh
		p.log.Info("exiting towerpro")
		p.Done <- nil
	}()
}

func (p *Performer) onExit() {
	p.Done <- nil
}
