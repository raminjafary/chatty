package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// var contextLogger *logrus.Entry

func GetLoggingEnv() string {
	checkRunningEnv := os.Getenv("HEX_ARCH_ENV")
	if checkRunningEnv == "release" {
		return "structured"
	}
	return "stdout"
}

// SetupLogger setup logger.Log
func SetupLogger() {
	Log = CreateLogger()
}

func LogWithFields(fields interface{}) (*logrus.Entry, error) {
	data, err := StructToMap(fields)

	if err != nil {
		log.Fatal("LogWithFields failed!")
		return nil, err
	}

	contextLogger := Log.WithFields(data)
	return contextLogger, nil
}

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}

// CreateLogger creates a logger instance with the configuration
func CreateLogger() *logrus.Logger {
	logInstance := logrus.New()
	logInstance.SetOutput(io.MultiWriter(os.Stdout))
	logInstance.SetReportCaller(true)

	if GetLoggingEnv() == "structured" {
		logInstance.SetLevel(logrus.ErrorLevel)
		logInstance.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		// if not in production, then default to plain stdout DebugLevel please
		// this may mean slowdown in UAT/QA but overall, it's probably worth it
		logInstance.SetLevel(logrus.DebugLevel)
		logInstance.SetFormatter(&myFormatter{logrus.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
			DisableLevelTruncation: true,
			DisableColors:          false,
		}})
	}
	return logInstance
}

type myFormatter struct {
	logrus.TextFormatter
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	strList := strings.Split(entry.Caller.File, "/")
	fileName := strList[len(strList)-1]

	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 31 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("[%s] - %s - [line:%d] - \x1b[%dm%s\x1b[0m - %s. Data: %v\n", entry.Time.Format(f.TimestampFormat), fileName, entry.Caller.Line, levelColor,
		strings.ToUpper(entry.Level.String()), entry.Message, entry.Data)), nil
}
