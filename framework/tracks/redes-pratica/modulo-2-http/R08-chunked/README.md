# R08 — chunked

Template: **Implementação**

Implement `Transfer-Encoding: chunked` for responses where `Content-Length` is unknown.

## Tasks

1. Implement chunked encoding: send size in hex, then data, then `\r\n`
2. Stream a response that generates data over time (one chunk per second for 5 seconds)
3. Client reads chunked response and reconstructs the full body

## Restriction

Do **NOT** use `net/http`. Build on R6 or R7 server.

## Validation

```bash
curl --raw http://localhost:9000/stream
# Should show raw chunks:
# 5\r\n
# hello\r\n
# 5\r\n
# world\r\n
# 0\r\n
# \r\n
```
