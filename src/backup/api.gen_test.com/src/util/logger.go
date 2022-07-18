package util

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/andodeki/api.gen_test.com/src/config"
	
	// "github.com/andodeki/api.propertylist.com/util"
	"github.com/jackc/pgx/v4"
	gommon "github.com/labstack/gommon/log"

	// "github.com/neoxelox/odin/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

var (
	ZlevelToGlevel = map[zerolog.Level]gommon.Lvl{
		zerolog.DebugLevel: gommon.DEBUG,
		zerolog.InfoLevel:  gommon.INFO,
		zerolog.WarnLevel:  gommon.WARN,
		zerolog.ErrorLevel: gommon.ERROR,
		zerolog.Disabled:   gommon.OFF,
	}

	GlevelToZlevel = map[gommon.Lvl]zerolog.Level{
		gommon.DEBUG: zerolog.DebugLevel,
		gommon.INFO:  zerolog.InfoLevel,
		gommon.WARN:  zerolog.WarnLevel,
		gommon.ERROR: zerolog.ErrorLevel,
		gommon.OFF:   zerolog.Disabled,
	}

	ZlevelToPlevel = map[zerolog.Level]pgx.LogLevel{
		zerolog.TraceLevel: pgx.LogLevelTrace,
		zerolog.DebugLevel: pgx.LogLevelDebug,
		zerolog.InfoLevel:  pgx.LogLevelInfo,
		zerolog.WarnLevel:  pgx.LogLevelWarn,
		zerolog.ErrorLevel: pgx.LogLevelError,
		zerolog.Disabled:   pgx.LogLevelNone,
	}

	PlevelToZlevel = map[pgx.LogLevel]zerolog.Level{
		pgx.LogLevelTrace: zerolog.TraceLevel,
		pgx.LogLevelDebug: zerolog.DebugLevel,
		pgx.LogLevelInfo:  zerolog.InfoLevel,
		pgx.LogLevelWarn:  zerolog.WarnLevel,
		pgx.LogLevelError: zerolog.ErrorLevel,
		pgx.LogLevelNone:  zerolog.Disabled,
	}
)

type Logger struct {
	configuration config.Config
	logger        zerolog.Logger
	level         zerolog.Level
	out           io.Writer
	prefix        string
	header        string
	verbose       bool
}

func NewLogger(configuration config.Config) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "timestamp"
	zerolog.CallerSkipFrameCount = 3

	level := zerolog.DebugLevel
	// if "dev" != configuration {
	// 	level = zerolog.InfoLevel
	// }

	out := diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Fprintf(os.Stderr, "Logger dropped %!d(MISSING) messages", missed)
	})

	return &Logger{
		configuration: configuration,
		logger:        zerolog.New(out).With().Str("service", configuration.HttpServer.ServerName()).Timestamp().Logger().Level(level),
		level:         level,
		out:           out,
		prefix:        configuration.HttpServer.ServerName(),
		header:        "",
		verbose:       level == zerolog.DebugLevel,
	}
}

func (logg Logger) Logger() *zerolog.Logger {
	return &logg.logger
}

func (logg *Logger) SetLogger(l zerolog.Logger) {
	logg.logger = l
}

func (logg Logger) Flush() {
	os.Stderr.Sync()
}

func (logg Logger) Close(ctx context.Context) error {
	logg.Flush()
	dw, _ := logg.out.(diode.Writer)
	return dw.Close()
}

func (logg Logger) Output() io.Writer {
	return logg.out
}

func (logg *Logger) SetOutput(w io.Writer) {
	logg.logger = logg.logger.Output(w)
	logg.out = w
}

func (logg Logger) Prefix() string {
	return logg.prefix
}

func (logg *Logger) SetPrefix(p string) {
	logg.prefix = p
}

func (logg Logger) GLevel() gommon.Lvl {
	return ZlevelToGlevel[logg.level]
}

func (logg *Logger) SetGLevel(l gommon.Lvl) {
	zlevel := GlevelToZlevel[l]
	logg.logger = logg.logger.Level(zlevel)
	logg.level = zlevel
}

func (logg Logger) PLevel() pgx.LogLevel {
	return ZlevelToPlevel[logg.level]
}

func (logg *Logger) SetPLevel(l pgx.LogLevel) {
	zlevel := PlevelToZlevel[l]
	logg.logger = logg.logger.Level(zlevel)
	logg.level = zlevel
}

func (logg Logger) ZLevel() zerolog.Level {
	return logg.level
}

func (logg *Logger) SetZLevel(l zerolog.Level) {
	zlevel := l
	logg.logger = logg.logger.Level(zlevel)
	logg.level = zlevel
}

func (logg *Logger) Header() string {
	return logg.header
}

func (logg *Logger) SetHeader(h string) {
	logg.header = h
}

func (logg Logger) Verbose() bool {
	return logg.verbose
}

func (logg Logger) SetVerbose(v bool) {
	logg.verbose = v
}

func (logg Logger) Print(i ...interface{}) {
	logg.logger.Log().Msg(fmt.Sprint(i...))
}

func (logg Logger) Printf(format string, i ...interface{}) {
	logg.logger.Log().Msgf(format, i...)
}

func (logg Logger) Debug(i ...interface{}) {
	logg.logger.Debug().Msg(fmt.Sprint(i...))
}

func (logg Logger) Debugf(format string, i ...interface{}) {
	logg.logger.Debug().Msgf(format, i...)
}

func (logg Logger) Info(i ...interface{}) {
	logg.logger.Info().Msg(fmt.Sprint(i...))
}

func (logg Logger) Infof(format string, i ...interface{}) {
	logg.logger.Info().Msgf(format, i...)
}

func (logg Logger) Warn(i ...interface{}) {
	logg.logger.Warn().Msg(fmt.Sprint(i...))
}

func (logg Logger) Warnf(format string, i ...interface{}) {
	logg.logger.Warn().Msgf(format, i...)
}

func (logg Logger) Error(i ...interface{}) {
	if !logg.configuration.IsDev { //logg.configuration.Environment == internal.Environment.PRODUCTION {
		logg.logger.Error().Msg(fmt.Sprint(i...))
	} else {
		if i != nil {
			if len(i) == 1 {
				err, _ := i[0].(error)
				if ierr, ok := i[0].(*Error); ok {
					err = ierr.Unwrap()
				}
				fmt.Printf("%!v(MISSING)\n", err)
			} else {
				fmt.Printf("%!v(MISSING)\n", i...)
			}
		}
	}
}

func (logg Logger) Errorf(format string, i ...interface{}) {
	logg.logger.Error().Msgf(format, i...)
}

func (logg Logger) Fatal(i ...interface{}) {
	logg.logger.Fatal().Msg(fmt.Sprint(i...))
}

func (logg Logger) Fatalf(format string, i ...interface{}) {
	logg.logger.Fatal().Msgf(format, i...)
}

func (logg Logger) Panic(i ...interface{}) {
	logg.logger.Panic().Msg(fmt.Sprint(i...))
}

func (logg Logger) Panicf(format string, i ...interface{}) {
	logg.logger.Panic().Msgf(format, i...)
}
