//chunk file, assign sequences numbers to chunks
//track acks
//re-transmit if necessary
//get final checksum

package main

var seqNum uint32 = 0

//no way to know recipients buffer size. assume very small, then increment up until
//theres a failure. 

for {
    chunk, err := ReadChunk(file, seqNum)
    if err == io.EOF {
        break
    }

    pkt := Packet{
        SeqNum:      seqNum,
        Ack:         false,
        EOF:         isLastChunk,       
        PayloadSize: uint16(len(chunk)), //payload size: SenderChunk(prev size unless failed)
        Payload:     chunk,
        Checksum:    MakeChecksum(chunk),
    }

    encoded := EncodePacket(pkt)
    conn.WriteToUDP(encoded, addr)

    // Wait for ACK or handle timeout...

    seqNum++ 
}

