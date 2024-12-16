package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 定义自定义的 TraceLevel
const TraceLevel zapcore.Level = -2

// Logger 封装日志记录器
type Logger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

// NewLogger 创建一个新的 Logger 实例
func NewLogger() *Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "",
		MessageKey:     "message",
		StacktraceKey:  "",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   nil,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		TraceLevel, // 设置最低日志级别为自定义的 TraceLevel
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	sugar := logger.Sugar()

	return &Logger{logger: logger, sugar: sugar}
}

// logWithCaller 格式化日志记录逻辑
func (l *Logger) logWithCaller(level zapcore.Level, format string, args []interface{}, fields ...zap.Field) {
	caller := getCaller()
	msg := fmt.Sprintf(format, args...)
	// 创建一个字段切片用于添加日志字段
	var allFields []zap.Field
	allFields = append(allFields, zap.Object("casa", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
		enc.AddObject("src", zapcore.ObjectMarshalerFunc(func(srcEnc zapcore.ObjectEncoder) error {
			srcEnc.AddString("file", caller.File)
			srcEnc.AddString("function", caller.Function)
			srcEnc.AddInt("line", caller.Line)
			return nil
		}))
		return nil
	})))
	// 添加用户传递的字段
	allFields = append(allFields, fields...)

	// 根据日志级别调用相应的方法
	switch level {
	case TraceLevel:
		l.logger.Debug(msg, allFields...) // 使用 Debug 方法来记录 Trace 日志
	case zapcore.DebugLevel:
		l.logger.Debug(msg, allFields...)
	case zapcore.InfoLevel:
		l.logger.Info(msg, allFields...)
	case zapcore.WarnLevel:
		l.logger.Warn(msg, allFields...)
	case zapcore.ErrorLevel:
		l.logger.Error(msg, allFields...)
	case zapcore.FatalLevel:
		l.logger.Fatal(msg, allFields...)
	case zapcore.PanicLevel:
		l.logger.Panic(msg, allFields...)
	}
}

// Tracef 记录 trace 级别的格式化日志
func (l *Logger) Tracef(format string, args ...interface{}) {
	l.logWithCaller(TraceLevel, format, args)
}

// Infof 记录 info 级别的格式化日志
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logWithCaller(zapcore.InfoLevel, format, args)
}

// Warnf 记录 warn 级别的格式化日志
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logWithCaller(zapcore.WarnLevel, format, args)
}

// Errorf 记录 error 级别的格式化日志
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logWithCaller(zapcore.ErrorLevel, format, args)
}

// Debugf 记录 debug 级别的格式化日志
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logWithCaller(zapcore.DebugLevel, format, args)
}

// Fatalf 记录 fatal 级别的格式化日志并终止程序
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logWithCaller(zapcore.FatalLevel, format, args)
	os.Exit(1)
}

// Panicf 记录 panic 级别的格式化日志并触发 panic
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logWithCaller(zapcore.PanicLevel, format, args)
}

// getCaller 获取调用者信息
func getCaller() zapcore.EntryCaller {
	if pc, file, line, ok := runtime.Caller(3); ok {
		function := runtime.FuncForPC(pc).Name()
		return zapcore.EntryCaller{
			File:     file,
			Function: function,
			Line:     line,
		}
	}
	return zapcore.EntryCaller{}
}

func main() {
	logger := NewLogger()
	defer logger.logger.Sync()
	<-time.After(10 * time.Second)
	// 使用封装的方法记录不同级别的格式化日志
	logger.Tracef("this is a trace message: %s", "details")
	logger.Infof("this is an info message: %s", "details")
	logger.Warnf("this is a warning message: %d", 123)
	logger.Errorf("this is an error message: %v", map[string]int{"key": 1})
	logger.Debugf("this is a debug message: %.2f", 3.1415)
	// Fatalf 和 Panicf 会终止程序或触发 panic
	// logger.Fatalf("this is a fatal message: %s", "fatal error")
	// logger.Panicf("this is a panic message: %s", "panic error")
	tb := false
	InitTest(tb)
}

func InitTest(b bool) {
	fmt.Println(b)
}
