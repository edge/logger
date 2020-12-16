// Edge Network
// (c) 2019 Edge Network technologies Ltd.

package logger

import (
	"github.com/sirupsen/logrus"
)

// Instance is a type alias for logrus.FieldLogger.
type Instance struct {
	*logrus.Logger
	uiHook uiHook
}

const (
	PanicLevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
)

// LevelToInt maps logrus.Levels to an int value.
var LevelToInt = map[uint16]logrus.Level{
	0: DebugLevel,
	1: InfoLevel,
	2: WarnLevel,
	3: ErrorLevel,
	4: FatalLevel,
	5: PanicLevel,
}

// GetLevelsFromInt returns levels from an integer rating.
func GetLevelsFromInt(l uint16) (levels []logrus.Level) {
	for i, _ := range LevelToInt {
		if i >= l {
			levels = append(levels, LevelToInt[i])
		}
	}
	return
}

// GetLevelsFromLevel returns levels from a logrus.Level.
func GetLevelsFromLevel(l logrus.Level) []logrus.Level {
	for i, level := range LevelToInt {
		if level == l {
			return GetLevelsFromInt(i)
		}
	}

	// default to all levels.
	return logrus.AllLevels
}

// DefaultLogger is the default logger instance, calls to Context() etc are forwarded to this instance.
var DefaultLogger *Instance

func init() {
	DefaultLogger = New()
}

// New returns a new Logger Instance.
func New() *Instance {
	l := logrus.New()
	customFormatter := &logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	l.SetFormatter(customFormatter)
	i := &Instance{
		Logger: l,
		uiHook: uiHook{
			uiData: make(map[string]interface{}),
		},
	}

	// Add ui hook.
	i.AddHook(i.uiHook)
	return i
}

// Context adds field context to logs.
func (i *Instance) Context(c string) *logrus.Entry {
	return i.WithFields(logrus.Fields{
		"context": c,
	})
}

// Context adds field context to logs.
func Context(c string) *logrus.Entry {
	return DefaultLogger.WithFields(logrus.Fields{
		"context": c,
	})
}

// SetLogLevel sets the level of logging.
func (i *Instance) SetLogLevel(level string) {
	switch level {
	case "panic":
		i.SetLevel(logrus.PanicLevel)
	case "fatal":
		i.SetLevel(logrus.FatalLevel)
	case "error":
		i.SetLevel(logrus.ErrorLevel)
	case "warn":
		i.SetLevel(logrus.WarnLevel)
	case "info":
		i.SetLevel(logrus.InfoLevel)
	case "debug":
		i.SetLevel(logrus.DebugLevel)
	}
}

// SetLogLevel sets the level of logging.
func SetLogLevel(level string) {
	DefaultLogger.SetLogLevel(level)
}

// AddLogUI adds a log UI.
func AddLogUI(key string, value interface{}) {
	DefaultLogger.AddLogUI(key, value)
}

// AddLogUI adds a log UI.
func (i *Instance) AddLogUI(key string, value interface{}) {
	i.uiHook.uiData[key] = value
}
