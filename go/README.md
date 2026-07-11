# Estudos Go

Roteiro completo em [roteiro.md](./roteiro.md).

## Exercicios

| # | Projeto | Conceitos |
|---|---------|-----------|
| 01 | currency-conversor | CLI, sentinel errors, errors.Is |
| 02 | bank-account | Pointer receiver, erro tipado, errors.As |
| 03 | worker-pool | Goroutines, channels, WaitGroup, context |
| 04 | paralell-query | Fan-in simple, time.Duration, N goroutines |
| 05 | fan-in | select multi-canal, merger, graceful shutdown |
| 06 | rate-limiter | sync.Mutex, token bucket, goroutine background |
| 07 | linked-list | Ponteiros, nil seguro, Remove/Reverse |
| 08 | shape-interface | Interface implicita, type switch |
| 09 | stack-queue | Thread-safe, pub/priv split, memory leak |
| 10 | io-reader-writer | io.Reader/Writer, delegacao, io.Copy |
| 11 | bst | Arvore binaria, recursao, Delete (3 casos) |
| 12 | cache-ttl | RWMutex, lazy eviction, cleanup goroutine, chan struct{} |
| 13 | products-api | HTTP CRUD, ServeMux, RWMutex, middleware |
| 14 | nil-interface-gotcha | interface (type, value) pair, nil pointer vs nil interface, reflect |
| 15 | middleware-chain | Middleware chain, Logger, Recoverer, Auth, CORS |
| 16 | select-sem-default | select sem default, send como case, cancelamento ctx.Done() |
| 17 | error-is-as | errors.Is vs errors.As, valor sentinela vs tipo, %w wrapping |
| 18 | slice-leak | Slice backing array, memory leak em Pop vs Dequeue |
| 19 | nil-interface-revisao | Interface (type, value) pair, nil pointer vs nil interface |


## Para retomar

Leia o [roteiro.md](./roteiro.md) e me chame: "Professor, estou no Modulo ___. Vamos comecar o exercicio ___."
