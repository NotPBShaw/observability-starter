package telemetry

type Span struct {
	Name string
}

func StartSpan(name string) Span {
	return Span{Name: name}
}
