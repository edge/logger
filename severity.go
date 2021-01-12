package logger

const (
	// Fatal error/panic, not recoverable
	Fatal     Severity = 0
	fatalText string   = "FATAL"

	// Error encountered, possibly recoverable
	Error     Severity = 1
	errorText string   = "ERROR"

	// Warn of potential operating concerns or failure trajectory
	Warn     Severity = 2
	warnText string   = "WARN"

	// Info updates
	Info     Severity = 3
	infoText string   = "INFO"

	// Debug data. Solely used in development
	Debug     Severity = 4
	debugText string   = "DEBUG"

	// Trace data. Solely used in development
	Trace     Severity = 5
	traceText string   = "TRACE"
)

// Severity of a log entry is represented by a number where 0 is most severe (Fatal) and N is least.
type Severity uint8

// GetSeverities returns a dictionary of severity text labels indexed by their internal values.
func GetSeverities() map[Severity]string {
	return map[Severity]string{
		Fatal: fatalText,
		Error: errorText,
		Warn:  warnText,
		Info:  infoText,
		Debug: debugText,
		Trace: traceText,
	}
}

func severityIntToText(s Severity) (text string, ok bool) {
	ok = true
	switch s {
	case Debug:
		text = debugText
		break
	case Error:
		text = errorText
		break
	case Fatal:
		text = fatalText
		break
	case Info:
		text = infoText
		break
	case Trace:
		text = traceText
		break
	case Warn:
		text = warnText
		break
	default:
		ok = false
		break
	}
	return
}

func severityTextToInt(text string) (s Severity, ok bool) {
	ok = true
	switch text {
	case debugText:
		s = Debug
		break
	case errorText:
		s = Error
		break
	case fatalText:
		s = Fatal
		break
	case infoText:
		s = Info
		break
	case traceText:
		s = Trace
		break
	case warnText:
		s = Warn
		break
	default:
		ok = false
		break
	}
	return
}
