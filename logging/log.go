// chiwa - change i3 wallpapers automatically
//
// Author: slowpoke <mail+git@slowpoke.io>
//
// This program is free software under the non-terms
// of the Anti-License. Do whatever the fuck you want.
//
// Github: https://www.github.com/proxypoke/chiwa
// (Shortlink: https://git.io/chiwa)

// Package logging implements logging facilities for chiwa.
package logging

import (
	"log"
	"os"
)

// Logging levels
var (
	debug   bool
	verbose bool
)

// Logger
var (
	warningLogger = log.New(os.Stderr, "[WARNING] ", 0)
	verboseLogger = log.New(os.Stdout, "", 0)
	debugLogger   = log.New(os.Stderr, "[DEBUG] ", log.Ldate|log.Ltime|log.Llongfile)
)

// Activate Debug output.
func SetDebug() {
	debug = true
}

// Activate verbose output.
func SetVerbose() {
	verbose = true
}

// Log a warning. These cannot be disabled.
func Warnf(format string, v ...interface{}) {
	warningLogger.Printf(format, v...)
}

// Log a verbose message. These can be enabled with SetVerbose.
func Logf(format string, v ...interface{}) {
	if verbose {
		verboseLogger.Printf(format, v...)
	}
}

// Log a debug message. These can be enabled with SetDebug.
func Debugf(format string, v ...interface{}) {
	if debug {
		debugLogger.Printf(format, v...)
	}
}
