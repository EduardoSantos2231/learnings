# D1 — Cache TTL

> Design | 60 min | Go stdlib

## Objetivo

Projete e implemente um cache concorrente que expire entradas por TTL.

## Faca

1. Defina `Get`, `Set` e `Delete`.
2. Escolha expiracao sob demanda ou limpeza periodica.
3. Proteja leituras e escritas concorrentes.
4. Teste expiracao sem `Sleep` longo.

## Pronto quando

- Entrada expirada nunca e retornada.
- `go test -race ./...` passa.
- O design documenta lock, cleanup e chave ausente.

## Responda

- O que acontece se a goroutine de cleanup parar?
- Qual trade-off existe entre lazy eviction e cleanup periodico?

> Confianca: [1-5]
