# Capstone 4 — API de Blog com Cache

> Template: Capstone | Síntese do Módulo D | Sem scaffolding

## Contexto

Você vai construir uma API REST completa para um blog, com cache de posts populares
e uma chain de middlewares production-ready — tudo com stdlib.

## Requisitos

### Funcionalidades obrigatórias

1. **CRUD de posts** (D2): `POST /posts`, `GET /posts/{id}`, `PUT /posts/{id}`,
   `DELETE /posts/{id}`, `GET /posts` (listagem com paginação `?page=1&size=10`).
   Store thread-safe com `sync.RWMutex`.
2. **Cache TTL para leitura** (D1): posts acessados frequentemente são cacheados
   por N segundos. `GET /posts/{id}` verifica o cache antes de ir ao store.
   Cache tem lazy eviction e cleanup goroutine.
3. **Middleware chain** (D3): na ordem correta —
   `CORS → Auth (bearer token) → Recoverer → Logger → mux`.
   - Logger: loga método, path, status, duração
   - Recoverer: captura panics, retorna 500
   - Auth: header `Authorization: Bearer <token>` — token fixo para simplificar
   - CORS: `Access-Control-Allow-*` headers, responde a OPTIONS preflight
4. **Graceful shutdown**: ao receber SIGINT, fecha o servidor HTTP e drena
   requisições em andamento (timeout de 5s).

### Requisitos não-funcionais

- Apenas `net/http` e stdlib
- Sem frameworks de router
- JSON como formato de dados
- IDs auto-incrementais

## Tarefas

### Fase 1: Design

Desenhe:
- Estrutura de um Post (ID, Title, Content, CreatedAt, UpdatedAt)
- Store interface (`Create`, `Get`, `Update`, `Delete`, `List`)
- Cache interface (`Get`, `Set`, `Delete`, `Clear`)
- Fluxo de uma requisição GET: middleware chain → cache lookup → store fallback
- Fluxo de uma requisição POST/PUT: middleware chain → store → cache invalidation

### Fase 2: Implementação

### Fase 3: Testes

- CRUD: criar, ler, atualizar, deletar, listar com paginação
- Cache: GET retorna cached após primeiro acesso; PUT invalida cache
- Concorrência: múltiplas requisições simultâneas sem race condition
- Middleware: OPTIONS retorna 200 (CORS antes de Auth)

### Fase 4: Retrospectiva

## Conceitos envolvidos

- HTTP CRUD, ServeMux, RWMutex — D2 (products-api)
- Cache TTL, lazy eviction, cleanup goroutine — D1 (cache-ttl)
- Middleware chain, Logger, Recoverer, Auth, CORS — D3 (middleware-chain)
