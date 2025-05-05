<h1>This project concerns implementation of TCP-style semantics and functionality within UDP.</h1>

<p1>Functionality to inherit:</p1>
1.dynamic wait times, packet sizes implied from recipient buffer size, network congestion/stability.
<br>
2. per-packet checksums
<br>
3. Initial, final checksum comparisons
<br>
4. In-order packet delivery
<br>
5. At-recipient packet ordering
<br>
6. Incremented packet sequencing
<br>
7. Duplicate packet restrictions
<br>
<br>
<br>

Why? This is just homework, man. *For the love of the Go Game*




<p2>rough draft pseudocode:</p2>

sender opens file in chunks.
<br>
sender checksums entire file.
<br>
sender adds a squence number and checksum to each chunk.
<br>
sender sends chunk
<br>
sender waits for ack from recipient before sending another chunk,
  and takes flow, congestion control, recipient buffer size into consideration. 
  <br>
recipient sends ack for previous chunk 
<br>
iterate until transmission is complete.
<br>
WHAT IF A CHUNK FAILS?
  -re-send that chunk, adjust wait time, packet size
  <br>
recipient correctly orders chunks 
<br>
recipient generates checksum for entire file
<br>
recipient sends checksum to sender
<br>
transmission successful if checksum matches. 
