package auxi

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

var Logger = Newlogger(os.Stdout, "[APP] ", 0)

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

type logger struct {
	*log.Logger
	flag int
}

func Newlogger(out io.Writer, prefix string, flag int) *logger {
	return &logger{
		log.New(out, prefix, flag),
		flag,
	}
}

func (l *logger) AddTimestampFlag() {
	l.flag |= log.LstdFlags
	l.SetFlags(l.flag)
}

func (l *logger) AddLineOfCodeFlag() {
	l.flag |= log.Llongfile
	l.SetFlags(l.flag)
}

func (l *logger) Errorln(v ...any) {
	l.Println(format(red, errorPrefix, fmt.Sprint(v...)))
}

func (l *logger) Errorf(s string, v ...any) {
	l.Println(format(red, errorPrefix, fmt.Sprintf(s, v...)))
}

func (l *logger) Fatalln(v ...any) {
	l.Fatalln(format(red, errorPrefix, fmt.Sprint(v...)))
}

func (l *logger) Fatalf(s string, v ...any) {
	l.Fatalln(format(red, errorPrefix, fmt.Sprintf(s, v...)))
}

func (l *logger) Infoln(v ...any) {
	l.Println(format(blue, infoPrefix, fmt.Sprint(v...)))
}

func (l *logger) Infof(s string, v ...any) {
	l.Println(format(blue, infoPrefix, fmt.Sprintf(s, v...)))
}

func (l *logger) Successln(v ...any) {
	l.Println(format(green, successPrefix, fmt.Sprint(v...)))
}

func (l *logger) Successf(s string, v ...any) {
	l.Println(format(green, successPrefix, fmt.Sprintf(s, v...)))
}

func (l *logger) Warnln(v ...any) {
	l.Println(format(yellow, warningPrefix, fmt.Sprint(v...)))
}

func (l *logger) Warnf(s string, v ...any) {
	l.Println(format(yellow, warningPrefix, fmt.Sprintf(s, v...)))
}

func format(color func(...any) string, prefix string, message string) string {
	return color(fmt.Sprintf("%v %v", prefix, message))
}
