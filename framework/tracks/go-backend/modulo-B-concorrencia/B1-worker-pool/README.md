# B1 — Worker Pool

> Implementacao | 60 min | Go stdlib

## Objetivo

Processe jobs com exatamente `N` workers e encerre sem deixar goroutines vivas.

## Faca

1. Receba jobs por canal e envie resultados por outro.
2. Inicie `N` workers configuravel.
3. Feche os canais na ordem correta.
4. Cancele o processamento com `context.Context`.

## Pronto quando

- Todo job concluido aparece uma vez.
- Cancelamento encerra workers e consumidores.
- `go test -race ./...` passa.

## Responda

- Quem fecha cada canal?
- O que acontece se o consumidor parar de ler?

> Confianca: [1-5]
