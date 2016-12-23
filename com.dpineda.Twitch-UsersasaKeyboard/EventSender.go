package main

import (
	"github.com/micmonay/keybd_event"
	"strings"
	"time"
	"log"
	"fmt"
)

var kb keybd_event.KeyBonding

func initializeKeySender() error {
	kbtmp, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	kb = kbtmp

	time.Sleep(2 * time.Second)

	return nil

}

func sendEvent(word string, event string) {
	appendTextToLabel("doing event... " + word)

	if strings.HasPrefix(event, "PRESSKEY") {
		pressKeyEvent(event)
	} else if strings.HasPrefix(event, "WRITETEXT") {
		writeTextEvent(event)
	}

}

func pressKeyEvent(event string) {
	splitted := strings.Split(event, "#")
	keyToPress := splitted[len(splitted)-1]

	switch keyToPress {
	// Characters
	case "VK_A":
		kb.SetKeys(keybd_event.VK_A)
	case "VK_B":
		kb.SetKeys(keybd_event.VK_B)
	case "VK_C":
		kb.SetKeys(keybd_event.VK_C)
	case "VK_D":
		kb.SetKeys(keybd_event.VK_D)
	case "VK_E":
		kb.SetKeys(keybd_event.VK_E)
	case "VK_F":
		kb.SetKeys(keybd_event.VK_F)
	case "VK_G":
		kb.SetKeys(keybd_event.VK_G)
	case "VK_H":
		kb.SetKeys(keybd_event.VK_H)
	case "VK_I":
		kb.SetKeys(keybd_event.VK_I)
	case "VK_J":
		kb.SetKeys(keybd_event.VK_J)
	case "VK_K":
		kb.SetKeys(keybd_event.VK_K)
	case "VK_L":
		kb.SetKeys(keybd_event.VK_L)
	case "VK_M":
		kb.SetKeys(keybd_event.VK_M)
	case "VK_N":
		kb.SetKeys(keybd_event.VK_N)
	case "VK_O":
		kb.SetKeys(keybd_event.VK_O)
	case "VK_P":
		kb.SetKeys(keybd_event.VK_P)
	case "VK_Q":
		kb.SetKeys(keybd_event.VK_Q)
	case "VK_R":
		kb.SetKeys(keybd_event.VK_R)
	case "VK_S":
		kb.SetKeys(keybd_event.VK_S)
	case "VK_T":
		kb.SetKeys(keybd_event.VK_T)
	case "VK_U":
		kb.SetKeys(keybd_event.VK_U)
	case "VK_V":
		kb.SetKeys(keybd_event.VK_V)
	case "VK_W":
		kb.SetKeys(keybd_event.VK_W)
	case "VK_X":
		kb.SetKeys(keybd_event.VK_X)
	case "VK_Y":
		kb.SetKeys(keybd_event.VK_Y)
	case "VK_Z":
		kb.SetKeys(keybd_event.VK_Z)
	// Numerics
	case "VK_0":
		kb.SetKeys(keybd_event.VK_0)
	case "VK_1":
		kb.SetKeys(keybd_event.VK_1)
	case "VK_2":
		kb.SetKeys(keybd_event.VK_2)
	case "VK_3":
		kb.SetKeys(keybd_event.VK_3)
	case "VK_4":
		kb.SetKeys(keybd_event.VK_4)
	case "VK_5":
		kb.SetKeys(keybd_event.VK_5)
	case "VK_6":
		kb.SetKeys(keybd_event.VK_6)
	case "VK_7":
		kb.SetKeys(keybd_event.VK_7)
	case "VK_8":
		kb.SetKeys(keybd_event.VK_8)
	case "VK_9":
		kb.SetKeys(keybd_event.VK_9)
	// Function keys
	case "VK_F1":
		kb.SetKeys(keybd_event.VK_F1)
	case "VK_F2":
		kb.SetKeys(keybd_event.VK_F2)
	case "VK_F3":
		kb.SetKeys(keybd_event.VK_F3)
	case "VK_F4":
		kb.SetKeys(keybd_event.VK_F4)
	case "VK_F5":
		kb.SetKeys(keybd_event.VK_F5)
	case "VK_F6":
		kb.SetKeys(keybd_event.VK_F6)
	case "VK_F7":
		kb.SetKeys(keybd_event.VK_F7)
	case "VK_F8":
		kb.SetKeys(keybd_event.VK_F8)
	case "VK_F9":
		kb.SetKeys(keybd_event.VK_F9)
	case "VK_F10":
		kb.SetKeys(keybd_event.VK_F10)
	case "VK_F11":
		kb.SetKeys(keybd_event.VK_F11)
	case "VK_F12":
		kb.SetKeys(keybd_event.VK_F12)
	case "VK_RESERVED":
		kb.SetKeys(keybd_event.VK_RESERVED)
	case "VK_ESC":
		kb.SetKeys(keybd_event.VK_ESC)
	case "VK_MINUS":
		kb.SetKeys(keybd_event.VK_MINUS)
	case "VK_EQUAL":
		kb.SetKeys(keybd_event.VK_EQUAL)
	case "VK_BACKSPACE":
		kb.SetKeys(keybd_event.VK_BACKSPACE)
	case "VK_TAB":
		kb.SetKeys(keybd_event.VK_TAB)
	case "VK_LEFTBRACE":
		kb.SetKeys(keybd_event.VK_LEFTBRACE)
	case "VK_RIGHTBRACE":
		kb.SetKeys(keybd_event.VK_RIGHTBRACE)
	case "VK_ENTER":
		kb.SetKeys(keybd_event.VK_ENTER)
	case "VK_SEMICOLON":
		kb.SetKeys(keybd_event.VK_SEMICOLON)
	case "VK_APOSTROPHE":
		kb.SetKeys(keybd_event.VK_APOSTROPHE)
	case "VK_GRAVE":
		kb.SetKeys(keybd_event.VK_GRAVE)
	case "VK_BACKSLASH":
		kb.SetKeys(keybd_event.VK_BACKSLASH)
	case "VK_COMMA":
		kb.SetKeys(keybd_event.VK_COMMA)
	case "VK_DOT":
		kb.SetKeys(keybd_event.VK_DOT)
	case "VK_SLASH":
		kb.SetKeys(keybd_event.VK_SLASH)
	case "VK_KPASTERISK":
		kb.SetKeys(keybd_event.VK_KPASTERISK)
	case "VK_SPACE":
		kb.SetKeys(keybd_event.VK_SPACE)
	case "VK_CAPSLOCK":
		kb.SetKeys(keybd_event.VK_CAPSLOCK)
	case "VK_NUMLOCK":
		kb.SetKeys(keybd_event.VK_NUMLOCK)
	case "VK_SCROLLLOCK":
		kb.SetKeys(keybd_event.VK_SCROLLLOCK)
	case "VK_KPMINUS":
		kb.SetKeys(keybd_event.VK_KPMINUS)
	case "VK_KPPLUS":
		kb.SetKeys(keybd_event.VK_KPPLUS)
	case "VK_KPDOT":
		kb.SetKeys(keybd_event.VK_KPDOT)
	case "VK_KP0":
		kb.SetKeys(keybd_event.VK_KP0)
	case "VK_KP1":
		kb.SetKeys(keybd_event.VK_KP1)
	case "VK_KP2":
		kb.SetKeys(keybd_event.VK_KP2)
	case "VK_KP3":
		kb.SetKeys(keybd_event.VK_KP3)
	case "VK_KP4":
		kb.SetKeys(keybd_event.VK_KP4)
	case "VK_KP5":
		kb.SetKeys(keybd_event.VK_KP5)
	case "VK_KP6":
		kb.SetKeys(keybd_event.VK_KP6)
	case "VK_KP7":
		kb.SetKeys(keybd_event.VK_KP7)
	case "VK_KP8":
		kb.SetKeys(keybd_event.VK_KP8)
	case "VK_KP9":
		kb.SetKeys(keybd_event.VK_KP9)
	case "VK_UP":
		kb.SetKeys(keybd_event.VK_UP)
	case "VK_DOWN":
		kb.SetKeys(keybd_event.VK_DOWN)
	case "VK_LEFT":
		kb.SetKeys(keybd_event.VK_LEFT)
	case "VK_RIGHT":
		kb.SetKeys(keybd_event.VK_RIGHT)
	default:
		fmt.Println("default")
	}

	if strings.Contains(event, "SHIFT") {
		kb.HasSHIFT(true)
	}
	if strings.Contains(event, "CTRL") {
		kb.HasCTRL(true)
	}
	if strings.Contains(event, "ALT") {
		kb.HasALT(true)
	}

	kb.Launching()
}

func writeTextEvent(event string) {
	splitted := strings.Split(event, "#")
	textToWrite := splitted[len(splitted)-1]

	for i:=0 ; i<len(textToWrite); i++ {
		newEvent := "PRESSKEY"
		character := string(textToWrite[i])
		if character == strings.ToUpper(character) {
			// is mayus
			newEvent = newEvent + "#SHIFT"
		}
		newEvent = newEvent + "#VK_" + strings.ToUpper(character)
		if character == " " {
			newEvent = "PRESSKEY#VK_SPACE"
		}
		pressKeyEvent(newEvent)
	}

}