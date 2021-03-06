package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

var l *Logger

type writer struct {
	io.Writer
	timeFormat string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().UTC().Format(w.timeFormat)), b...))
}

const (
	red    = "\033[0;31m"
	green  = "\033[0;32m"
	yellow = "\033[0;33m"
	white  = "\033[0;37m"
	cyan   = "\033[0;36m"
	blue   = "\033[0;34m"
	end    = "\033[0m"
)

func init() {
	l = New()
	l.SetOutput(&writer{os.Stdout, "2006-01-02T15:04:05.999Z"})
}

type Logger struct {
	T *log.Logger
	D *log.Logger
	I *log.Logger
	W *log.Logger
	E *log.Logger
	t bool
	d bool
	i bool
	w bool
	e bool
}

func New() (l *Logger) {
	l = &Logger{
		T: log.New(os.Stdout, "[TRACE] ", log.Lshortfile),
		D: log.New(os.Stdout, "[DEBUG] ", log.Lshortfile),
		I: log.New(os.Stdout, "[INFO]  ", log.Lshortfile),
		W: log.New(os.Stdout, "[WARN]  ", log.Lshortfile),
		E: log.New(os.Stdout, "[ERROR] ", log.Lshortfile),
		t: true,
		d: true,
		i: true,
		w: true,
		e: true,
	}
	if runtime.GOOS == "linux" {
		// l.T.SetPrefix(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " " + green + l.T.Prefix() + end)
		// l.D.SetPrefix(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " " + cyan + l.D.Prefix() + end)
		// l.I.SetPrefix(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " " + blue + l.I.Prefix() + end)
		// l.W.SetPrefix(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " " + yellow + l.W.Prefix() + end)
		// l.E.SetPrefix(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " " + red + l.E.Prefix() + end)
		l.T.SetPrefix(" " + green + l.T.Prefix() + end)
		l.D.SetPrefix(" " + cyan + l.D.Prefix() + end)
		l.I.SetPrefix(" " + blue + l.I.Prefix() + end)
		l.W.SetPrefix(" " + yellow + l.W.Prefix() + end)
		l.E.SetPrefix(" " + red + l.E.Prefix() + end)
	}
	return
}

func SetOutput(w io.Writer) {
	l.SetOutput(w)
}

func SetLevel(s string) {
	l.SetLevel(s)
}


func (l *Logger) SetOutput(w io.Writer) {
	l.T.SetOutput(w)
	l.D.SetOutput(w)
	l.I.SetOutput(w)
	l.W.SetOutput(w)
	l.E.SetOutput(w)
}

func (l *Logger) SetLevel(s string) {
	l.t = true
	l.d = true
	l.i = true
	l.w = true
	l.e = true
	switch s {
	case "debug":
		l.t = false
	case "info":
		l.t = false
		l.d = false
	case "warn":
		l.t = false
		l.d = false
		l.i = false
	case "error":
		l.t = false
		l.d = false
		l.i = false
		l.w = false
	}
}

func Tracef(format string, v ...interface{}) {
	l.Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	l.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	l.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	l.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}

func Trace(v ...interface{}) {
	l.Tracef(fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	l.Debugf(fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	l.Infof(fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	l.Warnf(fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	l.Errorf(fmt.Sprint(v...))
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.t {
		l.T.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.d {
		l.D.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.i {
		l.I.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.w {
		l.W.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.e {
		l.E.Output(3, fmt.Sprintf(format, v...))
	}
}

// func printPid() string {
// 	return fmt.Sprintf("[PID-%v] ",os.Getpid())
// }
