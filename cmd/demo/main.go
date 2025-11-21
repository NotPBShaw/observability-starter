package main

import (
	"fmt"

	"observability-starter-kit/internal/telemetry"
)

func main() {
	c := telemetry.Counter{Name: "requests_total"}
	c.Inc()
	span := telemetry.StartSpan("demo-request")
	fmt.Println(telemetry.Info(span.Name), c.Value)
}
