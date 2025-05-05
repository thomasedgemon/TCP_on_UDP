//get, validate, ack packets
//detect duped or out of order packets
//write to file once complete
//send final checksum to sender once complete

package main
//list or set to maintain recordkeeping of already-received chunks

//func to send checksum back to sender

//func to send ACKs back to sender


//func to store chunk as they come in and map sequence number to chunk, for placement.
map[unit32][]byte 

//create checksum to send back to sender
const PrevCheck = MakeChecksum(data)


//wait for packet to come in
buf := make([]byte, MaxPacketSize)
n, addr, err := conn.ReadFromUDP(buf)


