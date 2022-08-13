package gologger

import (
	"fmt"
	"log"
	"os"
)

const (
	success = "\033[92m"
	nocolor = "\033[0m"
	info    = "\033[96m"
	warning = "\033[93m"
	errorf  = "\033[91m"
)

// CustomLogger is the main struct
type CustomLogger struct {
	Log *log.Logger
}

// Info prints output to the console colored and with and emoji
func (l *CustomLogger) Info(msg string) {
	l.Log.SetPrefix(fmt.Sprintf("%sINFO%s ", info, nocolor))
	l.Log.SetFlags(log.Lshortfile)
	l.Log.Printf("%s%s%s", info, msg, nocolor)
}

// Warn prints output to the console colored and with and emoji
func (l *CustomLogger) Warn(msg string) {
	l.Log.SetPrefix(fmt.Sprintf("%sWARNING%s ", warning, nocolor))
	l.Log.SetFlags(log.Lshortfile)
	l.Log.Printf("%s%s%s", warning, msg, nocolor)
}

// Fatal prints output to the console colored and with and emoji. Returns exit 1
func (l *CustomLogger) Fatal(err error) {
	l.Log.SetPrefix(fmt.Sprintf("%sFATAL%s ", errorf, nocolor))
	l.Log.SetFlags(log.Lshortfile)
	l.Log.Fatalf("%s%v%s", errorf, err, nocolor)
}

// Success prints output to the console colored and with and emoji
func (l *CustomLogger) Success(msg string) {
	l.Log.SetPrefix(fmt.Sprintf("%sSUCCESS%s ", success, nocolor))
	l.Log.SetFlags(log.Lshortfile)
	l.Log.Printf("%s%s%s", success, msg, nocolor)
}

// Logger is the main function
func Logger() CustomLogger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return CustomLogger{Log: logger}
}