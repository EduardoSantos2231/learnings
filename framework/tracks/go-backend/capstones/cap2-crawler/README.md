# C2 — Crawler concorrente

> Capstone | 90 min | Modulo B

## Objetivo

Consuma uma lista de URLs, processe-as em paralelo e produza um relatorio.

## Obrigatorio

1. Use um worker pool com `N` workers configuravel.
2. Limite chamadas por dominio com um rate limiter.
3. Combine resultados em um canal de saida.
4. Cancele com timeout e `SIGINT`, retornando resultados parciais.

## Fora de escopo

- Parsear HTML e descobrir links novos; a entrada ja e uma lista de URLs.

## Pronto quando

- Cada URL aparece uma vez no relatorio.
- O limite de workers e taxa e testado.
- Cancelamento nao deixa goroutines vivas.
- `go test -race ./...` passa.

## Responda

- Onde o contexto e fechado ou propagado?
- Qual canal possui cada responsabilidade?

> Confianca: [1-5]
