package client

import (
	"os"

	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
)

type OsWin struct {
	Log *logrus.Logger
}

func (o *OsWin) startHooks(toggle []string) {
	o.Log.Info("starting hooks")
	hook.Register(hook.KeyUp, toggle, func(e hook.Event) {
		o.Play()
	})
}

func (o *OsWin) FileDialog() *os.File {
	//placeholder for now
	file, err := os.Open("C:/Users/Tusks/Music/RocketMan_drums.mid")
	if err != nil {
		o.Log.Fatal("failed opening .mid file")
		return &os.File{}
	}

	return file
}

func (o *OsWin) Play() {
	//performs file
}
