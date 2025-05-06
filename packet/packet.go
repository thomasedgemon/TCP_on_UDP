//define packet structure, payloads, flags

package packet

//defines the structure of a data packet, not to be confused with an ack packet
import (
	"encoding/binary"
)

// Define ACK packet structure (can be reused across sender/receiver)
type AckPacket struct {
	AckedSeqNum uint32
	WindowSize  uint16 // optional for flow control
}

func EncodePacket(pkt Packet) []byte {
	header := make([]byte, 11)
	//"network byte order"
	binary.BigEndian.PutUint32(header[0:4], pkt.SeqNum)
	//define headers based on the type of packet being sent
	if pkt.EOF {
		header[8] = 1
	} else {
		header[4] = 0
	}

	binary.BigEndian.PutUint16(header[6:8], pkt.PayloadSize)
	binary.BigEndian.PutUint32(header[8:12], pkt.Checksum)

	return append(header, pkt.Payload...)
}

func DecodePacket(data []byte) Packet {
	//reference Packet struct for packet construction
	pkt := Packet{}
	//apply values to the fields in the packet struct
	pkt.SeqNum = binary.BigEndian.Uint32(data[0:4])
	pkt.EOF = data[4] == 1
	pkt.PayloadSize = binary.BigEndian.Uint16(data[5:7])
	pkt.Checksum = binary.BigEndian.Uint32(data[7:11])

	// Defensive check in case payload is too short
	if len(data) >= 11 {
		pkt.Payload = make([]byte, pkt.PayloadSize)
		copy(pkt.Payload, data[11:])
	}

	return pkt
}
