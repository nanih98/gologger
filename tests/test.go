package main

import (
	"log"
	"os"

	"github.com/nanih98/gologger"
)

func main() {
	log := gologger.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)

	log.Info("Test")
}