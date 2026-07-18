# R11 — dns-cache

Template: **Design**

Add caching to your DNS resolver from R10.

## Tasks

1. **Design the cache**: what's the key? (domain + record type), what's the value? (records + TTL)
2. Implement cache with TTL expiration (lazy eviction on access)
3. Before resolving, check cache; if hit, return cached; if miss or expired, resolve and cache
4. Handle TTL correctly: use the minimum TTL from the response
5. Test: resolve same domain twice — second call should be instant

## Validation

```bash
go run main.go example.com
# First call: iterative resolution with timing
# Resolved example.com in 850ms

go run main.go example.com
# Second call: cache hit
# Resolved example.com (cached) in 0ms
```
