//define chunk sizes, timeouts, window sizes, retries, ports

package config

import (
	"time"
)

const (
	DefaultChunkSize = 1024
	MaxPacketSize    = 1200
	InitialTimeout   = 500 * time.Millisecond
	MaxRetries       = 5
	AckWaitThreshold = 1000 * time.Millisecond
	DefaultPort      = 9000
)

// Tunable parameters (you can mutate these during runtime)
var (
	CurrentChunkSize = DefaultChunkSize
	CurrentTimeout   = InitialTimeout
)

// Reduce chunk size if previous packet failed
func AdjustChunkSizeOnFailure() {
	if CurrentChunkSize > 256 {
		CurrentChunkSize -= 128
	}
}

// Increase timeout if packet is slow to ACK
func AdjustTimeoutOnFailure() {
	CurrentTimeout += 200 * time.Millisecond
}

// Reset parameters after successful ACK
func ResetParameters() {
	CurrentChunkSize = DefaultChunkSize
	CurrentTimeout = InitialTimeout
}

// Use in sender to decide if timeout exceeded
func AcceptableWait(rtt time.Duration) bool {
	return rtt <= AckWaitThreshold
}

// Define ACK packet structure (can be reused across sender/receiver)
type AckPacket struct {
	AckedSeqNum uint32
	WindowSize  uint16 // optional for flow control
}
