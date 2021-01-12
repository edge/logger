package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHandler struct {
	cb func(e *Entry)
	s  *StdoutHandler
}

func newTestHandler(cb func(e *Entry)) *testHandler {
	return &testHandler{
		cb: cb,
		s:  NewStdoutHandler(),
	}
}

func (h *testHandler) Log(e *Entry) error {
	h.cb(e)
	h.s.Log(e)
	return nil
}

func newWithTestHandler(cb func(e *Entry)) *Instance {
	return New(newTestHandler(cb))
}

func Test_Logger_Context(t *testing.T) {
	a := assert.New(t)
	l := New()
	var e *Entry

	e = l.Context("test")
	a.Equal("test", e.Context)

	l.SetLabel("hello", "world")
	e = l.Context("test2")
	a.Equal("test2", e.Context)
	for k, v := range e.Labels {
		a.Equal(k, "hello")
		a.Equal(v, "world")
	}
}

func Test_Logger_DynamicLabel(t *testing.T) {
	a := assert.New(t)

	n := 0
	l := newWithTestHandler(func(e *Entry) {
		for k, v := range e.Labels {
			a.Equal("dyn", k)
			a.Equal(fmt.Sprintf("%d", n), v)
		}
	})
	l.SetDynamicLabel("dyn", func() string {
		n++
		return fmt.Sprintf("%d", n)
	})

	x := 0
	for {
		x++
		l.Context("test").Infof("dynamic log - should show dyn=%d", x)
		if x == 10 {
			break
		}
	}
}

func Test_Logger_Label(t *testing.T) {
	a := assert.New(t)

	l := New()
	l.SetLabel("test", "ABC")

	for k, v := range l.GetLabels() {
		a.Equal("test", k)
		a.Equal("ABC", v)
	}
}

func Test_Logger_Log(t *testing.T) {
	a := assert.New(t)
	var l *Instance

	// all tests at default severity, Info
	l = newWithTestHandler(func(e *Entry) {
		a.Equal("test", e.Context)
		for k, v := range e.Labels {
			a.Equal("hello", k)
			a.Equal("world", v)
		}
		a.Equal("MESSAGE", e.Message[0])
		a.Equal(Info, e.Severity)
	})
	// assert info is default MinSeverity, else set manually for testing purposes
	if !a.Equal(Info, l.MinSeverity) {
		l.SetMinSeverity("INFO")
	}
	l.Context("test").Label("hello", "world").Info("MESSAGE")
	l.Context("test").Label("hello", "world").Infof("%s%s", "MESS", "AGE")

	// simple checks for all severities
	for i, textSeverity := range GetSeverities() {
		// check severity value is correct
		l = newWithTestHandler(func(e *Entry) {
			a.Equal(i, e.Severity)
		})
		l.SetMinSeverity(textSeverity)
		logWithSeverity(l.Context("test"), textSeverity, nil)
		logfWithSeverity(l.Context("test"), textSeverity, "%s%s", "MESS", "AGE")

		// check that output is blocked if insufficiently severe. skip FATAL as it cannot be blocked
		if i > Fatal {
			l = newWithTestHandler(func(e *Entry) {
				a.Failf("Unexpected logging", "Entry with severity \"%s\" should not be output", textSeverity)
			})
			l.MinSeverity = i - 1
			logWithSeverity(l.Context("test"), textSeverity, nil)
		}
	}
}

func Test_Logger_MinSeverity(t *testing.T) {
	a := assert.New(t)
	l := New()
	for k, v := range l.GetLabels() {
		l.SetMinSeverity(v)
		a.Equal(k, l.MinSeverity)
		a.Equal(v, l.GetMinSeverity())
	}
}

func Test_Logger_MinSeverity_Error(t *testing.T) {
	a := assert.New(t)
	l := New()

	origSeverity := l.MinSeverity
	err := l.SetMinSeverity("")
	a.Error(err)
	a.Equal(origSeverity, l.MinSeverity)
}

func logWithSeverity(e *Entry, textSeverity string, args ...interface{}) {
	ok := true
	switch textSeverity {
	case "TRACE":
		e.Trace(args)
		break
	case "DEBUG":
		e.Debug(args)
		break
	case "INFO":
		e.Info(args)
		break
	case "WARN":
		e.Warn(args)
		break
	case "ERROR":
		e.Error(args)
		break
	case "FATAL":
		e.Fatal(args)
		break
	default:
		ok = false
		break
	}
	if !ok {
		panic(fmt.Errorf("Invalid severity level"))
	}
}

func logfWithSeverity(e *Entry, textSeverity string, msg string, args ...interface{}) {
	ok := true
	switch textSeverity {
	case "TRACE":
		e.Tracef(msg, args...)
		break
	case "DEBUG":
		e.Debugf(msg, args...)
		break
	case "INFO":
		e.Infof(msg, args...)
		break
	case "WARN":
		e.Warnf(msg, args...)
		break
	case "ERROR":
		e.Errorf(msg, args...)
		break
	case "FATAL":
		e.Fatalf(msg, args...)
		break
	default:
		ok = false
		break
	}
	if !ok {
		panic(fmt.Errorf("Invalid severity level"))
	}
}
