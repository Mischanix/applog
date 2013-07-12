// Package applog provides global logging methods so that multiple packages may
// log to a single application-defined stream.
//
// If any logging methods are called before SetOutput is called, os.Stdout will
// be used as a fallback.
package applog

import (
  "fmt"
  "io"
  "log"
  "os"
)

type LogLevel int

// Level determines the current level of logging output, from only Panics
// (PanicLevel) to everything (DebugLevel)
var Level LogLevel

const (
  PanicLevel LogLevel = iota
  ErrorLevel
  WarnLevel
  InfoLevel
  DebugLevel
  numLevels int = iota
)

var levelStrings = [numLevels]string{
  "panic",
  "error",
  "warn",
  "info",
  "debug",
}

const logFlags = log.LstdFlags | log.Lshortfile

var logger *log.Logger

func write(msg string, level LogLevel) {
  if Level >= level {
    if logger == nil {
      SetOutput(os.Stdout)
    }

    // 3 = Output, 2 = write, 1 = ..., 0 = caller
    // logger comes with its own mutex
    logger.Output(3, fmt.Sprintf(
      "[%s] %s",
      levelStrings[level],
      msg,
    ))
  }
}

func SetOutput(w io.Writer) {
  logger = log.New(w, "", logFlags)
}

func Panic(format string, v ...interface{}) {
  s := fmt.Sprintf(format, v...)
  write(s, PanicLevel)
  panic(s)
}

func Error(format string, v ...interface{}) {
  write(fmt.Sprintf(format, v...), ErrorLevel)
}

func Warn(format string, v ...interface{}) {
  write(fmt.Sprintf(format, v...), WarnLevel)
}

func Info(format string, v ...interface{}) {
  write(fmt.Sprintf(format, v...), InfoLevel)
}

func Debug(format string, v ...interface{}) {
  write(fmt.Sprintf(format, v...), DebugLevel)
}
