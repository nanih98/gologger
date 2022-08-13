// package main

// import (
// 	"log"
// 	"os"
// )

// var (
// 	SuccessLogger *log.Logger
// 	WarningLogger *log.Logger
// 	InfoLogger    *log.Logger
// 	ErrorLogger   *log.Logger
// )

// func init() {
// 	// file, err := os.OpenFile("/tmp/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 	// if err != nil {
// 	//     log.Fatal(err)
// 	// }
// 	SuccessLogger = log.New(os.Stdout, "SUCCESS |✅| ", log.Ldate|log.Ltime|log.Lshortfile)
// 	InfoLogger = log.New(os.Stdout, "INFO |ℹ️ | ", log.Ldate|log.Ltime|log.Lshortfile)
// 	WarningLogger = log.New(os.Stdout, "WARNING |⚠️ | ", log.Ldate|log.Ltime|log.Lshortfile)
// 	ErrorLogger = log.New(os.Stdout, "ERROR |🔥| ", log.Ldate|log.Ltime|log.Lshortfile)
// }

// func main() {
// 	SuccessLogger.Println("Successfully blablabla...")
// 	InfoLogger.Println("Starting the application...")
// 	InfoLogger.Println("Something noteworthy happened")
// 	WarningLogger.Println("There is something you should know about")
// 	ErrorLogger.Println("Something went wrong")
// }

package main

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
func (l *CustomLogger) Info(msg ...any) {
	InfoLogger := log.New(os.Stdout, fmt.Sprintf("%sINFO |ℹ️ | %s", info, nocolor), log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger.Printf("%s Info log %s", info, nocolor)
}

// Warn prints output to the console colored and with and emoji
func (l *CustomLogger) Warn(msg ...any) {
	WarningLogger := log.New(os.Stdout, fmt.Sprintf("%sWARNING |⚠️ |%s", warning, nocolor), log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger.Printf("%s Warning log %s", warning, nocolor)
}

// Fatal prints output to the console colored and with and emoji. Returns exit 1
func (l *CustomLogger) Fatal(msg ...any) {
	FatalLogger := log.New(os.Stdout, fmt.Sprintf("%sERROR |❌| %s", errorf, nocolor), log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger.Fatal(fmt.Sprintf("%s Fatal log %s", errorf, nocolor))
}

// Success prints output to the console colored and with and emoji
func (l *CustomLogger) Success(msg ...any) {
	SuccessLogger := log.New(os.Stdout, fmt.Sprintf("%sSUCCESS |✅| %s", success, nocolor), log.Ldate|log.Ltime|log.Lshortfile)
	SuccessLogger.Printf("%s Success log %s", success, nocolor)
}

// Logger is the main function
func Logger() CustomLogger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	return CustomLogger{Log: logger}
}

func main() {
	log := Logger()

	log.Info("Info message")
	log.Warn("Warning message")
	log.Success("Success logger")
	log.Fatal("Fatal message")
}
