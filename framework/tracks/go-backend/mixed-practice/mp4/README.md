# MP4 — APIs em contexto

> Revisao pratica | Faca um cenario por sessao.

Escolha o cenario indicado pela CLI. Reproduza o problema antes de corrigir.

## Cenarios

1. **Middleware:** corrija preflight CORS bloqueado por autenticacao.
2. **Cache:** corrija a race condition entre leitura, expiracao e cleanup.
3. **Logging:** corrija o vazamento causado por uma goroutine por requisicao.

## Pronto quando

- O bug e reproduzido por um teste.
- A correcao passa sem alterar o contrato HTTP.
- `go test -race ./...` passa quando aplicavel.

## Responda

- Qual era a causa raiz?
- Como impediria a regressao?

> Confianca: [1-5]
