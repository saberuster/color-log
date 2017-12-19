package colorlog

import (
	"log"
	"fmt"
	"io"
	"sync"
)

type color int
type displayMethod int
type level int

const (
	BLACK   color = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
	VERDANT
	WHITE
)

const (
	DEBUG level = 1 << iota
	INFO
	WARN
	ERROR
	FATAL
)

const Frontend = 30
const Backend = 40

const Default displayMethod = 0
const Highlight displayMethod = 1
const Underline displayMethod = 4
const Blinking displayMethod = 5
const Reverse displayMethod = 7

type Logger struct {
	dm     displayMethod //
	fc, bc color
	lv     level
	mu     sync.Mutex
	log.Logger
}

func New(out io.Writer, prefix string, flag int) *Logger {
	l := log.New(out, prefix, flag)
	return &Logger{
		Logger: *l,
		fc:     Frontend + WHITE,
		bc:     Backend + BLACK,
		dm:     Default,
		lv:     DEBUG,
	}
}

func (l *Logger) SetLevel(lv level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lv = lv
}

func (l *Logger) SetFrontColor(c color) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fc = Frontend + c
}

func (l *Logger) SetBackColor(c color) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.bc = Backend + c
}

func (l *Logger) SetDisplayMethod(dm displayMethod) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.dm = dm
}

func (l *Logger) formatMsg(v ...interface{}) string {
	wr := make([]byte, 0)
	dm, fc, bc := l.colorHeader()
	wr = append(wr, fmt.Sprintf("%c[%d;%d;%dm", 0x1B, dm, fc, bc)...)

	for _, value := range v {
		wr = append(wr, fmt.Sprint(value)...)
	}

	wr = append(wr, fmt.Sprintf("%c[0m", 0x1B)...)
	return string(wr)
}

func (l *Logger) colorHeader() (displayMethod, color, color) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.dm, l.fc, l.bc
}

func (l *Logger) Debug(v ...interface{}) {
	if l.lv <= DEBUG {
		l.Output(2, l.formatMsg(v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.lv <= INFO {
		l.Output(2, l.formatMsg(v...))
	}
}

func (l *Logger) Warning(v ...interface{}) {
	if l.lv <= WARN {
		l.Output(2, l.formatMsg(v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.lv <= ERROR {
		l.Output(2, l.formatMsg(v...))
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.lv <= FATAL {
		l.Output(2, l.formatMsg(v...))
	}
}
