package logger

import (
	"fmt"
)

// Handler instances output log entries.
type Handler interface {
	Log(e *Entry) error
}

// Instance of a logger.
type Instance struct {
	// Labels that all entries inherit.
	Labels map[string]string
	// MinSeverity specifies the minimum severity to qualify an entry for logging. If an entry is less severe than this, it won't be output.
	MinSeverity Severity

	h Handler
}

// New logger instance. An optional Handler may be provided, otherwise the logger uses a StdoutHandler.
func New(h ...Handler) *Instance {
	if h == nil {
		h = []Handler{NewStdoutHandler()}
	}
	l := &Instance{
		Labels:      make(map[string]string, 0),
		MinSeverity: Info,

		h: h[0],
	}
	return l
}

// Context for a new log entry. De facto entry constructor.
func (l *Instance) Context(c string) (e *Entry) {
	e = &Entry{
		Context: c,
		Labels:  make(map[string]string, 0),
		l:       l,
	}
	// copy default labels to entry
	for k, v := range l.Labels {
		e.Labels[k] = v
	}
	return
}

// GetMinSeverity returns a text label representing MinSeverity.
func (l *Instance) GetMinSeverity() string {
	if textSeverity, ok := severityIntToText(l.MinSeverity); ok {
		return textSeverity
	}
	return ""
}

// Label adds a label that all entries inherit.
func (l *Instance) Label(k, v string) {
	l.Labels[k] = v
}

// Log an entry to output.
func (l *Instance) Log(e *Entry) {
	if e.Severity > l.MinSeverity {
		return
	}
	l.h.Log(e)
}

// SetMinSeverity parses a text label to replace MinSeverity. If an invalid label is given, an error will be returned with no side effects.
func (l *Instance) SetMinSeverity(textSeverity string) error {
	s, ok := severityTextToInt(textSeverity)
	if !ok {
		return fmt.Errorf("Invalid severity \"%s\", keeping current value \"%s\"", textSeverity, l.GetMinSeverity())
	}
	l.MinSeverity = s
	return nil
}
