# B4 — Rate Limiter

> Implementacao | 60 min | Go stdlib

## Objetivo

Limite chamadas a uma taxa configuravel sem corrida de dados.

## Faca

1. Implemente um token bucket simples.
2. Permita esperar ou recusar quando nao houver token.
3. Proteja o estado compartilhado.
4. Permita encerrar a reposicao de tokens.

## Restricoes

- Nao use biblioteca externa.
- Nao crie uma goroutine por chamada.

## Pronto quando

- A taxa observada respeita o limite.
- `go test -race ./...` passa.
- O encerramento libera todos os recursos.

## Responda

- Qual estado precisa do mutex?
- Qual trade-off existe entre esperar e recusar?

> Confianca: [1-5]
