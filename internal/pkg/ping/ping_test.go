package ping

import (
	"testing"
	"time"
)

const layout = "15:04:00"

func TestPing(t *testing.T) {
	currentTime := time.Now().UTC()
	t.Log("PING")
	t.Logf("PING current time UTC: %v\n", currentTime.Format(layout))
}

func TestPong(t *testing.T) {
	currentTime := time.Now()
	t.Log("PONG")
	t.Logf("PONG current time: %v\n", currentTime.Format(layout))
}
