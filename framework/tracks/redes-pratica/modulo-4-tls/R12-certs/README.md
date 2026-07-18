# R12 — certs

Template: **Explicação**

Generate certificates and explain the chain of trust.

## Tasks

1. Generate a self-signed CA root certificate with `crypto/x509`
2. Generate a server certificate signed by your CA
3. **Explain in plain language** (max 300 words) how certificate validation works
4. Write a validator that checks: certificate signed by CA?, not expired?, hostname matches SAN/CN?
5. Demonstrate: invalid cert (wrong hostname) → validation fails

## Validation

```bash
go run main.go
# Expected: generates ca.pem, ca-key.pem, server.pem, server-key.pem
# Validator: server.pem passes validation for "localhost"
# Validator: server.pem fails validation for "wronghost"
```
