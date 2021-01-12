package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_severityIntToText(t *testing.T) {
	a := assert.New(t)
	severities := GetSeverities()
	for i, ta := range severities {
		tb, ok := severityIntToText(i)
		if a.True(ok) {
			a.Equal(ta, tb)
		}
	}
}

func Test_severityIntToText_NotOk(t *testing.T) {
	a := assert.New(t)
	textSeverity, ok := severityIntToText(6)
	if a.False(ok) {
		a.Equal("", textSeverity)
	}
}

func Test_severityTextToInt(t *testing.T) {
	a := assert.New(t)
	severities := GetSeverities()
	for ia, t := range severities {
		ib, ok := severityTextToInt(t)
		if a.True(ok) {
			a.Equal(ia, ib)
		}
	}
}

func Test_severityTextToInt_NotOk(t *testing.T) {
	a := assert.New(t)
	// Severity zero-value is valid as an internal value, so we only test the ok return
	_, ok := severityTextToInt("")
	a.False(ok)
}
