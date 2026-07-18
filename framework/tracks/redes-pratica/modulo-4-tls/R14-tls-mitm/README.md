# R14 — tls-mitm

Template: **Debug**

Build an educational MITM proxy that intercepts TLS connections.

## Tasks

1. Proxy listens on `localhost:8080`, forwards to real server
2. Generate certificate for the target hostname on-the-fly, signed by your CA (R12)
3. Client connects to proxy → proxy connects to real server → proxy re-encrypts with fake cert
4. Print decrypted traffic to stdout (**educational purpose ONLY**)
5. **Explain**: why does this work? Why does it fail if the client doesn't trust your CA?

## Validation

```bash
# Terminal 1: Start MITM proxy
go run proxy.go

# Terminal 2: Test with curl (trusting your CA)
curl --cacert ca.pem https://localhost:8080/
# Proxy intercepts, prints decrypted traffic, forwards to real server
```
