package log

import (
	"fmt"
	"ittech24/rest/apidemo/bootstrap"
	"ittech24/rest/apidemo/strcolor"
	"os"
	"strings"

	"github.com/fatih/color"
)

// CmdLogger Command Line Logger implementation
type CmdLogger struct{}

// Logger Ansi Colors
const (
	SuccessColor  = color.FgGreen
	InfoColor     = color.FgHiWhite
	NoticeColor   = color.FgHiCyan
	WarningColor  = color.FgYellow
	ErrorColor    = color.FgRed
	DebugColor    = color.FgMagenta
	TraceColor    = color.FgHiMagenta
	CommandColor  = color.FgBlue
	DisabledColor = color.FgHiBlack
)

// Log Log information message
func (l *CmdLogger) Log(format string, level Level, words ...string) {
	switch level {
	case 0:
		printMessage(format, "error", false, false, words...)
	case 1:
		printMessage(format, "warn", false, false, words...)
	case 2:
		printMessage(format, "info", false, false, words...)
	case 3:
		printMessage(format, "debug", false, false, words...)
	case 4:
		printMessage(format, "trace", false, false, words...)
	}
}

// LogHighlight Log information message
func (l *CmdLogger) LogHighlight(format string, level Level, highlightColor strcolor.ColorCode, words ...string) {
	for i := range words {
		words[i] = strcolor.GetColorString(strcolor.ColorCode(highlightColor), words[i])
	}

	switch level {
	case 0:
		printMessage(format, "error", false, false, words...)
	case 1:
		printMessage(format, "warn", false, false, words...)
	case 2:
		printMessage(format, "info", false, false, words...)
	case 3:
		printMessage(format, "debug", false, false, words...)
	case 4:
		printMessage(format, "trace", false, false, words...)
	}
}

// Info log information message
func (l *CmdLogger) Info(format string, words ...string) {
	printMessage(format, "info", false, false, words...)
}

// Success log message
func (l *CmdLogger) Success(format string, words ...string) {
	printMessage(format, "success", false, false, words...)
}

// TaskSuccess log message
func (l *CmdLogger) TaskSuccess(format string, isComplete bool, words ...string) {
	printMessage(format, "success", true, isComplete, words...)
}

// Warn log message
func (l *CmdLogger) Warn(format string, words ...string) {
	printMessage(format, "warn", false, false, words...)
}

// TaskWarn log message
func (l *CmdLogger) TaskWarn(format string, words ...string) {
	printMessage(format, "warn", true, false, words...)
}

// Command log message
func (l *CmdLogger) Command(format string, words ...string) {
	printMessage(format, "command", false, false, words...)
}

// Disabled log message
func (l *CmdLogger) Disabled(format string, words ...string) {
	printMessage(format, "disabled", false, false, words...)
}

// Notice log message
func (l *CmdLogger) Notice(format string, words ...string) {
	printMessage(format, "notice", false, false, words...)
}

// Debug log message
func (l *CmdLogger) Debug(format string, words ...string) {
	printMessage(format, "debug", false, false, words...)
}

// Trace log message
func (l *CmdLogger) Trace(format string, words ...string) {
	printMessage(format, "trace", false, false, words...)
}

// Error log message
func (l *CmdLogger) Error(format string, words ...string) {
	printMessage(format, "error", false, false, words...)
}

// LogError log message
func (l *CmdLogger) LogError(message error) {
	if message != nil {
		printMessage(message.Error(), "error", false, false)
	}
}

// TaskError log message
func (l *CmdLogger) TaskError(format string, isComplete bool, words ...string) {
	printMessage(format, "error", true, isComplete, words...)
}

// Fatal log message
func (l *CmdLogger) Fatal(format string, words ...string) {
	printMessage(format, "error", false, true, words...)
}

// FatalError log message
func (l *CmdLogger) FatalError(e error, format string, words ...string) {
	l.Error(format, words...)
	if e != nil {
		panic(e)
	}
}

// printMessage Prints a message in the system
func printMessage(format string, level string, isTask bool, isComplete bool, words ...string) {
	agentID := os.Getenv("AGENT_ID")
	isPipeline := false
	if len(agentID) != 0 {
		isPipeline = true
	}

	if !isPipeline {
		format = format + "\u001b[0m" + "\n"
	} else {
		if (level == "warn" || level == "error") && isTask {
			format = format + "\n"
		} else {
			format = format + "\033[0m" + "\n"
		}
	}

	successWriter := color.New(SuccessColor).PrintfFunc()
	warningWriter := color.New(WarningColor).PrintfFunc()
	errorWriter := color.New(ErrorColor).PrintfFunc()
	debugWriter := color.New(DebugColor).PrintfFunc()
	traceWriter := color.New(TraceColor).PrintfFunc()
	infoWriter := color.New(InfoColor).PrintfFunc()
	noticeWriter := color.New(NoticeColor).PrintfFunc()
	commandWriter := color.New(CommandColor).PrintfFunc()
	disableWriter := color.New(DisabledColor).PrintfFunc()

	formatedWords := make([]interface{}, len(words))
	for i := range words {
		words[i] = strings.TrimSpace(words[i])
		words[i] = strings.ReplaceAll(words[i], "\n\n", "\n")
		if words[i][0] == 27 {
			switch strings.ToLower(level) {
			case "success":
				if isPipeline {
					words[i] += "\033[" + fmt.Sprint(SuccessColor) + "m"
				} else {
					words[i] += "\u001b[" + fmt.Sprint(SuccessColor) + "m"
				}
			case "warn":
				if isPipeline {
					if !isTask {
						words[i] += "\033[" + fmt.Sprint(WarningColor) + "m"
					}
				} else {
					words[i] += "\u001b[" + fmt.Sprint(WarningColor) + "m"
				}
			case "error":
				if isPipeline {
					if !isTask {
						words[i] += "\033[" + fmt.Sprint(ErrorColor) + "m"
					}
				} else {
					words[i] += "\u001b[" + fmt.Sprint(ErrorColor) + "m"
				}
			case "debug":
				if isPipeline {
					words[i] += "\033[" + fmt.Sprint(DebugColor) + "m"
				} else {
					words[i] += "\u001b[" + fmt.Sprint(DebugColor) + "m"
				}
			case "trace":
				if isPipeline {
					words[i] += "\033[" + fmt.Sprint(TraceColor) + "m"
				} else {
					words[i] += "\u001b[" + fmt.Sprint(TraceColor) + "m"
				}
			case "info":
				if isPipeline {
					words[i] += "\033[" + fmt.Sprint(InfoColor) + "m"
				} else {
					words[i] += "\u001b[" + fmt.Sprint(InfoColor) + "m"
				}
			case "notice":
				if isPipeline {
					words[i] += "\033[" + fmt.Sprint(NoticeColor) + "m"
				} else {
					words[i] += "\u001b[" + fmt.Sprint(NoticeColor) + "m"
				}
			case "command":
				if isPipeline {
					words[i] += "\033[" + fmt.Sprint(CommandColor) + "m"
				} else {
					words[i] += "\u001b[" + fmt.Sprint(CommandColor) + "m"
				}
			case "disabled":
				if isPipeline {
					words[i] += "\033[" + fmt.Sprint(DisabledColor) + "m"
				} else {
					words[i] += "\u001b[" + fmt.Sprint(DisabledColor) + "m"
				}
			}
			formatedWords[i] = words[i]
		} else {
			formatedWords[i] = words[i]
		}
	}

	switch strings.ToLower(level) {
	case "success":
		if isPipeline {
			format = "\033[" + fmt.Sprint(SuccessColor) + "m" + format
			format = "##[section]" + format
			fmt.Printf(format, formatedWords...)
			if isTask && isComplete {
				fmt.Printf("\033[" + fmt.Sprint(SuccessColor) + "m" + "##vso[task.complete result=Succeeded;]\n")
			}
		} else {
			successWriter(format, formatedWords...)
		}

		if isComplete {
			if isPipeline && isTask {
				fmt.Printf("\033[" + fmt.Sprint(SuccessColor) + "m" + "##[section] Completed\n")
			} else {
				successWriter("Completed")
			}
			bootstrap.Exit(0)
		}
	case "warn":
		if isPipeline {
			if isTask {
				format = "##vso[task.LogIssue type=warning;]" + format
				fmt.Printf(format, formatedWords...)
			} else {
				format = "\033[" + fmt.Sprint(WarningColor) + "m" + format
				fmt.Printf(format, formatedWords...)
			}
		} else {
			warningWriter(format, formatedWords...)
		}
	case "error":
		if isPipeline {
			if isTask {
				format = "##vso[task.LogIssue type=error;]" + format
				fmt.Printf(format, formatedWords...)
			} else {
				format = "\033[" + fmt.Sprint(ErrorColor) + "m" + format
				fmt.Printf(format, formatedWords...)
			}
		} else {
			errorWriter(format, formatedWords...)
		}

		if isComplete {
			if isPipeline && isTask {
				format = "\033[" + fmt.Sprint(ErrorColor) + "m" + format
				fmt.Printf("##vso[task.complete result=Failed;]\n")
				bootstrap.Exit(0)
			} else {
				errorWriter("Failed\n")
				bootstrap.Exit(1)
			}
		}
	case "debug":
		if isPipeline {
			format = "\033[" + fmt.Sprint(DebugColor) + "m" + format
			fmt.Printf(format, formatedWords...)
		} else {
			debugWriter(format, formatedWords...)
		}
	case "trace":
		if isPipeline {
			format = "\033[" + fmt.Sprint(TraceColor) + "m" + format
			fmt.Printf(format, formatedWords...)
		} else {
			traceWriter(format, formatedWords...)
		}
	case "info":
		if isPipeline {
			format = "\033[" + fmt.Sprint(InfoColor) + "m" + format
			fmt.Printf(format, formatedWords...)
		} else {
			infoWriter(format, formatedWords...)
		}
	case "notice":
		if isPipeline {
			format = "\033[" + fmt.Sprint(NoticeColor) + "m" + format
			fmt.Printf(format, formatedWords...)
		} else {
			noticeWriter(format, formatedWords...)
		}
	case "command":
		if isPipeline {
			format = "\033[" + fmt.Sprint(CommandColor) + "m" + format
			format = "##[command]" + format
			fmt.Printf(format, formatedWords...)
		} else {
			commandWriter(format, formatedWords...)
		}
	case "disabled":
		if isPipeline {
			format = "\033[" + fmt.Sprint(DisabledColor) + "m" + format
			fmt.Printf(format, formatedWords...)
		} else {
			disableWriter(format, formatedWords...)
		}
	}
}
