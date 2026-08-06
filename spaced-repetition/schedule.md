# Spaced Repetition Schedule

> A agenda ativa e controlada por `schedule.json` e pelo CLI `tracking`.
> Use `tracking session` no inicio de cada sessao.

## Regra

- Uma sessao escolhe um desafio novo ou uma revisao pratica.
- No maximo uma revisao consecutiva quando houver desafio disponivel.
- Blocos avancam progressivamente: `1d -> 7d -> 30d`.
- Falha agenda um reparo pratico em `1d`; nao recria quatro intervalos.
- Revisoes antigas foram arquivadas no rebaseline e nao precisam ser recuperadas.

## Track ativa

| Track | Papel |
|-------|-------|
| go-backend | Principal |
| docker-devops | Complementar |

## Proxima sessao

```bash
./framework/tracking/tracking session
```

Estado atual: diagnostico pratico de Go pendente.

## Blocos de revisao Go

| Bloco | Origem | Cenarios |
|-------|--------|----------|
| RB-go-fundamentos | A1-A4 | undo, erros, plugins |
| RB-go-concurrency | B1-B5 | lote, gateway, agregador |
| RB-go-structures | C1-C5 | historico, logs, ranking |
| RB-go-http | D1-D3 | middleware, cache, goroutines |

Cada bloco so entra na agenda depois que sua Mixed Practice for concluida.

## Docker complementar

Docker aparece como checkpoint operacional sobre os desafios Go:

- build e run no diagnostico;
- porta e healthcheck na API;
- PID 1 e sinais na concorrencia;
- multi-stage e Compose nos capstones.

## Historico

As revisoes antigas permanecem em `schedule.json` com status `archived`.
Elas nao contam como pendencias e nao serao repetidas automaticamente.
