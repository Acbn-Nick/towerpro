package client

import (
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/dlgs"
	"github.com/getlantern/systray"
)

func (p *Performer) onReady() {
	ico, err := ioutil.ReadFile("./assets/icon.ico")
	if err != nil {
		p.log.Fatal("error loading systray icon: " + err.Error())
	}

	time.Sleep(500 * time.Millisecond) // Add 500ms delay to fix issue with systray.AddMenuItem() in goroutines on Windows.

	systray.SetIcon(ico)
	systray.SetTitle("TowerPro")
	systray.SetTooltip("TowerPro Performer")

	openFile := systray.AddMenuItem("Open MIDI file...", "Select a MIDI file to open")

	go func() {
		<-openFile.ClickedCh

		f, b, err := dlgs.File("Select file", "Audio (*.mid, *.midi)", false)
		if err != nil {
			p.log.Warnf("error with file dialog: %+v", err)
			return
		}

		if !b {
			p.log.Info("no file selected")
			return
		}

		if !strings.HasSuffix(strings.ToUpper(f), ".MID") || !strings.HasSuffix(strings.ToUpper(f), ".MIDI") {
			p.log.Warnf("invalid file format on file: %s", f)
			return
		}

		file, err := os.Open(f)
		if err != nil {
			p.log.Warnf("error retrieving file: %+v", err)
		}

		p.music = file
	}()

	systray.AddSeparator()
	drums := systray.AddMenuItem("Drums", "Use drumset bindings")
	drums.Check()

	//more instrument options and bindings later

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
