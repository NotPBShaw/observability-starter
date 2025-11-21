package telemetry

import "fmt"

func Info(msg string) string {
	return fmt.Sprintf("level=info msg=%q", msg)
}
