package receiver

import (
	"fmt"
	"networking/config"
	"networking/packet"
	"networking/sender"
)

//get, validate, ack packets
//detect duped or out of order packets
//write to file once complete
//send final checksum to sender once complete
//func to send checksum back to sender
//list or set to maintain recordkeeping of already-received chunks
//func to send ACKs back to sender

//func to store chunk as they come in and map sequence number to chunk, for placement.

func GetData() {
	var receivedChunks = make(map[uint32][]byte)

	//wait for packet to come in
	buf := make([]byte, sender.MaxPacketSize)
	n, addr, err := conn.ReadFromUDP(buf)

	//in this line, "packet." refers to the file that function lives in.
	received := packet.DecodePacket(buf[:n])

	if received.Checksum != config.MakeChecksum(received.Payload) {
		fmt.Println("Checksum mismatch — discarding packet")
		return // skip storing or ACKing
	}
	if _, exists := receivedChunks[received.SeqNum]; exists {
		fmt.Println("Duplicate packet — re-ACKing")
		// Still send an ACK for reliability
		config.sendAck(received.SeqNum, conn, addr)
		return
	}

	receivedChunks[received.SeqNum] = received.Payload

	sendAck(received.SeqNum, conn, addr)

	if received.EOF {
		reassembleFile(receivedChunks)
	}
}
