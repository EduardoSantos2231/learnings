# R06 — http-server

Template: **Implementação**

Build on R5 parser to create a minimal HTTP server that responds to GET requests.

## Tasks

1. Accept TCP connections, parse request with your R5 parser
2. Route `GET /` to return 200 with `Hello World`
3. Route `GET /file/{name}` to serve files from `./www/` directory
4. Return proper HTTP responses with status line and headers
5. Return 404 for unknown paths, 405 for non-GET methods

## Restriction

Do **NOT** use `net/http`. Use your R5 parser with raw TCP.

## Validation

```bash
# Terminal 1: Start server
go run server.go
# Creates ./www/ directory if missing, add some files

# Terminal 2: Test with curl
curl -v http://localhost:9000/
# < HTTP/1.1 200 OK
# Hello World

curl -v http://localhost:9000/file/test.txt
# < HTTP/1.1 200 OK
# (contents of ./www/test.txt)

curl -v http://localhost:9000/nonexistent
# < HTTP/1.1 404 Not Found

curl -v -X POST http://localhost:9000/
# < HTTP/1.1 405 Method Not Allowed
```
