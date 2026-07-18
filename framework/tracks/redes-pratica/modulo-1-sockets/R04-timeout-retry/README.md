# R04 — timeout-retry

Template: **Otimização**

Working TCP client but no timeout handling — it hangs forever if server is down.

## Tasks

1. Identify the problem: current code has no deadline
2. Add `SetDeadline` with configurable timeout
3. Implement retry with exponential backoff (1s, 2s, 4s, max 3 retries)
4. Test different buffer sizes (64B, 1KB, 64KB) — measure throughput difference

Prove with benchmarks.

## Restriction

Only `net` stdlib and `time` stdlib.

## Validation

```bash
# Test: server not running — client should timeout and retry, not hang forever
go run client.go
# Expected: "connection timed out, retrying in 1s..."
#          "connection timed out, retrying in 2s..."
#          "connection timed out, retrying in 4s..."
#          "max retries exceeded"

# Benchmark: compare 64B vs 64KB buffer throughput
go test -bench=. -benchtime=5s
```
