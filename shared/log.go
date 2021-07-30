package shared

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

// Log Creates a log entry
func Log(value string) {
	color.White(value)
}

// LogHighlight Creates a log entry with highlight words
func LogHighlight(format string, values ...string) {
	formatColorString(color.FgWhite, color.FgHiYellow, format, values...)
}

// LogInfo Creates a information log entry
func LogInfo(value string) {
	color.HiCyan(value)
}

// LogInfoHighlight Creates a information log entry with highlight words
func LogInfoHighlight(format string, values ...string) {
	formatColorString(color.FgHiCyan, color.FgHiYellow, format, values...)
}

// LogCmd Creates a command log entry
func LogCmd(value string) {
	color.HiBlue(value)
}

// LogCmdHighlight Creates a command log entry with highlight words
func LogCmdHighlight(format string, values ...string) {
	formatColorString(color.FgHiBlue, color.FgHiYellow, format, values...)
}

// LogError Creates an error log entry
func LogError(value string) {
	color.HiRed(value)
}

// LogErrorHighlight Creates an error log entry with highlight words
func LogErrorHighlight(format string, values ...string) {
	formatColorString(color.FgHiRed, color.FgHiBlack, format, values...)
}

// LogWarning Creates a warning log entry
func LogWarning(value string) {
	color.Yellow(value)
}

// LogWarnHighlight Creates a warning log entry with highlight words
func LogWarnHighlight(format string, values ...string) {
	formatColorString(color.FgYellow, color.FgHiWhite, format, values...)
}

// LogSuccess Creates a success log entry
func LogSuccess(value string) {
	color.HiGreen(value)
}

// LogSuccessHighlight Creates a success log entry with highlight words
func LogSuccessHighlight(format string, values ...string) {
	formatColorString(color.FgHiGreen, color.FgHiYellow, format, values...)
}

// LogDebug Creates a debug log entry
func LogDebug(value string) {
	debug := os.Getenv("debug")
	if debug == "true" {
		color.Magenta(value)
	}
}

// LogTrace Creates a trace log entry
func LogTrace(value string) {
	debug := os.Getenv("debug")
	if debug == "true" {
		color.HiMagenta(value)
	}
}

// Highlight Highlights a string
func Highlight(value string) string {
	highlight := color.New(color.FgHiYellow)
	highlight = highlight.Add(color.Bold)
	return highlight.Sprint(value)
}

// HighlightRed Highlights with red a string
func HighlightRed(value string) string {
	highlight := color.New(color.FgRed)
	highlight = highlight.Add(color.Bold)
	return highlight.Sprint(value)
}

func formatColorString(colorName color.Attribute, highlightColorName color.Attribute, format string, values ...string) {
	text := color.New(colorName)
	highlight := color.New(highlightColorName)
	highlight = highlight.Add(color.Bold)
	var words []string
	for _, word := range values {
		words = append(words, highlight.Sprint(word))
	}
	formatWords := strings.Split(format, "%v")
	var st string
	i := 0
	for idx, w := range formatWords {
		st += text.Sprint(w)
		if i <= len(words)-1 {
			st += words[idx]
			i++
		}
	}
	color.White(st)
}
