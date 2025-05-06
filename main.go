package main

import (
	"fmt"
	"networking/config"
	"networking/packet"
	"os"
)

func main() {
	file, err := os.Open("testfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seqNum := uint32(0)
	chunk, err := config.ReadChunk(file, seqNum)
	if err != nil {
		panic(err)
	}

	isLastChunk := len(chunk) < config.CurrentChunkSize

	pkt := packet.Packet{
		SeqNum:      seqNum,
		EOF:         isLastChunk,
		PayloadSize: uint16(len(chunk)),
		Payload:     chunk,
		Checksum:    checksum.MakeChecksum(chunk),
	}

	encoded := packet.EncodePacket(pkt)

	fmt.Printf("Encoded packet: %x\n", encoded) // Just for testing
}

// parse args, init config, call sender, receiver
//handle high level debugging, logging, and exit conditions
