// Edge Network
// (c) 2019 Edge Network technologies Ltd.

package logger

import (
	"github.com/sirupsen/logrus"
)

// serviceLogger allows us to add hooks to service logger.
type uiHook struct {
	uiData map[string]interface{}
}

// Levels specify the levels to which we want to attach the hook.
func (uiHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

// Fire is fired before any log.
func (u uiHook) Fire(e *logrus.Entry) error {
	for key, val := range u.uiData {

		// if val is a function, get it's output
		outFunc, ok := val.(func() string)
		if ok {
			val = safeRunFunc(outFunc)
		}

		// do not show empty fields
		if val == "" {
			continue
		}
		e.Data[key] = val
	}
	return nil
}

// safeRunFunc safely runs a user function
func safeRunFunc(f func() string) string {
	defer recover()
	return f()
}
