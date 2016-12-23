package main

import (
	"github.com/thoj/go-ircevent"
	"time"
)

var lastEvent int64 = 0
var aggregatedEvents map[string]int64

func initializeAggregatedEvents() map[string]int64 {
	mapa := make(map[string]int64)
	for i:=0;i<len(Cfg.Keysbinding);i++ {
		mapa[Cfg.Keysbinding[i].Word] = 0
	}
	return mapa
}

func ProcessMessages(events chan *irc.Event) error {
	err := initializeKeySender()
	if err != nil {
		setRunning(false)
		appendTextToLabel(err.Error())
		return err
	}
	aggregatedEvents = initializeAggregatedEvents()

	setRunning(true)

	for {
		event := <- events

		aggregateEvent(event)

		currentTime := time.Now().UnixNano() / 1000000
		if currentTime > lastEvent + Cfg.EventFrequencyInMs {
			// do event
			doEvent()
			// clear results
			aggregatedEvents = initializeAggregatedEvents()
			// reset the last event
			lastEvent = time.Now().UnixNano() / 1000000
		}
	}
	return nil
}

func aggregateEvent(event *irc.Event) {
	for i:=0;i<len(Cfg.Keysbinding);i++ {
		if Cfg.Keysbinding[i].Word == event.Message() {
			aggregatedEvents[event.Message()] = aggregatedEvents[event.Message()] + 1
			return
		}
	}
}

func doEvent() {
	bestKey := ""
	var bestValue int64 = 0
	for k, v := range aggregatedEvents {
		if v > bestValue {
			bestKey = k
			bestValue = v
		}
	}

	for i:=0;i<len(Cfg.Keysbinding);i++ {
		if Cfg.Keysbinding[i].Word == bestKey {
			sendEvent(Cfg.Keysbinding[i].Word, Cfg.Keysbinding[i].Event)
			return
		}
	}
}