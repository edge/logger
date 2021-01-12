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
	timeLayout string

	stderr *log.Logger
	stdout *log.Logger
}

// NewStdoutHandler creates a new StdoutHandler.
func NewStdoutHandler() *StdoutHandler {
	return &StdoutHandler{
		timeLayout: "2006-01-02 15:04:05.000 -0700 MST",

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
	logTime := time.Now().Format(s.timeLayout)
	message := s.formatMessage(e)
	textSeverity, _ := severityIntToText(e.Severity)
	return fmt.Sprintf("%[2]s %- 7[1]s %[3]s [%[4]s]: %[5]s",
		fmt.Sprintf("[%s]", textSeverity),
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
