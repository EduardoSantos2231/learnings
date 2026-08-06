# C4 — API de posts com cache

> Capstone | 90 min | Modulo D

## Objetivo

Evolua a Products API para uma API de posts com cache de leitura.

## Obrigatorio

1. Implemente `POST`, `GET`, `PUT` e `DELETE /posts/{id}`.
2. Proteja o store com `sync.RWMutex`.
3. Use cache TTL em `GET` e invalide-o em `PUT` e `DELETE`.
4. Adicione logger e recoverer como middlewares.

## Fora de escopo

- Autenticacao, paginacao, banco externo e frontend.

## Pronto quando

- CRUD, cache e invalidacao possuem testes.
- Panic do handler retorna `500` sem derrubar o servidor.
- `go test -race ./...` passa.

## Responda

- Onde o cache pode devolver dado antigo?
- Qual ordem os middlewares precisam seguir?

> Confianca: [1-5]
