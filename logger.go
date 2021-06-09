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
	//
	// The convenience functions GetMinSeverity() and SetMinSeverity() allow you to set this using a text label instead of integer value, and are recommended for long-term compatibility.
	MinSeverity Severity

	// Labels that all entries inherit. A dynamic label is calculated each time an entry is created and may be useful for logging changeable state data.
	labels *Labels

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
		labels:      &Labels{},
		h:           h[0],
		m:           &sync.RWMutex{},
	}
	return l
}

// Context for a new log entry. De facto entry constructor.
func (l *Instance) Context(c string) (e *Entry) {
	e = &Entry{
		Context: c,
		Labels:  l.labels,
		l:       l,
	}
	// copy default labels to entry
	for _, lbl := range *l.labels {
		e.Label(lbl.k, lbl.v)
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

// Log an entry to output.
func (l *Instance) Log(e *Entry) {
	if e.Severity >= l.MinSeverity {
		return
	}
	l.h.Log(e)
}

// SetDynamicLabel configures a label that is calculated for each entry when it is created (at which time instance labels are copied).
func (l *Instance) SetDynamicLabel(k string, v func() string) *Instance {
	return l.SetLabel(k, v)
}

// SetLabel configures a label that all entries inherit. The same instance is returned.
func (l *Instance) SetLabel(k string, v interface{}) *Instance {
	l.m.Lock()
	defer l.m.Unlock()
	// Search for existing label and update
	for _, lbl := range *l.labels {
		if lbl.k == k {
			lbl.v = v
			goto SORT
		}
	}
	// Insert new label
	*l.labels = append(*l.labels, &Label{k, v})
SORT:
	l.labels.Sort()
	return l
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
