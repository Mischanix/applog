package applog

import (
  "bytes"
  "os"
  "testing"
)

var buf = bytes.NewBuffer(nil)

func TestSetOutput(t *testing.T) {
  if logger != nil {
    t.Error("logger non-nil at start")
  }

  SetOutput(buf)
  if logger == nil {
    t.Error("logger nil after SetOutput")
  }
}

func TestPanic(t *testing.T) {
  lenBefore := buf.Len()
  SetOutput(buf)
  Level = PanicLevel

  defer func() {
    recover()
    if buf.Len() <= lenBefore {
      t.Error("Panic output nothing")
    }
  }()
  Panic("Panic test")
}

func TestError(t *testing.T) {
  lenBefore := buf.Len()
  SetOutput(buf)
  Level = PanicLevel
  Error("Error test (DND)")
  if buf.Len() > lenBefore {
    t.Error("Error output something at PanicLevel")
  }
  lenBefore = buf.Len()

  Level = ErrorLevel
  Error("Error test")
  if buf.Len() <= lenBefore {
    t.Error("Error output nothing at ErrorLevel")
  }
}

func TestWarn(t *testing.T) {
  lenBefore := buf.Len()
  SetOutput(buf)
  Level = PanicLevel
  Warn("Warn test (DND)")
  if buf.Len() > lenBefore {
    t.Error("Warn output something at PanicLevel")
  }
  lenBefore = buf.Len()

  Level = WarnLevel
  Warn("Warn test")
  if buf.Len() <= lenBefore {
    t.Error("Warn output nothing at WarnLevel")
  }
}

func TestInfo(t *testing.T) {
  lenBefore := buf.Len()
  SetOutput(buf)
  Level = PanicLevel
  Info("Info test (DND)")
  if buf.Len() > lenBefore {
    t.Error("Info output something at PanicLevel")
  }
  lenBefore = buf.Len()

  Level = InfoLevel
  Info("Info test")
  if buf.Len() <= lenBefore {
    t.Error("Info output nothing at InfoLevel")
  }
}

func TestDebug(t *testing.T) {
  lenBefore := buf.Len()
  SetOutput(buf)
  Level = PanicLevel
  Debug("Debug test (DND)")
  if buf.Len() > lenBefore {
    t.Error("Debug output something at PanicLevel")
  }
  lenBefore = buf.Len()

  Level = DebugLevel
  Debug("Debug test")
  if buf.Len() <= lenBefore {
    t.Error("Debug output nothing at DebugLevel")
  }
}

func TestSample(t *testing.T) {
  // Throw some sample outputs in the test
  SetOutput(os.Stdout)
  Level = DebugLevel
  defer func() {
    recover()
  }()
  Debug("Debug sample")
  Info("Info sample")
  Warn("Warn sample")
  Error("Error sample")
  Panic("Panic sample")
}
