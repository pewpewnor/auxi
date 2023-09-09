package auxi

import (
	"fmt"
	"io"
	"log"
	"os"

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

var logger = &Logger{
	log.New(os.Stdout, "[APP] ", 0),
	0,
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		log.New(out, prefix, flag),
		flag,
	}
}

func AddTimestampFlag(flag int) {
	logger.flag |= log.LstdFlags
	logger.SetFlags(logger.flag)
}

func AddLineOfCodeFlag(flag int) {
	logger.flag |= log.Llongfile
	logger.SetFlags(logger.flag)
}

func Fatal(v ...any) {
	logger.Fatalln(format(red, errorPrefix, fmt.Sprint(v...)))
}

func Fatalf(s string, v ...any) {
	logger.Fatalf(format(red, errorPrefix, fmt.Sprintf(s, v...)))
}

func Warn(v ...any) {
	logger.Println(format(yellow, warningPrefix, fmt.Sprint(v...)))
}

func Warnf(s string, v ...any) {
	logger.Println(format(yellow, warningPrefix, fmt.Sprintf(s, v...)))
}

func Error(v ...any) {
	logger.Println(format(red, errorPrefix, fmt.Sprint(v...)))
}

func Errorf(s string, v ...any) {
	logger.Println(format(red, errorPrefix, fmt.Sprintf(s, v...)))
}

func Success(v ...any) {
	logger.Println(format(green, successPrefix, fmt.Sprint(v...)))
}

func Successf(s string, v ...any) {
	logger.Println(format(green, successPrefix, fmt.Sprintf(s, v...)))
}

func Info(v ...any) {
	logger.Println(format(blue, infoPrefix, fmt.Sprint(v...)))
}

func Infof(s string, v ...any) {
	logger.Println(format(blue, infoPrefix, fmt.Sprintf(s, v...)))
}
