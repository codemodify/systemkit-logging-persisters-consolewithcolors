package persisters

import (
	logging "github.com/codemodify/systemkit-logging"
)

// ConsoleLoggerColorDelegate -
type ConsoleLoggerColorDelegate func(string, ...interface{}) string

// BlackStringYellowBG -
func BlackStringYellowBG(format string, a ...interface{}) string {
	c := New(FgBlack, BgYellow)
	return c.Sprintf(format, a...)
}

// BlackStringWhiteBG -
func BlackStringWhiteBG(format string, a ...interface{}) string {
	c := New(FgBlack, BgWhite)
	return c.Sprintf(format, a...)
}

// NewDefaultColors -
func NewDefaultColors() map[logging.LogType]ConsoleLoggerColorDelegate {
	return map[logging.LogType]ConsoleLoggerColorDelegate{
		logging.TypeDisable: WhiteString,
		logging.TypeTrace:   BlackStringYellowBG,
		logging.TypePanic:   RedString,
		logging.TypeFatal:   RedString,
		logging.TypeError:   RedString,
		logging.TypeWarning: YellowString,
		logging.TypeInfo:    WhiteString,
		logging.TypeSuccess: GreenString,
		logging.TypeDebug:   CyanString,
	}
}

// NewConsoleLoggerCustomColors -
func NewConsoleLoggerCustomColors(colors map[logging.LogType]ConsoleLoggerColorDelegate) logging.CoreLogger {
	return NewConsoleLogger(colors)
}

// NewConsoleLoggerDefaultColors -
func NewConsoleLoggerDefaultColors() logging.CoreLogger {
	return NewConsoleLoggerCustomColors(NewDefaultColors())
}

// NewConsoleLoggerCustomColorsEasy -
func NewConsoleLoggerCustomColorsEasy(colors map[logging.LogType]ConsoleLoggerColorDelegate) logging.Logger {
	return logging.NewLoggerImplementation(NewConsoleLoggerCustomColors(colors))
}

// NewConsoleLoggerDefaultColorsEasy -
func NewConsoleLoggerDefaultColorsEasy() logging.Logger {
	return NewConsoleLoggerCustomColorsEasy(NewDefaultColors())
}
