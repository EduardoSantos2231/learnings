# R05 — http-parser

Template: **Implementação**

Parse raw HTTP/1.1 requests from bytes. You receive raw bytes from a TCP connection and must extract method, path, headers, and body.

## Tasks

1. Parse request line: `GET /index.html HTTP/1.1`
2. Parse headers until empty line (`\r\n\r\n`)
3. If `Content-Length` header exists, read exactly that many bytes for the body
4. Return a struct with `Method`, `Path`, `Headers map[string]string`, `Body []byte`

Use `bufio.Reader`.

## Restriction

Do **NOT** use `net/http`. Raw TCP + your own parser only.

## Validation

```bash
# Terminal 1: Start your raw TCP server
go run server.go

# Terminal 2: Test with curl
curl -v -X POST -d "hello" http://localhost:9000/
```
