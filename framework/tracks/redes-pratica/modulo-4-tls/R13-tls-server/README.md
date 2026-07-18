# R13 — tls-server

Template: **Implementação**

Build a TLS server and client that verify each other.

## Tasks

1. Create TLS server with `crypto/tls.Listen` using cert from R12
2. Create TLS client with `crypto/tls.Dial` that verifies the server certificate
3. Test with self-signed cert: client should fail (unknown authority)
4. Add CA cert to client's root pool: client should now succeed
5. Explore `tls.Config`: `InsecureSkipVerify`, `MinVersion`, `CurvePreferences`

## Validation

```bash
# Terminal 1: Start TLS server
go run server.go

# Terminal 2: Test with curl
curl -v --cacert ca.pem https://localhost:8443/

# Without CA: should fail
curl -v https://localhost:8443/
# Expected: SSL certificate problem: self-signed certificate
```
