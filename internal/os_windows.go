package client

import (
	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
)

type OsWin struct {
	Log *logrus.Logger
}

func (o *OsWin) startHooks(toggle []string, p *Performer) {
	o.Log.Info("starting hooks")
	hook.Register(hook.KeyUp, toggle, func(e hook.Event) {
		p.Play()
	})
}

func (o *OsWin) PlayNote() {

}
