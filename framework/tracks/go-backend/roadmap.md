# Roadmap — Go Backend

> Track de Go do zero a APIs production-ready.
> Siga a ordem. O professor anuncia o próximo automaticamente.

## Posição atual

| Campo | Valor |
|-------|-------|
| Módulo atual | A — Fundamentos |
| Último concluído | — |
| Próximo desafio | A1-currency-conversor |
| Próximo formato | Implementação |

---

## Módulo A — Fundamentos da Linguagem

> Scaffolding: alto (enunciados detalhados, dicas de stdlib)

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| A1 | currency-conversor | Implementação | CLI, sentinel errors, errors.Is | ✅ |
| A2 | bank-account | Implementação | Pointer receiver, erro tipado, errors.As | ✅ |
| A3 | linked-list | Implementação | Ponteiros, nil seguro, Remove/Reverse | ✅ |
| A4 | shape-interface | Implementação | Interface implícita, type switch | ✅ |

---

## Mixed Practice 1

> Interleaving: decidir qual padrão usar sem saber de antemão

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| MP1 | escolha-ferramentas | Mixed Practice | Erros, interfaces, estruturas | ⬜ |

---

## Capstone 1 — Calculadora de Expressões CLI

> Síntese do Módulo A

| # | Desafio | Template | Conceitos envolvidos | Status |
|---|---------|----------|---------------------|--------|
| C1 | calculadora | Capstone | Erros (A1, A2), Interfaces (A4), Estruturas (A3) | ⬜ |

---

## Módulo B — Concorrência

> Scaffolding: alto

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| B1 | worker-pool | Implementação | Goroutines, channels, WaitGroup, context | ✅ |
| B2 | parallel-query | Debug | Fan-in simples, time.Duration | ✅ |
| B3 | fan-in | Implementação | Select multi-canal, merger, graceful shutdown | ✅ |
| B4 | rate-limiter | Implementação | sync.Mutex, token bucket, goroutine background | ✅ |
| B5 | select-sem-default | Otimização | Select sem default, ctx.Done(), send como case | ✅ |

---

## Mixed Practice 2

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| MP2 | padrao-concorrencia | Mixed Practice | Worker pool vs fan-in vs rate limiter | ⬜ |

---

## Capstone 2 — Crawler Web Concorrente

> Síntese do Módulo B

| # | Desafio | Template | Conceitos envolvidos | Status |
|---|---------|----------|---------------------|--------|
| C2 | crawler | Capstone | Worker pool (B1), Fan-in (B2, B3), Rate limiter (B4), Context (B5) | ⬜ |

---

## Módulo C — Estruturas de Dados & I/O

> Scaffolding: médio (sem dicas de implementação)

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| C1 | stack-queue | Explicação | Thread-safe, pub/priv split, memory leak | ✅ |
| C2 | io-reader-writer | Implementação | io.Reader/Writer, delegação, io.Copy | ✅ |
| C3 | bst | Implementação | Árvore binária, recursão, Delete (3 casos) | ✅ |
| C4 | error-is-as | Debug | errors.Is vs errors.As, %w wrapping | ✅ |
| C5 | slice-leak | Debug | Backing array, memory leak em Pop vs Dequeue | ✅ |

---

## Mixed Practice 3

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| MP3 | estrutura-certa | Mixed Practice | Stack, Queue, BST, io.Reader, slices | ⬜ |

---

## Capstone 3 — Cache com Persistência

> Síntese do Módulo C

| # | Desafio | Template | Conceitos envolvidos | Status |
|---|---------|----------|---------------------|--------|
| C3 | cache-persistente | Capstone | Estruturas (C1), I/O (C2), BST (C3), Error wrap (C4), Slices (C5) | ⬜ |

---

## Módulo D — HTTP & APIs

> Scaffolding: médio

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| D1 | cache-ttl | Design | RWMutex, lazy eviction, cleanup goroutine | ✅ |
| D2 | products-api | Implementação | HTTP CRUD, ServeMux, RWMutex, middleware | ✅ |
| D3 | middleware-chain | Implementação | Logger, Recoverer, Auth, CORS, chain() | ✅ |

---

## Mixed Practice 4

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| MP4 | debugging-apis | Mixed Practice | Middleware, concorrência, erros em APIs | ⬜ |

---

## Capstone 4 — API de Blog com Cache

> Síntese do Módulo D

| # | Desafio | Template | Conceitos envolvidos | Status |
|---|---------|----------|---------------------|--------|
| C4 | api-blog | Capstone | CRUD (D2), Cache TTL (D1), Middleware (D3) | ⬜ |

---

## Módulo E — Armadilhas da Linguagem

> Scaffolding: baixo (apenas o problema)

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| E1 | nil-interface | Debug | Interface (type, value) pair, nil pointer vs nil interface | ✅ |
| E2 | nil-interface-revisao | Explicação | Revisão aprofundada do gotcha | ✅ |

---

## Capstone 5 — Auditoria de Código Go

> Síntese do Módulo E + revisão geral

| # | Desafio | Template | Conceitos envolvidos | Status |
|---|---------|----------|---------------------|--------|
| C5 | auditoria | Capstone | Nil interface (E1, E2), Slice leak (C5), Error wrap (C4), Select (B5) | ⬜ |

---

## Legenda

- ✅ Concluído
- ⬜ Pendente
- 🔄 Revisão pendente (spaced repetition)
