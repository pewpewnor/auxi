package auxi

import (
	"fmt"
	"io"
	"log"

	"github.com/fatih/color"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	purple = color.New(color.FgMagenta).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	gray   = color.New(color.FgHiBlack).SprintFunc()
)

const (
	warningPrefix = "[WARNING]"
	errorPrefix   = "[ERROR]"
	successPrefix = "[SUCCESS]"
	infoPrefix    = "[INFO]"
)

func format(color func(...any) string, prefix string, message string) string {
	return color(prefix + " " + fmt.Sprint(message))
}

type Logger struct {
	*log.Logger
	flag int
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		log.New(out, prefix, flag),
		flag,
	}
}

func (l *Logger) AddTimestampFlag(flag int) {
	l.flag |= log.LstdFlags
	l.SetFlags(l.flag)
}

func (l *Logger) AddLineOfCodeFlag(flag int) {
	l.flag |= log.Llongfile
	l.SetFlags(l.flag)
}

func (l *Logger) Fatal(v ...any) {
	l.Fatalln(format(red, errorPrefix, fmt.Sprint(v...)))
}

func (l *Logger) Fatalf(s string, v ...any) {
	l.Fatalf(format(red, errorPrefix, fmt.Sprintf(s, v...)))
}

func (l *Logger) Warn(v ...any) {
	l.Println(format(yellow, warningPrefix, fmt.Sprint(v...)))
}

func (l *Logger) Warnf(s string, v ...any) {
	l.Println(format(yellow, warningPrefix, fmt.Sprintf(s, v...)))
}

func (l *Logger) Error(v ...any) {
	l.Println(format(red, errorPrefix, fmt.Sprint(v...)))
}

func (l *Logger) Errorf(s string, v ...any) {
	l.Println(format(red, errorPrefix, fmt.Sprintf(s, v...)))
}

func (l *Logger) Success(v ...any) {
	l.Println(format(green, successPrefix, fmt.Sprint(v...)))
}

func (l *Logger) Successf(s string, v ...any) {
	l.Println(format(green, successPrefix, fmt.Sprintf(s, v...)))
}

func (l *Logger) Info(v ...any) {
	l.Println(format(blue, infoPrefix, fmt.Sprint(v...)))
}

func (l *Logger) Infof(s string, v ...any) {
	l.Println(format(blue, infoPrefix, fmt.Sprintf(s, v...)))
}
