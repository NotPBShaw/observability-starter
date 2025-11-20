package telemetry

type Counter struct {
	Name  string
	Value int
}

func (c *Counter) Inc() {
	c.Value++
}
