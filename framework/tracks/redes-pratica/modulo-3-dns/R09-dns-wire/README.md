# R09 — dns-wire

Template: **Implementação**

Build a DNS query from scratch in wire format (binary).

## Tasks

1. Construct DNS header: ID (random 16 bits), flags (standard query, recursion desired), QDCOUNT=1
2. Encode domain name in DNS label format (`www.example.com` → `3www7example3com0`)
3. Encode question: QTYPE=A (1), QCLASS=IN (1)
4. Send query via UDP to `8.8.8.8:53` and receive response
5. Parse and print the response (at minimum: the IP address in the answer)

Use `encoding/binary BigEndian`.

## Restriction

Do **NOT** use `net.LookupHost` or any DNS library. Build the wire format yourself.

## Validation

```bash
go run main.go www.example.com
# Expected output: parsed IP address(es) for www.example.com
```
