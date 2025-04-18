package logx

import (
	"fmt"
	"log"
)

// Configurable flags
var (
	IsDebug = true
	Prefix  = "" // e.g., "[Server]"
)

// Color constants
const (
	colorReset   = "\033[0m"
	colorGreen   = "\033[32m"   // Success
	colorYellow  = "\033[33m"   // Warning
	colorMagenta = "\033[35m"   // Info
	colorRed     = "\033[31m"   // Error
	colorBoldRed = "\033[1;31m" // Fatal
)

// Core formatter
func logWithColor(color, label string, format string, args ...any) {
	if Prefix != "" {
		Prefix += " " // Ensure a space between prefix and log label
	}
	log.Printf("%s%s[%s]%s\t%s\n", color, Prefix, label, colorReset, fmt.Sprintf(format, args...))
}

// Logging functions
func Info(format string, args ...any) {
	if IsDebug {
		logWithColor(colorMagenta, "INFO", format, args...)
	}
}

func Success(format string, args ...any) {
	if IsDebug {
		logWithColor(colorGreen, "SUCCESS", format, args...)
	}
}

func Warning(format string, args ...any) {
	if IsDebug {
		logWithColor(colorYellow, "WARNING", format, args...)
	}
}

func Error(format string, args ...any) {
	logWithColor(colorRed, "ERROR", format, args...)
}

func Fatal(format string, args ...any) {
	log.Fatalf("%s%s[FATAL]%s\t%s\n", colorBoldRed, Prefix, colorReset, fmt.Sprintf(format, args...))
}
