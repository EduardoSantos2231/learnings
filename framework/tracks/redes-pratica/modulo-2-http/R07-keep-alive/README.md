# R07 — keep-alive

Template: **Design**

HTTP/1.1 uses persistent connections by default. Your R6 server closes after each response. Design and implement keep-alive support.

## Tasks

1. **Design**: how will you keep the connection open? How will you know when the client is done?
2. **Implement**: after sending response, read next request on same connection
3. Add timeout: close idle connections after 10s
4. Handle `Connection: close` header from client

## Restriction

Do **NOT** use `net/http`. Build on R6.

## Validation

```bash
curl -v --keepalive-time 5 http://localhost:9000/
# Should reuse same connection for multiple requests
```
