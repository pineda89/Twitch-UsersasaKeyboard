package main

import (
	"github.com/thoj/go-ircevent"
	"github.com/andlabs/ui"
	"strings"
	"os"
)

var configfile string = "../application.yml"

var TextLabel *ui.Label
var RunningLabel *ui.Label
var running bool

var windowname string = "ircbot"

func main() {
	if len(os.Args) > 1 {
		configfile = os.Args[1]
	}

	err := LoadConfig(configfile)
	if err != nil {
		appendTextToLabel(err.Error())
		setRunning(false)
		return
	}

	err = ui.Main(func() {
		button := ui.NewButton("Start")
		RunningLabel = ui.NewLabel("")
		setRunning(false)
		TextLabel = ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(RunningLabel, false)
		box.Append(button, false)
		box.Append(TextLabel, false)
		window := ui.NewWindow(windowname, 300, 200, false)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			if !running {
				setRunning(true)
				go doProcess()
			}
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func doProcess() {
	chanmessages := make(chan *irc.Event, 10000)
	err := ReadFromIRC(Cfg.Server, Cfg.Port, Cfg.User, Cfg.Password, Cfg.Channel, chanmessages)
	if err != nil {
		appendTextToLabel(err.Error())
		setRunning(false)
		return
	}
	err = ProcessMessages(chanmessages)
	if err != nil {
		appendTextToLabel(err.Error())
		setRunning(false)
		closeConnection()
	}
}

func appendTextToLabel(text string) {
	splitted := strings.Split(TextLabel.Text(), "\n")
	if len(splitted) > 10 {
		TextLabel.SetText("")
	}

	TextLabel.SetText(TextLabel.Text() + "\n" + text)
}

func setRunning(running2 bool) {
	var txt string
	if running2 {
		txt = "true"
	} else {
		txt = "false"
	}
	running = running2
	RunningLabel.SetText("running: " + txt)
}