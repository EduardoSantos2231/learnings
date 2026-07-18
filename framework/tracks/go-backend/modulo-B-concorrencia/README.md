# Módulo B — Concorrência

> Scaffolding: alto

**Objetivo:** Dominar goroutines, channels, sincronização e padrões de concorrência.

## Desafios

| # | Desafio | O que você construiu |
|---|---------|---------------------|
| B1 | worker-pool | Pool de N workers processando jobs de um channel com WaitGroup e context |
| B2 | parallel-query | Consulta paralela a múltiplas fontes com fan-in simples e timeout |
| B3 | fan-in | Merge de múltiplos channels com select, graceful shutdown via context |
| B4 | rate-limiter | Token bucket com sync.Mutex, reabastecimento em goroutine background |
| B5 | select-sem-default | Uso correto de select sem default, send como case, cancelamento via ctx.Done() |

## Conceitos dominados ao final do módulo

- Goroutines — `go func()` e ciclo de vida
- Channels — buffered vs unbuffered, send/receive, close
- `sync.WaitGroup` — Add/Done/Wait para sincronização
- `context.Context` — cancelamento e timeout propagados
- Select — multiplexação de channels
- Fan-in / Fan-out — distribuir e coletar trabalho
- `sync.Mutex` — proteção de estado compartilhado
- Token bucket — algoritmo de rate limiting
- Armadilha do select sem default — bloqueio vs não-bloqueio

## Para revisitar

Consulte o `spaced-repetition/schedule.md` para as datas de revisão agendadas.
