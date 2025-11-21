package telemetry

import "testing"

func TestCounterIncrements(t *testing.T) {
	c := Counter{Name: "x"}
	c.Inc()
	if c.Value != 1 {
		t.Fatalf("expected 1 got %d", c.Value)
	}
}
