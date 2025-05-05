//define chunk sizes, timeouts, window sizes, retries, ports

// func to ensure chunk is always smaller than buffer, make chunck smaller if timeout on
// previous send
package main

func ChunkSize() {
	//init as fixed value, maybe 1024 bytes
	//decrement with failed packets
}

func Timeout() {
	//init as fixed val
	//increments with failed packets
}

func Retry() {
	//if prev packet fails (no ack after 500ms, what else?)
}

func TimeAck() {
	//time delay between packet send and ack response
}

func ReadChunk() {
	//as titled
}

func AcceptableWait() {
	//static set. if response takes this time or less, maintain packet size, speed, etc
}

type AckPacket struct {
	AckedSeqNum uint32
}
