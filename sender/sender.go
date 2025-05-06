//chunk file, assign sequences numbers to chunks
//track acks
//re-transmit if necessary
//get final checksum

package sender

import (
	"io"
	"networking/config"
	"networking/packet"
	"time"
)

var seqNum uint32 = 0

const MaxRetries int = 5
const MaxPacketSize int = 1200

//no way to know recipients buffer size. assume x, then decrement if
//theres a failure.

// all executable logic must be inside a function!!
func sendFile() {
	for {
		chunk, err := config.ReadChunk(file, seqNum)
		if err == io.EOF {
			break
		}

		pkt := config.Packet{
			SeqNum:      seqNum,
			Ack:         false,
			EOF:         isLastChunk,
			PayloadSize: uint16(len(chunk)), //payload size: SenderChunk(1200 bytes unless failed)
			Payload:     chunk,
			Checksum:    config.MakeChecksum(chunk),
		}

		encoded := packet.EncodePacket(pkt)
		retries := 0
		for retries < MaxRetries {
			con.WriteToUDP(encoded, addr)
		}
		conn.WriteToUDP(encoded, addr)

		//wait for ack
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		//make buffer to catch what comes back
		buf := make([]byte, MaxPacketSize)
		//fill said buffer with data
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			retries++
			continue
		}
		ack := config.DecodeAckPacket(buf[:n])
		if ack.AckedSeqNum == seqNum {
			break
		}
		retries++
	}
	if retries == maxRetries {
		break
	}
	seqNum++
}
