package middleware

import (
	"fmt"
	"os"
	"time"
)

var (
	logFilePath = "/var/log/vpn.log"
)

type Logger struct{}

func (l *Logger) LogErr(err error) {
	if err == nil {
		return
	}
	logText := fmt.Sprintf("[Error] %s %v", time.Now().UTC(), err)
	l.LogToFile(logText)
	fmt.Println(logText)
}

func (l *Logger) LogInfo(info string) {
	if info == "" {
		return
	}
	logText := fmt.Sprintf("[Info] %v", info)
	l.LogToFile(logText)
	fmt.Println(logText)
}

func (l *Logger) LogDebug(debug string) {
	logText := fmt.Sprintf("[Debug] %v", debug)
	fmt.Println(logText)
}

func (l *Logger) LogToFile(data string) {
	f, _ := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	_, _ = f.WriteString(data)
}
