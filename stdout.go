package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// StdoutHandler is a logger handler that outputs to stdout/stderr.
type StdoutHandler struct {
	// format will be used in fmt.Sprintf() and is provided the following replacement strings:
	// 1. severity text flag e.g. ERROR, INFO
	// 2. timestamp
	// 3. context
	// 4. labels in key1=value key2=value format
	// 5. message
	format string

	timeLayout string

	stderr *log.Logger
	stdout *log.Logger
}

// NewStdoutHandler creates a new StdoutHandler.
func NewStdoutHandler() *StdoutHandler {
	return &StdoutHandler{
		format:     "%[1]s[%[2]s] %[5]s context=%[3]s %[4]s",
		timeLayout: time.RFC3339Nano,

		stderr: log.New(os.Stderr, "", 0),
		stdout: log.New(os.Stdout, "", 0),
	}
}

// Log an entry to stdout/stderr.
func (s *StdoutHandler) Log(e *Entry) error {
	line := s.formatEntry(e)
	if e.Severity <= Error {
		s.stderr.Println(line)
	} else {
		s.stdout.Println(line)
	}
	return nil
}

func (s *StdoutHandler) formatEntry(e *Entry) string {
	labels := s.formatLabels(e)
	logTime := time.Now().UTC().Format(s.timeLayout)
	message := s.formatMessage(e)
	textSeverity, _ := severityIntToText(e.Severity)
	return fmt.Sprintf(s.format,
		textSeverity,
		logTime,
		e.Context,
		labels,
		message,
	)
}

func (*StdoutHandler) formatLabels(e *Entry) (labels string) {
	ll := []string{}
	for k, v := range e.Labels {
		ll = append(ll, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(ll, " ")
}

func (*StdoutHandler) formatMessage(e *Entry) string {
	return fmt.Sprint(e.Message...)
}
