# MP2 — Concorrencia em contexto

> Revisao pratica | Faca um cenario por sessao.

Escolha o cenario indicado pela CLI. Implemente a protecao necessaria e prove o encerramento.

## Cenarios

1. **Lote:** processe 10.000 itens com no maximo oito workers.
2. **Gateway:** proteja um servico limitado a 50 chamadas por segundo.
3. **Agregador:** consulte tres fontes e encerre em no maximo 1,5 segundo.

## Pronto quando

- O limite de concorrencia ou taxa e testado.
- Cancelamento encerra todas as goroutines.
- `go test -race ./...` passa.

## Responda

- Por que escolheu worker pool, rate limiter ou fan-in?
- O que acontece quando uma fonte demora demais?

> Confianca: [1-5]
