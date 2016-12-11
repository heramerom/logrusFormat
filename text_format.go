package logrusFormat 

import (
	"bytes"
	"fmt"
	"github.com/Sirupsen/logrus"
	"runtime"
	"strings"
)

const (
	_nocolor = "\033[0m"
	_red     = "\033[91m"
	_green   = "\033[92m"
	_yellow  = "\033[93m"
	_magenta = "\033[95m"
	_cyan    = "\033[96m"
)

const (
	_time_format = "2006-01-02 15:04:05"
)

var (
	isTerminal bool
)

func init() {
	isTerminal = logrus.IsTerminal()
}

type TextFormat struct {
	ForceColors bool
}

func (f *TextFormat) Format(entry *logrus.Entry) ([]byte, error) {
	levelText := strings.ToUpper(entry.Level.String())[0:4]
	buf := bytes.NewBuffer(make([]byte, 0, 32))
	if (f.ForceColors || isTerminal) && runtime.GOOS != "windows" {
		color := _nocolor
		switch entry.Level {
		case logrus.DebugLevel:
			color = _cyan
		case logrus.InfoLevel:
			color = _green
		case logrus.WarnLevel:
			color = _yellow
		case logrus.ErrorLevel:
			color = _magenta
		case logrus.PanicLevel, logrus.FatalLevel:
			color = _red
		}
		buf.WriteString(color)
	}
	buf.WriteString(fmt.Sprintf("[%s] ", entry.Time.Format(_time_format)))
	buf.WriteString(fmt.Sprintf("[%s] ", levelText))
	for k, v := range entry.Data {
		buf.WriteString(fmt.Sprintf("[%s=%v] ", k, v))
	}
	buf.WriteString(entry.Message)
	if (f.ForceColors || isTerminal) && runtime.GOOS != "windows" {
		buf.WriteString(_nocolor)
	}
	buf.WriteString("\n")
	return buf.Bytes(), nil
}
