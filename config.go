//define chunk sizes, timeouts, window sizes, retries, ports

package config

import (
	"io"
	"net"
	"os"
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

func ReadChunk(file *os.File, seqNum uint32) ([]byte, error) {
	//offset pairs sequence number with the next file chunk to be read in
	offset := int64(seqNum) * int64(config.CurrentChunkSize)
	buffer := make([]byte, config.CurrentChunkSize)

	_, err := file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return buffer[:n], err
}

func MakeChecksum() {
	//logic
}

func sendAck(seqNum uint32, conn *net.UDPConn, addr *net.UDPAddr) {
	ack := config.AckPacket{
		AckedSeqNum: seqNum,
		WindowSize:  0, // optional for now
	}
	encoded := config.EncodeAck(ack)
	conn.WriteToUDP(encoded, addr)
}
