package Logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
)

const placeholder = "{}"

func initLog(logPtr **log.Logger, prefix string) {
	if *logPtr == nil {
		*logPtr = log.New(os.Stdout, prefix, log.Ldate|log.Ltime)
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
	initLog(&warningLog, "WARN: "+getCallerString(filePath, line))
	warningLog.Println(replacePlaceholders(str, args...))
}

func INFO(str string, args ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	initLog(&infoLog, "INFO: "+getCallerString(filePath, line))
	infoLog.Println(replacePlaceholders(str, args...))
}

func ERROR(str string, args ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	initLog(&errorLog, "ERROR: "+getCallerString(filePath, line))
	errorLog.Println(replacePlaceholders(str, args...))
}

func getCallerString(filePath string, line int) string {
	baseDir, _ := filepath.Abs(filepath.Dir("."))
	rootPath, _ := filepath.Rel(baseDir, filePath)
	return "\\" + rootPath + ":" + strconv.Itoa(line) + " "
}
