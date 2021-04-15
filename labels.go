package logger

import (
	"fmt"
	"sort"
)

type Labels []*Label

func (l Labels) Len() int           { return len(l) }
func (l Labels) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l Labels) Less(i, j int) bool { return l[i].k < l[j].k }

// Sort triggers Labels sorting method.
func (l *Labels) Sort() {
	sort.Sort(l)
}

type Label struct {
	k string
	v interface{}
}

// Out outputs the label as a key=val string.
func (l *Label) Out() string {
	return fmt.Sprintf(`%s=%s`, l.k, l.Val())
}

// Val returns the value of the label.
func (l *Label) Val() string {
	switch v := l.v.(type) {
	case string:
		return v
	case func() string:
		return v()
	default:
		return ""
	}
}
