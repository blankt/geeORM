package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

//简单的记录log

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	logs     = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range logs {
		logger.SetOutput(os.Stdout)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
}
