This project concerns implementation of TCP-style semantics and functionality within UDP.

Functionality to inherit:
1.dynamic wait times, packet sizes implied from recipient buffer size, network congestion/stability.
<br>
2. per-packet checksums
3. Initial, final checksum comparisons
4. In-order packet delivery
5. At-recipient packet ordering
6. Incremented packet sequencing
7. Duplicate packet restrictions

Why? This is just homework, man. *For the love of the Go Game*




rough draft pseudocode:

sender opens file in chunks.
<br>
sender checksums entire file.
sender adds a squence number and checksum to each chunk.
sender sends chunk
sender waits for ack from recipient before sending another chunk,
  and takes flow, congestion control, recipient buffer size into consideration. 
recipient sends ack for previous chunk 
iterate until transmission is complete.
WHAT IF A CHUNK FAILS?
  -re-send that chunk 
recipient correctly orders chunks 
recipient generates checksum for entire file
recipient sends checksum to sender
transmission successful if checksum matches. 
