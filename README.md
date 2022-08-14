<p align="center" >
  <img src="logo.png" alt="logo" width="250"/>
  <h3 align="center">go-logger</h3>
  <p align="center">Simple package for logging in golang</p>
</p>

<p align="center" >
  <img alt="Go report card" src="https://goreportcard.com/badge/github.com/nanih98/gologger">
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/nanih98/gologger">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/nanih98/gologger">
</p>

---

# Simple usage

```go
package main

import (
 "fmt"
 "log"
 "os"

 "github.com/nanih98/gologger"
)

func main() {
 log := gologger.New(os.Stdout, "", log.Ldate|log.Ltime)
 log.Info("Info log")
 log.Warn("Warning log")
 log.Success("Success log")
 log.Fatal(fmt.Errorf("Fatal error"))
 //....
}
```
# Example
  <img alt="Gologger example" src="gologger-example.png">

# About

This library is a simple implementation of the native library log. Is almost the same but with colored output.

It is not for production environments, just to learn how to create your own library :)

For production environments or real apps, use the following libraries:

* Zap
* Zerolog
* Logrus (my favourite)

