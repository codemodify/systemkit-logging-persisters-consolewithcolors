package persisters

import (
	"fmt"

	logging "github.com/codemodify/systemkit-logging"
	colors "github.com/codemodify/systemkit-terminal-colors"
)

// ConsoleLoggerColorDelegate -
type ConsoleLoggerColorDelegate func(string, ...interface{}) string

func whiteString(format string, a ...interface{}) string {
	return colors.ColoredText(fmt.Sprintf(format, a), colors.Foreground.White).String()
}

func blackStringYellowBG(format string, a ...interface{}) string {
	return colors.ColoredTextWithBG(fmt.Sprintf(format, a), colors.Foreground.Black, colors.Background.LightYellow).String()
}

func redString(format string, a ...interface{}) string {
	return colors.ColoredText(fmt.Sprintf(format, a), colors.Foreground.LightRed).String()
}

func yellowString(format string, a ...interface{}) string {
	return colors.ColoredText(fmt.Sprintf(format, a), colors.Foreground.LightYellow).String()
}

func greenString(format string, a ...interface{}) string {
	return colors.ColoredText(fmt.Sprintf(format, a), colors.Foreground.LightGreen).String()
}

func cyanString(format string, a ...interface{}) string {
	return colors.ColoredText(fmt.Sprintf(format, a), colors.Foreground.Cyan).String()
}

// NewDefaultColors -
func NewDefaultColors() map[logging.LogType]ConsoleLoggerColorDelegate {
	return map[logging.LogType]ConsoleLoggerColorDelegate{
		logging.TypeDisable: whiteString,
		logging.TypeTrace:   blackStringYellowBG,
		logging.TypePanic:   redString,
		logging.TypeFatal:   redString,
		logging.TypeError:   redString,
		logging.TypeWarning: yellowString,
		logging.TypeInfo:    whiteString,
		logging.TypeSuccess: greenString,
		logging.TypeDebug:   cyanString,
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
