package Logger

import (
	"fmt"
	"log"
	"os"
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
		*logPtr = log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Llongfile)
	}
}

func replacePlaceholders(str string, args ...interface{}) string {
	for _, arg := range args {
		str = strings.Replace(str, placeholder, fmt.Sprint(arg), 1)
	}
	return str
}

func WARN(str string, args ...interface{}) {
	initLog(&warningLog, "WARN: ")
	warningLog.Println(replacePlaceholders(str, args...))
}

func INFO(str string, args ...interface{}) {
	initLog(&infoLog, "INFO: ")
	infoLog.Println(replacePlaceholders(str, args...))
}

func ERROR(str string, args ...interface{}) {
	initLog(&errorLog, "ERROR: ")
	errorLog.Println(replacePlaceholders(str, args...))
}
