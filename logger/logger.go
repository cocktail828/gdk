package logger

import (
	"log"
	"os"

	"github.com/cocktail828/go-tools/z/stringx"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// EnvGdkMode indicates environment name for gdk work mode.
const EnvGdkMode = "GDK_MODE"

const (
	// DebugMode indicates gdk mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gdk mode is release.
	ReleaseMode = "release"
	// TestMode indicates gdk mode is test.
	TestMode = "test"
)

var (
	modeName = DebugMode
)

func init() {
	log.SetPrefix("[GDK] ")
	mode := os.Getenv(EnvGdkMode)
	if !stringx.Contains([]string{DebugMode, ReleaseMode, TestMode, ""}, mode) {
		log.Fatal("env 'GDK_MODE' should be oneof debug|release|test")
	}
	if mode != "" {
		modeName = mode
	}
	log.Println("gdk work mode:", modeName)
}

func NewLogger(filename string) *slog.Logger {
	switch modeName {
	case DebugMode:
		out, _ := os.OpenFile("/log/server/"+filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		return slog.New(slog.NewJSONHandler(
			out,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			}))
	case TestMode:
		out, _ := os.OpenFile("/log/server/"+filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		return slog.New(slog.NewJSONHandler(
			out,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			}))
	case ReleaseMode:
		return slog.New(slog.NewJSONHandler(
			&lumberjack.Logger{
				Filename:   "/log/server/" + filename,
				MaxSize:    10,
				MaxAge:     7,
				MaxBackups: 10,
				LocalTime:  true,
			},
			&slog.HandlerOptions{
				AddSource: false,
				Level:     slog.LevelWarn,
			}))
	default:
		log.Fatal("gdk mode unknown: " + modeName + " (available mode: debug release test)")
		return nil
	}
}
