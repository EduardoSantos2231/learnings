# R01 — echo-server

Template: **Implementação**

TCP echo server. Read from connection, write back the same bytes. Then write a client that sends a message and reads the response. Use `net.Listen` and `net.Dial`.

## Tasks

1. Implement server that listens on `:9000` and echoes back whatever it receives
2. Implement client that connects and sends a message, then prints the response
3. Test with multiple simultaneous clients using goroutines

## Restriction

Only `net` stdlib. No third-party packages.

## Validation

```bash
# Terminal 1: Start the server
go run server.go

# Terminal 2: Test with netcat
nc localhost 9000
> hello
hello
> world
world
```
