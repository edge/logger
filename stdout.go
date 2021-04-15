package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const newLine string = "\n"

// StdoutHandler is a logger handler that outputs to stdout/stderr.
type StdoutHandler struct {
	timeLayout string
}

// NewStdoutHandler creates a new StdoutHandler.
func NewStdoutHandler() *StdoutHandler {
	return &StdoutHandler{
		timeLayout: "2006-01-02 15:04:05.000 -0700 MST",
	}
}

// Log an entry to stdout/stderr.
func (s *StdoutHandler) Log(e *Entry) error {
	if e.Severity <= Error {
		s.formatEntry(os.Stderr, e)
		return nil
	}
	s.formatEntry(os.Stdout, e)
	return nil
}

func (s *StdoutHandler) formatEntry(w io.Writer, e *Entry) {
	labels := s.formatLabels(e)
	logTime := time.Now().Format(s.timeLayout)
	message := s.formatMessage(e)
	textSeverity, _ := severityIntToText(e.Severity)
	severityColor := severityIntToColor(e.Severity)
	fmt.Fprintf(w, "\x1b[%[1]dm%- 5[2]s\x1b[0m %[3]s %[4]s [%[5]s]: %[6]s %s",
		severityColor,
		fmt.Sprintf("%s", textSeverity),
		logTime,
		e.Context,
		labels,
		message,
		newLine,
	)
}

func (*StdoutHandler) formatLabels(e *Entry) (labels string) {
	ll := []string{}
	for _, v := range *e.Labels {
		ll = append(ll, v.Out())
	}
	return strings.Join(ll, " ")
}

func (*StdoutHandler) formatMessage(e *Entry) string {
	return fmt.Sprint(e.Message...)
}
