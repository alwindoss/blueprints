package tracer

import (
	"bytes"
	"testing"
)

func TestTrace(t *testing.T) {
	var buf bytes.Buffer
	trcr := New(&buf)
	if trcr == nil {
		t.Errorf("Return from New should not be nil")
	} else {
		trcr.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write: '%s'.", buf.String())
		}
	}
}

func TestTracerOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
