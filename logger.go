package logger

import (
	"fmt"
	"sync"
)

// Handler instances output log entries.
type Handler interface {
	Log(e *Entry) error
}

// Instance of a logger.
type Instance struct {
	// MinSeverity specifies the minimum severity to qualify an entry for logging. If an entry is less severe than this, it won't be output.
	MinSeverity Severity

	// Labels that all entries inherit.
	labels map[string]string

	h Handler
	m *sync.RWMutex
}

// New logger instance. An optional Handler may be provided, otherwise the logger uses a StdoutHandler.
func New(h ...Handler) *Instance {
	if h == nil {
		h = []Handler{NewStdoutHandler()}
	}
	l := &Instance{
		MinSeverity: Info,

		labels: make(map[string]string, 0),

		h: h[0],
		m: &sync.RWMutex{},
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
	for k, v := range l.GetLabels() {
		e.Labels[k] = v
	}

	return
}

// GetLabels returns all the labels configured on this logger.
func (l *Instance) GetLabels() map[string]string {
	l.m.Lock()
	defer l.m.Unlock()

	labels := map[string]string{}
	for k, v := range l.labels {
		labels[k] = v
	}
	return labels
}

// GetMinSeverity returns a text label representing MinSeverity.
func (l *Instance) GetMinSeverity() string {
	if textSeverity, ok := severityIntToText(l.MinSeverity); ok {
		return textSeverity
	}
	return ""
}

// Log an entry to output.
func (l *Instance) Log(e *Entry) {
	if e.Severity > l.MinSeverity {
		return
	}
	l.h.Log(e)
}

// SetLabel configures a label that all entries inherit.
func (l *Instance) SetLabel(k, v string) {
	l.m.Lock()
	defer l.m.Unlock()
	l.labels[k] = v
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
