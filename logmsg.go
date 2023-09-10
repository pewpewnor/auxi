package auxi

import (
	"fmt"

	"github.com/fatih/color"
)

var Logmsg = logmsg{}

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

type logmsg struct {
	flag int
}

func (l *logmsg) Err(v ...any) string {
	return format(red, errorPrefix, fmt.Sprint(v...))
}

func (l *logmsg) Errf(s string, v ...any) string {
	return format(red, errorPrefix, fmt.Sprintf(s, v...))
}

func (l *logmsg) Fatal(v ...any) string {
	return format(red, errorPrefix, fmt.Sprint(v...))
}

func (l *logmsg) Fatalf(s string, v ...any) string {
	return format(red, errorPrefix, fmt.Sprintf(s, v...))
}

func (l *logmsg) Info(v ...any) string {
	return format(blue, infoPrefix, fmt.Sprint(v...))
}

func (l *logmsg) Infof(s string, v ...any) string {
	return format(blue, infoPrefix, fmt.Sprintf(s, v...))
}

func (l *logmsg) Success(v ...any) string {
	return format(green, successPrefix, fmt.Sprint(v...))
}

func (l *logmsg) Successf(s string, v ...any) string {
	return format(green, successPrefix, fmt.Sprintf(s, v...))
}

func (l *logmsg) Warn(v ...any) string {
	return format(yellow, warningPrefix, fmt.Sprint(v...))
}

func (l *logmsg) Warnf(s string, v ...any) string {
	return format(yellow, warningPrefix, fmt.Sprintf(s, v...))
}

func format(color func(...any) string, prefix string, message string) string {
	return color(fmt.Sprintf("%v %v", prefix, message))
}
