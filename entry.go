package logger

import "fmt"

// Entry represents a single log entry.
type Entry struct {
	Context  string
	Labels   *Labels
	Severity Severity
	Message  []interface{}

	l *Instance
}

// Debug logs a DEBUG message.
func (e *Entry) Debug(args ...interface{}) {
	e.vlog(Debug, args)
}

// Debugf logs a DEBUG message.
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.vlogf(Debug, format, args)
}

// Error logs an ERROR.
func (e *Entry) Error(args ...interface{}) {
	e.vlog(Error, args)
}

// Errorf logs an ERROR.
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.vlogf(Error, format, args)
}

// Fatal logs a FATAL error.
func (e *Entry) Fatal(args ...interface{}) {
	e.vlog(Fatal, args)
}

// Fatalf logs a FATAL error.
func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.vlogf(Fatal, format, args)
}

// Info logs INFO.
func (e *Entry) Info(args ...interface{}) {
	e.vlog(Info, args)
}

// Infof logs INFO.
func (e *Entry) Infof(format string, args ...interface{}) {
	e.vlogf(Info, format, args)
}

// Label sets a label on the entry. The same entry is returned.
func (e *Entry) Label(k string, v interface{}) *Entry {
	for _, lbl := range *e.Labels {
		if lbl.k == k {
			lbl.v = v
			goto SORT
		}
	}
	// Insert new label
	*e.Labels = append(*e.Labels, &Label{k, v})
SORT:
	e.Labels.Sort()
	return e
}

// Trace logs a TRACE message.
func (e *Entry) Trace(args ...interface{}) {
	e.vlog(Trace, args)
}

// Tracef logs a TRACE message.
func (e *Entry) Tracef(format string, args ...interface{}) {
	e.vlogf(Trace, format, args)
}

// Warn logs a WARN(ing).
func (e *Entry) Warn(args ...interface{}) {
	e.vlog(Warn, args)
}

// Warnf logs a WARN(ing).
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.vlogf(Warn, format, args)
}

func (e *Entry) vlog(s Severity, args []interface{}) {
	e.Severity = s
	e.Message = args
	e.l.Log(e)
}

func (e *Entry) vlogf(s Severity, format string, args []interface{}) {
	finalMsg := fmt.Sprintf(format, args...)
	e.vlog(s, []interface{}{finalMsg})
}
