package receiver

import ("networking/config"
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
var receivedChunks = make(map[unit32][]byte) 

//create checksum to send back to sender
const PrevCheck = MakeChecksum(data)


const chunk = ReadChunk()


//wait for packet to come in
buf := make([]byte, MaxPacketSize)
n, addr, err := conn.ReadFromUDP(buf)

//in this line, "packet." refers to the file that function lives in. 
received := packet.DecodePacket(buf[:n])

if received.Checksum != checksum.MakeChecksum(received.Payload) {
	fmt.Println("Checksum mismatch — discarding packet")
	return // skip storing or ACKing
}
if _, exists := receivedChunks[received.SeqNum]; exists {
	fmt.Println("Duplicate packet — re-ACKing")
	// Still send an ACK for reliability
	sendAck(received.SeqNum, conn, addr)
	return
}

receivedChunks[received.SeqNum] = received.Payload

sendAck(received.SeqNum, conn, addr)

if received.EOF {
	reassembleFile(receivedChunks)
}