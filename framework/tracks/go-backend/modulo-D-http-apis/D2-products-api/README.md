# D2 — Products API

> Implementacao | 60 min | Go stdlib

## Objetivo

Implemente uma API HTTP para cadastrar, consultar, atualizar e remover produtos.

## Faca

1. Crie `POST /products` e `GET /products`.
2. Crie `GET`, `PUT` e `DELETE /products/{id}`.
3. Valide JSON e campos obrigatorios.
4. Proteja o armazenamento contra acesso concorrente.

## Pronto quando

- Cada rota retorna o status HTTP correto.
- JSON invalido retorna `400`.
- Produto ausente retorna `404`.
- `go test -race ./...` passa.

## Responda

- Onde deve ficar a validacao?
- O que mudaria para persistir os produtos?

> Confianca: [1-5]
