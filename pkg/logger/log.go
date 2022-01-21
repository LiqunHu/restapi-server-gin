package logger

import (
	"io"
	"strconv"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/LiqunHu/restapi-server-gin/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Setup() {
	var cores []zapcore.Core
	if setting.AppSetting.DebugFlag {
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	} else {
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, highPriority))
	}

	if setting.AppSetting.DebugFlag {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(getWriter(setting.AppSetting.LogSavePath+setting.AppSetting.AppName+".log")), lowPriority))
	} else {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(getWriter(setting.AppSetting.LogSavePath+setting.AppSetting.AppName+".log")), highPriority))
	}

	// if c.MultiFile {
	// 	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(getWriter(setting.LogPath+c.AppName+"_info.log")), infoLevel))
	// 	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(getWriter(setting.LogPath+c.AppName+"_warn.log")), warnLevel))
	// 	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(getWriter(setting.LogPath+c.AppName+"_error.log")), errorLevel))
	// 	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(getWriter(setting.LogPath+c.AppName+"_fatal.log")), fatalLevel))
	// 	if c.Debug {
	// 		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(getWriter(c.LogPath+c.AppName+"_debug.log")), debugLevel))
	// 	}
	// } else {

	// }
	setLogger(zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.Development(),
		zap.AddCallerSkip(1),
	).Sugar())
}

// TimeEncoder time encoder .
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strconv.FormatInt(time.Now().Unix(), 10))
}

// Debugf log
func Debugf(msg string, args ...interface{}) {
	Logger().Debugf(msg, args...)
}

// Debug log
func Debug(args ...interface{}) {
	Logger().Debug(args...)
}

// Infof log
func Infof(msg string, args ...interface{}) {
	Logger().Infof(msg, args...)
}

// Info log
func Info(args ...interface{}) {
	Logger().Info(args...)
}

// Errorf log
func Errorf(msg string, args ...interface{}) {
	Logger().Errorf(msg, args...)
}

// Error log
func Error(args ...interface{}) {
	Logger().Error(args...)
}

// Warnf log
func Warnf(msg string, args ...interface{}) {
	Logger().Warnf(msg, args...)
}

// Warn log
func Warn(args ...interface{}) {
	Logger().Warn(args...)
}

// Fatalf send log fatalf
func Fatalf(msg string, args ...interface{}) {
	Logger().Fatalf(msg, args...)
}

// Fatal send log fatal
func Fatal(args ...interface{}) {
	Logger().Fatal(args...)
}

func getWriter(filename string) io.Writer {
	hook := lumberjack.Logger{
		Filename:   filename,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// hook, err := rotatelogs.New(
	// 	strings.Replace(filename, ".log", "", -1)+"_%Y%m%d%H.log",
	// 	rotatelogs.WithLinkName(filename),
	// 	rotatelogs.WithMaxAge(time.Hour*24*7),  // 默认保存时间为7天
	// 	rotatelogs.WithRotationTime(time.Hour), // 每小时滚动一次存储
	// )
	// if err != nil {
	// 	panic(err)
	// }
	return &hook
}

func Logger() *zap.SugaredLogger {
	return (*zap.SugaredLogger)(atomic.LoadPointer(&loggerImpl))
}

func setLogger(l *zap.SugaredLogger) {
	atomic.StorePointer(&loggerImpl, unsafe.Pointer(l))
}

var loggerImpl unsafe.Pointer = unsafe.Pointer(
	zap.New(zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority)),
		zap.AddCaller(),
		zap.Development(),
		zap.AddCallerSkip(1),
	).Sugar())
