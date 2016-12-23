#Twitch - Users as a Keyboard
This implementation allows you to control a computer using an IRC channel as a "remote control". It works on mac, linux and windows

![Alt text](/doc/first_view_UI.png?raw=true "First view")

##Motivation
The purpose of this project is simple. Allow a group of users from a chat (such as twitch) to control a video game on a single computer, which should be broadcasting their screen.

##Why golang?
why not?

##Architecture
This application must be listening to an IRC channel, and at the same time streaming the content.
The users will write the events, and with this they get the video to execute what they want

![Alt text](/doc/Componentes.png?raw=true "Componentes")

##Internal design
The design of the application is relatively simple.

The IRC component handles the connection to the IRC server, as well as the channel, and is listening to the "new message" event. When one of these events occurs, it adds them to the channel.

The Processor component is responsible for listening to the previous channel. It also controls the times, performs the event aggregates, and notifies the next component to generate the final event.

The EventSender component expects Processor to report a new event. When this occurs, it performs the event action against the operating system API

![Alt text](/doc/InternalDesign.png?raw=true "Internal Design")

##Preconditions to use
* Golang installed

##Libraries useds
* github.com/thoj/go-ircevent : IRC Client library
* github.com/andlabs/ui : Platform-native GUI library
* gopkg.in/yaml.v2 : YAML support
* github.com/micmonay/keybd_event : For simulate key press

##How to use

To compile, simply perform a "go build"

To start the executable, just an application.yml file as the example in the same path as the executable (./binary). Otherwise, it is injected by parameter (./binary /my/path/my.yml)

###Configuration

Attribute  | Description | Example (twitch.tv)
------------- | ------------- | -------------
server | Irc server | irc.chat.twitch.tv
port | Irc server port | 443
user | Irc username | humiltat (my twitch username)
password | Irc auth password | oauth:XXXk1XXXpbetXXX1iiedXXX6k2XX7 (oauth, can get on https://twitchapps.com/tmi/)
channel | Irc channel to read data | humiltat (twitch channel where read, in this case my username)
eventfrequencyinms | Execute an event each X miliseconds | 500
usetls | Flag to indicates if uses secure connection to IRC | true
debug | Flag to indicates if print debug traces | false
keybinding | Contains a list of pairs (word-event). When any user prints the word, the event will be executed (see key list) | word: A && event: PRESSKEY#VK_A


###Configuration example

```yaml
server: irc.chat.twitch.tv
port: 443
user: humiltat
password: oauth:vcyXXXXXXXXXXXXXXXXXXXXXXx7
channel: humiltat
eventfrequencyinms: 500
usetls: true
debug: false
keysbinding:
  - word: UP
    event: PRESSKEY#VK_UP
  - word: DOWN
    event: PRESSKEY#VK_DOWN
  - word: LEFT
    event: PRESSKEY#VK_LEFT
  - word: RIGHT
    event: PRESSKEY#VK_RIGHT
```

In this case, if write UP in twitch chat (on humiltat channel) the computer press UP key. Same for DOWN, LEFT and RIGHT

Is posible to use:

* Basic key -> PRESSKEY#{KEY}
* Key + shift -> PRESSKEY#SHIFT#{KEY}
* Key + ctrl -> PRESSKEY#CTRL#{KEY}
* Key + alt -> PRESSKEY#ALT#{KEY}
* Also can combine like -> PRESSKEY#SHIFT#CTRL#{KEY}

##Key control
To avoid pressing too many keys simultaneously, the eventfrequencyinms configuration parameter is available, which limits to a minimum of milliseconds between events.

As an example, if we indicate a value of 5000 (5 seconds), and in that time window several messages appear in the chat, a count will be made of all of them, and the most requested event will be executed.

The chat could look like this:

```yaml
User1: UP
User2: DOWN
User1: UP
```

And the application would execute the action "UP", since it has been most requested in that time window.

##key list

####Characters (A-Z)
* VK_A
* VK_B
* VK_C
* VK_D
* VK_E
* VK_F
* VK_G
* VK_H
* VK_I
* VK_J
* VK_K
* VK_L
* VK_M
* VK_N
* VK_O
* VK_P
* VK_Q
* VK_R
* VK_S
* VK_T
* VK_U
* VK_V
* VK_W
* VK_X
* VK_Y
* VK_Z

####Numbers (0-9)
* VK_0
* VK_1
* VK_2
* VK_3
* VK_4
* VK_5
* VK_6
* VK_7
* VK_8
* VK_9

####Numbers keypad (0-9)
* VK_KP0
* VK_KP1
* VK_KP2
* VK_KP3
* VK_KP4
* VK_KP5
* VK_KP6
* VK_KP7
* VK_KP8
* VK_KP9

####Function keys (F1-F12)
* VK_F1
* VK_F2
* VK_F3
* VK_F4
* VK_F5
* VK_F6
* VK_F7
* VK_F8
* VK_F9
* VK_F10
* VK_F11
* VK_F12

####Direction keys
* VK_UP
* VK_DOWN
* VK_RIGHT
* VK_LEFT

####Other keys
* VK_BACKSPACE
* VK_TAB
* VK_ENTER
* VK_BACKSLASH
* VK_COMMA
* VK_DOT
* VK_SPACE

The complete list is available at EventSender.go

### Detail on console

![Alt text](/doc/detail_UI_console.png?raw=true "Console view")

## TO-DO's

* Remove some hardcoded texts / configs
* UI Improvement
* More datasources (slack, telegram)
* More events (write text, computer action, administration, whitelists/blacklists)
