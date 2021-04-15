package logger

const (
	// Fatal error/panic, not recoverable
	Fatal      Severity = 0
	fatalText  string   = "FATAL"
	fatalColor uint16   = 31 // red

	// Error encountered, possibly recoverable
	Error      Severity = 1
	errorText  string   = "ERROR"
	errorColor uint16   = 31 // red

	// Warn of potential operating concerns or failure trajectory
	Warn      Severity = 2
	warnText  string   = "WARN"
	warnColor uint16   = 33 // yellow

	// Info updates
	Info      Severity = 3
	infoText  string   = "INFO"
	infoColor uint16   = 36 // blue

	// Debug data. Solely used in development
	Debug      Severity = 4
	debugText  string   = "DEBUG"
	debugColor uint16   = 37 // grey

	// Trace data. Solely used in development
	Trace      Severity = 5
	traceText  string   = "TRACE"
	traceColor uint16   = 37 // grey
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

// GetSeverityColors returns a dictionary of severity colors indexed by their internal values.
func GetSeverityColors() map[Severity]uint16 {
	return map[Severity]uint16{
		Fatal: fatalColor,
		Error: errorColor,
		Warn:  warnColor,
		Info:  infoColor,
		Debug: debugColor,
		Trace: traceColor,
	}
}

func severityIntToColor(s Severity) uint16 {
	switch s {
	case Debug:
		return debugColor
	case Error:
		return errorColor
	case Fatal:
		return fatalColor
	case Info:
		return infoColor
	case Trace:
		return traceColor
	case Warn:
		return warnColor
	default:
		return debugColor
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
