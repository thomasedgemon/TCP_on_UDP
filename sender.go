//chunk file, assign sequences numbers to chunks
//track acks
//re-transmit if necessary
//get final checksum

package main


import (
    "io"
    "net"
    "time"
    "fmt"
)


var seqNum uint32 = 0
const MaxRetries int = 5
const MaxPacketSize int = 1200

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
        PayloadSize: uint16(len(chunk)), //payload size: SenderChunk(1200 bytes unless failed)
        Payload:     chunk,
        Checksum:    MakeChecksum(chunk),
    }

    encoded := EncodePacket(pkt)
	retries := 0
	for retries < MaxRetries {
		con..WriteToUDP(encoded, addr)
	}
    conn.WriteToUDP(encoded, addr)

	//wait for ack 
    conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		//make buffer to catch what comes back
        buf := make([]byte, MaxPacketSize)
		//fill said buffer with data
        n, _, err := conn.ReadFromUDP(buf)
        if err != nil {
            //fmt.Println("Timeout or read error; retrying...")
            retries++
            continue
        }
        ack := DecodeAckPacket(buf[:n])
        if ack.AckedSeqNum == seqNum {
            //fmt.Printf("ACK received for chunk %d\n", seqNum)
            break
        }
        //fmt.Printf("Received unexpected ACK %d (expected %d)\n", ack.AckedSeqNum, seqNum)
        retries++
    }
    if retries == maxRetries {
        //fmt.Printf("Failed to send chunk %d after %d retries\n", seqNum, maxRetries)
        break
    }

    seqNum++
}

