package pools

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
)

type Tracker struct {
	file string
	line int
	disposed bool
	typeName string
}

func NewTracker(value interface{}) *Tracker {
	_, file, line, _ := runtime.Caller(2)
	tracker := &Tracker{
		file:     file,
		line:     line,
		disposed: false,
		typeName: reflect.TypeOf(value).String(),
	}
	runtime.SetFinalizer(tracker, func(f *Tracker) { f.Finalize() })
	return tracker
}

func (t *Tracker) Dispose() {
	t.disposed = true
	runtime.SetFinalizer(t, nil)
}

func (t *Tracker) Finalize() {
	if !t.disposed {
		_, _ = fmt.Fprintf(os.Stderr, "LEAK: %s was leaked. Created at: %s (%d)\n", t.typeName, t.file, t.line)
	}
}
