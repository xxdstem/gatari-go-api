package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var latest chan log

type Logger struct {
	Mode uint8
}

type log struct {
	Time time.Time
	Mode uint8
	Body string
}

type header struct {
	Prefix string
	Color  *color.Color
}

func Init() {
	latest = make(chan log, 100000)
	// latest <- log{
	// 	Time: time.Now(),
	// 	Mode: 1,
	// 	Body: "nigger",
	// }
	go func() {
		for {
			c := <-latest

			var h header

			switch c.Mode {
			case MODE_INFO:
				{
					h.Color = blackOnWhite
					h.Prefix = " INFO "
				}
			case MODE_ERR:
				{
					h.Color = whiteOnBrightRed
					h.Prefix = " ERR* "
				}
			}

			printLog(c.Time, h, c.Body)
		}
	}()
}

func New(opts ...interface{}) *Logger {
	var L Logger

	for _, o := range opts {
		if o == NOCOLOR {
			L.Mode = 1
		}
	}

	return &L
}

func printLog(T time.Time, h header, text string) {
	t := T.Format(timeLayout)

	h.Color.Printf(" ")
	fmt.Printf(" [ %s ] ", t)

	h.Color.Printf(h.Prefix)

	fmt.Printf(" %s\n", text)
}

func (l *Logger) Info(text string) {
	latest <- log{
		Time: time.Now(),
		Mode: MODE_INFO,
		Body: text,
	}
}

func (l *Logger) Error(text string) {
	latest <- log{
		Time: time.Now(),
		Mode: MODE_ERR,
		Body: text,
	}
}
