package client

import (
	"io/ioutil"
	"runtime"
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
		for {
			<-openFile.ClickedCh

			var format string

			switch runtime.GOOS {
			case "linux":
				format = "Audio (*.mid,*.midi) | *.mid *.midi"
			case "windows":
				format = "Audio (*.mid, *.midi)\x00*.mid;*.midi\x00All Files (*.*)\x00*.*\x00\x00"
			case "darwin":
				format = "public.audio"
			}

			f, b, err := dlgs.File("Select file", format, false)
			if err != nil {
				p.log.Warnf("error with file dialog: %+v", err)
				return
			}

			if !b {
				p.log.Info("no file selected")
				return
			}

			p.LoadMusic(f)
		}
	}()

	systray.AddSeparator()
	drums := systray.AddMenuItem("Drums", "Use drumset bindings")
	drums.Check()

	//more instrument options and bindings later

	systray.AddSeparator()
	reload := systray.AddMenuItem("Reload config", "Reload config.toml")

	go func() {
		for {
			<-reload.ClickedCh
			p.log.Info("reloading config")
			p.loadConfig()
		}
	}()

	quit := systray.AddMenuItem("Quit", "Exit TowerPro")

	go func() {
		<-quit.ClickedCh
		p.log.Info("exiting towerpro")
		p.Kill <- nil
	}()
}

func (p *Performer) onExit() {
	p.Kill <- nil
}
