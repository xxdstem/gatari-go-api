package logger

import (
	"fmt"
	"time"
)

var latest chan log

func Init() {
	latest = make(chan log)

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
			case MODE_WARN:
				{
					h.Color = whiteOnYellow
					h.Prefix = " WARN "
				}
			case MODE_DONE:
				{
					h.Color = whiteOnGreen
					h.Prefix = " DONE "
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

func (l *Logger) Done(text string) {
	latest <- log{
		Time: time.Now(),
		Mode: MODE_DONE,
		Body: text,
	}
}

func (l *Logger) Warn(text string) {
	latest <- log{
		Time: time.Now(),
		Mode: MODE_WARN,
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
