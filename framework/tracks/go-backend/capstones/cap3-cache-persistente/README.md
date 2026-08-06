# C3 — Cache persistente

> Capstone | 90 min | Modulo C

## Objetivo

Implemente um cache em memoria com TTL e salvamento em arquivo.

## Obrigatorio

1. Implemente `Get`, `Set`, `Delete` e `Clear` com protecao concorrente.
2. Expire entradas no acesso.
3. Salve e carregue os dados usando `io.Reader` e `io.Writer`.
4. Encapsule erros de arquivo com `%w`.

## Fora de escopo

- Cleanup em background e formato binario customizado.

## Pronto quando

- Item expirado nunca e retornado.
- Salvar e carregar preservam itens validos.
- `go test -race ./...` passa.

## Responda

- Qual estrutura armazena o cache e por que?
- Como voce evitaria referencias antigas apos `Delete`?

> Confianca: [1-5]
