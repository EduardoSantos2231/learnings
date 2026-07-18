# Módulo D — HTTP & APIs

> Scaffolding: médio

**Objetivo:** Construir APIs HTTP completas com cache, middleware e concorrência segura.

## Desafios

| # | Desafio | O que você construiu |
|---|---------|---------------------|
| D1 | cache-ttl | Cache com TTL, RWMutex, lazy eviction e cleanup goroutine |
| D2 | products-api | CRUD HTTP com ServeMux, RWMutex no store e middleware |
| D3 | middleware-chain | Chain de middlewares: Logger, Recoverer, Auth, CORS |

## Conceitos dominados ao final do módulo

- `sync.RWMutex` — múltiplos leitores, um escritor
- Lazy eviction — expirar entradas apenas no acesso (não em background)
- Cleanup goroutine — varrer entradas expiradas periodicamente
- `chan struct{}` — sinalização sem dados (tamanho zero)
- `net/http.ServeMux` — roteamento de handlers
- HTTP handlers — `func(w http.ResponseWriter, r *http.Request)`
- CRUD completo — POST, GET, PUT, DELETE
- Middleware — função que recebe e retorna http.Handler
- Chain — composição de múltiplos middlewares em ordem
- Logger, Recoverer, Auth (bearer token), CORS

## Para revisitar

Consulte o `spaced-repetition/schedule.md` para as datas de revisão agendadas.
