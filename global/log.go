package global

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var _logger *zap.Logger

func init() {
	// 获取配置文件
	abs, err := filepath.Abs(_config.Log.LogDir)
	if err != nil {
		panic(err)
	}

	// 创建日志文件夹
	_, err = os.Stat(abs)
	if os.IsNotExist(err) {
		err = os.MkdirAll(abs, 0777)
		if err != nil {
			panic(err)
		}
	}

	// 持久化info
	infoHook, err := rotatelogs.New(
		abs+"/info-%Y%m%d.log",
		rotatelogs.WithLinkName(abs+"/info.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	// 持久化error
	errHook, err := rotatelogs.New(
		abs+"/error-%Y%m%d.log",
		rotatelogs.WithLinkName(abs+"/error.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写不加颜色编码器,
		EncodeTime:     TimeFormat,                     // ISO8601 UTC 时间格式,
		EncodeDuration: zapcore.SecondsDurationEncoder, //,
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.Level(_config.Log.LogLevel))

	infoCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                  // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(infoHook)), // 打印到文件
		atomicLevel, // 日志级别
	)

	errCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(errHook)),
		zap.ErrorLevel,
	)
	core := zapcore.NewTee(infoCore, errCore)
	// 构造日志
	_logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// TimeFormat 时间格式化
func TimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000 -0700"))
}

func Fatal(msg string, err error) {
	_logger.Fatal(msg, zap.Error(err))
}
func Error(msg string, err error) {
	_logger.Error(msg, zap.Error(err))
}
func Debug(msg string, fields ...zapcore.Field) {
	_logger.Debug(msg, fields...)
}
func Info(msg string, fields ...zapcore.Field) {
	_logger.Info(msg, fields...)
}
