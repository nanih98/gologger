package gologger

import (
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
	l.Log.Printf("%sINFO: %s %s", info, msg, nocolor)

}

// Warn prints output to the console colored and with and emoji
func (l *CustomLogger) Warn(msg string) {
	l.Log.Printf("%sWARNING: %s %s", warning, msg, nocolor)
}

// Fatal prints output to the console colored and with and emoji. Returns exit 1
func (l *CustomLogger) Fatal(msg string) {
	l.Log.Printf("%sFATAL: %s %s", errorf, msg, nocolor)
}

// Success prints output to the console colored and with and emoji
func (l *CustomLogger) Success(msg string) {
	l.Log.Printf("%sSUCCESS: %s %s", success, msg, nocolor)
}

// Logger is the main function
func Logger() CustomLogger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	return CustomLogger{Log: logger}
}