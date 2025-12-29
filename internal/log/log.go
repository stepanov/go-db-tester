package log

import "log"

// Simple wrapper around standard log to allow future replacement
func Info(v ...interface{})  { log.Println(v...) }
func Error(v ...interface{}) { log.Println(v...) }
func Fatal(v ...interface{}) { log.Fatalln(v...) }
