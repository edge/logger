# Edge Logger

This lightweight logging library is used throughout the Edge suite to standardise reporting.

Basic usage is simple:

```go
package myapp

import (
  "github.com/edge/logger"
)

func main() {
  l := logger.New().SetLabel("app", "com.example.myapp")
  go func() {
    l.Context("main.go_func").Info("Goroutine works")
  }()
  l.Context("main").Label("status", "OK").Info("All systems check")
}
```

## Using your own handler or middleware

The default logger handler (which writes to stdout/stderr) uses a fixed format with no options; if you want a more flexible logging library then you might want to use something like [logrus](https://github.com/sirupsen/logrus) instead.

However, it is possible to initialise a logger with a custom handler. Implement `logger.Handler` and pass an instance into `logger.New()` like so:

```go
package myapp

import (
  "github.com/edge/logger"
)

type MyHandler struct {
  s *logger.StdoutHandler
}

func NewMyHandler() *MyHandler {
  return &MyHandler{
    s: logger.NewStdoutHandler(),
  }
}

func (l *MyHandler) Log(e *logger.Entry) error {
  // always log as fatal error
  e.Severity = logger.Fatal
  return l.s.Log(e)
}

func main() {
  l := logger.New(NewMyHandler())
  l.Context("main").Info("Hello")
}
```
