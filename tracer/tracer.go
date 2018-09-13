package tracer

import (
	"fmt"
	"io"
)

// Tracer is the interface the defines contract for Tracer service
type Tracer interface {
	Trace(...interface{})
}

// New is the factory to create Tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

// Trace is the implementation of Tracer::Trace interface
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {

}

func Off() Tracer {
	return &nilTracer{}
}
