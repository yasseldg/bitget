package sLog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const EncodeTimeFormat = "2006.01.02 15:04:05"

var sugaredLogger *zap.SugaredLogger
var atomicLevel zap.AtomicLevel

func init() {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "time",
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.TimeEncoderOfLayout(EnvGet("LogTimeFormat", EncodeTimeFormat)),
	}

	// define default level as debug level
	atomicLevel = zap.NewAtomicLevel()
	err := atomicLevel.UnmarshalText([]byte(EnvGet("LogLevel", "DEBUG")))
	if err != nil {
		atomicLevel.SetLevel(zapcore.DebugLevel)
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, atomicLevel)
	sugaredLogger = zap.New(core).Sugar()

	Info("LOG_Level set to ( %s )", atomicLevel.String())
}

func SetLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}

func Fatal(template string, args ...interface{}) {
	sugaredLogger.Fatalf(template, args...)
}

func Error(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}

func Panic(template string, args ...interface{}) {
	sugaredLogger.Panicf(template, args...)
}

func Warn(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

func Info(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

func Debug(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}

// EnvGet
func EnvGet(env_name, defaults string) string {

	env, ok := os.LookupEnv(env_name)
	if ok {
		return env
	}
	return defaults
}
