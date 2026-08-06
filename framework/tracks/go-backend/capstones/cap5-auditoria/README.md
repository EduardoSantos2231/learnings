# C5 — Auditoria de codigo Go

> Capstone | 60 min | Revisao geral

## Objetivo

Audite um servico que passa nos testes, mas possui falhas de producao.

## Obrigatorio

1. Encontre nil interface, erro comparado com `==` e slice que retém memoria.
2. Encontre uma goroutine que ignora cancelamento.
3. Localize um acesso concorrente sem protecao.
4. Corrija cada causa sem reescrever o servico.

## Pronto quando

- Cada bug tem evidencia, severidade e correcao registrada.
- Os testes antigos continuam passando.
- `go test -race ./...` passa.

## Responda

- Qual bug teria maior impacto em producao?
- Qual teste teria detectado cada classe de falha?

> Confianca: [1-5]
