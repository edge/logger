package logger

/*
This file adds backward compatibility for existing edge code

Everything here is considered deprecated and will be removed after all affected code is refactored to not require it.
*/

var defaultLogger *Instance

// Context for a new log entry. DEPRECATED - will log a warning.
func Context(c string) (e *Entry) {
	if defaultLogger == nil {
		defaultLogger = New()
		defaultLogger.MinSeverity = Debug
	}
	defaultLogger.Context("logger").Warnf("Use of deprecated log.Context() in local context=\"%s\" - migrate to instance-oriented logging ASAP", c)
	e = defaultLogger.Context(c)
	return
}

// Printf logs an INFO entry.
//
// Provided for compatibility with bigcache.Logger.
func (l *Instance) Printf(format string, v ...interface{}) {
	l.Context("Printf").Infof(format, v...)
}
