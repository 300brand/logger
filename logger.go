package logger

import (
	"log"
)

const DefaultFlags = log.LstdFlags | log.Lshortfile | log.Lmicroseconds

var (
	Trace = log.New(NewColorWriter("b"), "TRACE ", DefaultFlags)
	Debug = log.New(NewColorWriter("c"), "DEBUG ", DefaultFlags)
	Info  = log.New(NewColorWriter("g"), "INFO  ", DefaultFlags)
	Warn  = log.New(NewColorWriter("y"), "WARN  ", DefaultFlags)
	Error = log.New(NewColorWriter("r"), "ERROR ", DefaultFlags)
)
