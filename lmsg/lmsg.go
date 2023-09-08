package lmsg

import (
	"fmt"

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

func Warn(v ...any) string {
	return format(yellow, warningPrefix, fmt.Sprint(v...))
}

func Warnf(s string, v ...any) string {
	return format(yellow, warningPrefix, fmt.Sprintf(s, v...))
}

func Error(v ...any) string {
	return format(red, errorPrefix, fmt.Sprint(v...))
}

func Errorf(s string, v ...any) string {
	return format(red, errorPrefix, fmt.Sprintf(s, v...))
}

func Success(v ...any) string {
	return format(green, successPrefix, fmt.Sprint(v...))
}

func Successf(s string, v ...any) string {
	return format(green, successPrefix, fmt.Sprintf(s, v...))
}

func Info(v ...any) string {
	return format(blue, infoPrefix, fmt.Sprint(v...))
}

func Infof(s string, v ...any) string {
	return format(blue, infoPrefix, fmt.Sprintf(s, v...))
}
