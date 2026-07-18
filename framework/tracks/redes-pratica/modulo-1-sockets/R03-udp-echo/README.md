# R03 — udp-echo

Template: **Implementação**

UDP echo server and client. Unlike TCP, UDP has no connection, no ordering guarantee, and messages can be lost.

## Tasks

1. UDP server with `net.ListenUDP`
2. UDP client that sends 10 messages with sequence numbers
3. Observe: do messages arrive in order? Are any lost?
4. Compare with TCP: what are the practical differences?

## Restriction

Only `net` stdlib. No third-party packages.

## Validation

```bash
# Terminal 1: Start server
go run server.go

# Terminal 2: Run client, observe output
go run client.go
# Expected: see 10 sends, note how many arrive and in what order
```
