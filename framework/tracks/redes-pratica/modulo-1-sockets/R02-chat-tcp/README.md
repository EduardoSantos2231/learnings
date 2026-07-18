# R02 — chat-tcp

Template: **Implementação**

Multi-client TCP chat. Server accepts connections, each client can send messages that get broadcasted to all other clients.

## Tasks

1. Server with one goroutine per connection
2. Broadcast: when a client sends a message, all other clients receive it prefixed with a sender ID
3. Handle disconnection: remove client from broadcast list on disconnect

Use a mutex for the clients map.

## Restriction

Only `net` stdlib. No third-party packages.

## Validation

```bash
# Terminal 1: Start server
go run server.go

# Terminal 2: Connect client A
nc localhost 9000
> Connected as client-1
> hello everyone          # broadcasted to all others

# Terminal 3: Connect client B
nc localhost 9000
> Connected as client-2
> [client-1] hello everyone   # received broadcast
```
