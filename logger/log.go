package logger

import (
	"fmt"
	"log"
)

var Log *_log

type _log struct {
	level int
	l     *log.Logger
}

const (
	Debug = iota
	Info
	Silent
)

func Setup(level int, out *log.Logger) {
	if level > 2 || level < 0 {
		panic(fmt.Errorf("invalid log level:%v, try to use logger.Debug etc", level))
	}

	Log = &_log{
		level: level,
		l:     out,
	}
}

func (L *_log) Infoln(a ...any) {
	if L.level <= Info {
		L.l.Print("INFO: ")
		L.l.Println(a...)
	}
}

func (L *_log) Infof(format string, a ...any) {
	if L.level <= Info {
		L.l.Print("INFO: ")
		L.l.Printf(format, a...)
	}
}

func (L *_log) Warnln(a ...any) {
	if L.level <= Silent {
		L.l.Print("WARN: ")
		L.l.Println(a...)
	}
}

func (L *_log) Warnf(format string, a ...any) {
	if L.level <= Silent {
		L.l.Print("WARN: ")
		L.l.Printf(format, a...)
	}
}

func (L *_log) Debugln(a ...any) {
	if L.level <= Debug {
		L.l.Print("DEBUG: ")
		L.l.Println(a...)
	}
}

func (L *_log) Debugf(format string, a ...any) {
	if L.level <= Debug {
		L.l.Print("DEBUG: ")
		L.l.Printf(format, a...)
	}
}
