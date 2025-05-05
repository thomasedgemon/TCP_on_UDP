//define chunk sizes, timeouts, window sizes, retries, ports

// func to ensure chunk is always smaller than buffer, make chunck smaller if timeout on
// previous send
package main

const SenderChunk = RecipientBuffer //unless last transmission failed, then decrement.

//func to determine timeout. static unless last chunk failed
