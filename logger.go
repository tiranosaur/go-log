package Logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
)

const placeholder = "{}"

func initLog(logPtr **log.Logger, prefix string) {
	if *logPtr == nil {
		*logPtr = log.New(os.Stdout, prefix, 0)
	}
}

func replacePlaceholders(str string, args ...interface{}) string {
	for _, arg := range args {
		str = strings.Replace(str, placeholder, fmt.Sprint(arg), 1)
	}
	return str
}

func WARN(str string, args ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	initLog(&warningLog, getCallerString("WARN: ", filePath, line))
	warningLog.Println(replacePlaceholders(str, args...))
}

func INFO(str string, args ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	initLog(&infoLog, getCallerString("INFO: ", filePath, line))
	infoLog.Println(replacePlaceholders(str, args...))
}

func ERROR(str string, args ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	initLog(&errorLog, getCallerString("ERROR: ", filePath, line))
	errorLog.Println(replacePlaceholders(str, args...))
}

func getCallerString(prefix string, filePath string, line int) string {
	currentTime := time.Now().Format("2006-01-02T15:04:05.999-07:00")
	baseDir, _ := filepath.Abs(filepath.Dir("."))
	rootPath, _ := filepath.Rel(baseDir, filePath)
	return fmt.Sprintf("%s \t %s \t %s:%d \t : ", currentTime, prefix, rootPath, line)
}
