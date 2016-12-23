package main

import (
	"github.com/thoj/go-ircevent"
	"crypto/tls"
)

var irccon *irc.Connection

func ReadFromIRC(server string, port string, username string, password string, channel string, chanmessages chan *irc.Event) error  {
	irccon = irc.IRC(username, username)
	irccon.VerboseCallbackHandler = false
	irccon.Debug = Cfg.Debug
	irccon.Password = password
	if Cfg.Usetls {
		irccon.UseTLS = true
		irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		irccon.UseTLS = false
	}
	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join("#"+channel)
	})
	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		chanmessages <- e
	})
	err := irccon.Connect(server + ":" + port)
	if err != nil {
		appendTextToLabel("Err %s" + err.Error())
		return err
	}
	appendTextToLabel("Connected to channel")
	return nil
}

func closeConnection() {
	irccon.Disconnect()
	irccon.Quit()
}