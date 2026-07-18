# R10 — dns-resolver

Template: **Implementação**

Implement iterative DNS resolution starting from root servers.

## Tasks

1. Hardcode root server IPs (`a.root-servers.net` = 198.41.0.4, etc.)
2. Query root for TLD nameserver (e.g., ask root for `.com` NS)
3. Query TLD nameserver for authoritative nameserver
4. Query authoritative nameserver for the A record
5. Handle CNAME chains (follow until you get an A record)
6. Handle NXDOMAIN (domain doesn't exist)

## Restriction

Do **NOT** use `net.LookupHost` or any DNS library. Use your R9 wire-format builder.

## Validation

```bash
go run main.go example.com
# Expected: iteratively resolve from root, print each step and final IP

go run main.go nonexistent-domain-xyz123.com
# Expected: NXDOMAIN
```
