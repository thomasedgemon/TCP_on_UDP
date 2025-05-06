//define chunk sizes, timeouts, window sizes, retries, ports

package config

import (
	"encoding/binary"
	"hash/crc32"
	"io"
	"net"
	"networking/packet"
	"os"
	"sort"
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
	offset := int64(seqNum) * int64(CurrentChunkSize)
	buffer := make([]byte, CurrentChunkSize)

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

func MakeChecksum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func EncodeAck(ack AckPacket) []byte {
	buf := make([]byte, 6)
	binary.BigEndian.PutUint32(buf[0:4], ack.AckedSeqNum)
	binary.BigEndian.PutUint16(buf[4:6], ack.WindowSize)
	return buf
}

func sendAck(seqNum uint32, conn *net.UDPConn, addr *net.UDPAddr) {
	ack := packet.AckPacket{
		AckedSeqNum: seqNum,
		WindowSize:  0,
	}
	encoded := EncodeAck(ack)
	conn.WriteToUDP(encoded, addr)
}

func reassembleFile(chunks map[uint32][]byte) {
	//open file to write to
	outFile, err := os.Create("output.txt")
	if err != nil {
		panic("failed to create output file:" + err.Error())
	}
	defer outFile.Close()
	//make sorted list of sequence numbers
	var seqNums []int
	for seq := range chunks {
		seqNums = append(seqNums, int(seq))
	}
	sort.Ints(seqNums)

	//write the chunks in order to the file
	for _, seq := range seqNums {
		_, err := outFile.Write(chunks[uint32](seq))
		if err != nil {
			panic("failed to write a chunk:" + err.Error())
		}
	}
	println("file reassembled.")
}

func computeFinalChecksum(filename string) uint32 {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	hasher := crc32.NewIEEE()
	_, err = io.Copy(hasher, f)
	if err != nil {
		panic(err)
	}

	return hasher.Sum32()
}

func DecodeAckPacket(seqNum uint32, conn *net.UDPConn, addr *net.UDPAddr) {

	//logic
}

func isLastChunk() {
	//logic
}
