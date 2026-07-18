# Capstone 1 — Proxy HTTP Reverso

> Síntese dos Módulos 1 e 2 | Sem scaffolding

## Contexto

Construa um proxy HTTP reverso: ele recebe requisições HTTP e as encaminha
para um servidor backend. Útil para load balancing, caching, SSL termination.

## Requisitos

1. Proxy escuta em :8080, encaminha para backend em :9000
2. Suporta múltiplos backends com round-robin
3. Healthcheck: se um backend falhar 3x, remove da rotação até voltar
4. Keep-alive: reusa conexões com backends
5. Timeout: se backend não responder em 5s, retorna 504

## Conceitos envolvidos

- TCP connect, read, write — R1 (echo-server)
- Múltiplas conexões concorrentes — R2 (chat-tcp)
- Timeout e retry — R4 (timeout-retry)
- HTTP parse e response — R5 (http-parser), R6 (http-server)
- Keep-alive — R7
