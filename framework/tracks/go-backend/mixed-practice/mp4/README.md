# Mixed Practice 4 — Debugging de APIs

> Interleaving: encontre e corrija problemas em APIs HTTP reais.

## Cenário 1: Middleware Fora de Ordem

Uma API tem os middlewares: `Auth → Logger → Recoverer → CORS → mux`.
Requisições OPTIONS (preflight) estão retornando 401 antes de chegar ao CORS.

**Pergunta:** Qual é o bug? Como corrigir? Qual a ordem correta e por quê?
Corrija o código.

## Cenário 2: Race Condition no Cache

Um cache TTL é usado por uma API. Sob carga (100 req/s), clientes diferentes
estão recebendo valores inconsistentes para a mesma chave.

**Pergunta:** Onde está a race condition? RWMutex, cleanup goroutine, ou lazy eviction?
Corrija.

## Cenário 3: Vazamento de Goroutine

Após 1 hora rodando, a API tem 5000 goroutines vivas. O código cria uma
goroutine por requisição para logging assíncrono.

**Pergunta:** Qual o problema e como corrigir? Channel bufferizado, worker pool,
ou log síncrono? Implemente a correção.
