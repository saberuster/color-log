package colorlog

import (
	"testing"
	"io"
	"bufio"
	"strings"
	"fmt"
)

func TestLogger_Debug(t *testing.T) {
	r, l := newLog()
	l.SetFrontColor(RED)
	l.SetBackColor(BLACK)
	l.SetDisplayMethod(Highlight)
	l.SetLevel(DEBUG)
	go l.Debug("debugging")
	br := bufio.NewReader(r)
	str, err := br.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(str) != fmt.Sprintf("%c[%d;%d;%dmdebugging%c[0m", 0x1B, Highlight, Frontend+RED, Backend+BLACK, 0x1B) {
		t.Fatal("str error")
	}
}

func TestLogger_Error(t *testing.T) {
	r, l := newLog()
	l.SetFrontColor(RED)
	go l.Error("error")
	br := bufio.NewReader(r)
	str, err := br.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(str) != fmt.Sprintf("%c[%d;%d;%dmerror%c[0m", 0x1B, Default, Frontend+RED, Backend+BLACK, 0x1B) {
		t.Fatal("str error")
	}
}

func TestLogger_Fatal(t *testing.T) {
	r, l := newLog()
	l.SetFrontColor(RED)
	go l.Fatal("fatal")
	br := bufio.NewReader(r)
	str, err := br.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(str) != fmt.Sprintf("%c[%d;%d;%dmfatal%c[0m", 0x1B, Default, Frontend+RED, Backend+BLACK, 0x1B) {
		t.Fatal("str error")
	}
}

func TestLogger_Info(t *testing.T) {
	r, l := newLog()
	l.SetFrontColor(RED)
	go l.Info("info")
	br := bufio.NewReader(r)
	str, err := br.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(str) != fmt.Sprintf("%c[%d;%d;%dminfo%c[0m", 0x1B, Default, Frontend+RED, Backend+BLACK, 0x1B) {
		t.Fatal("str error")
	}
}

func TestLogger_Warning(t *testing.T) {
	r, l := newLog()
	l.SetFrontColor(RED)
	go l.Warning("warning")
	br := bufio.NewReader(r)
	str, err := br.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(str) != fmt.Sprintf("%c[%d;%d;%dmwarning%c[0m", 0x1B, Default, Frontend+RED, Backend+BLACK, 0x1B) {
		t.Fatal("str error")
	}
}

func newLog() (io.Reader, *Logger) {
	r, w := io.Pipe()
	l := New(w, "", 0)
	return r, l
}
