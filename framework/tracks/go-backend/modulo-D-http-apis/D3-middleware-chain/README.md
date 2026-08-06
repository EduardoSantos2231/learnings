# D3 — Middleware chain

> Implementacao | 60 min | Go stdlib

## Objetivo

Adicione logger, recoverer, autenticacao e CORS a uma API HTTP existente.

## Faca

1. Registre metodo, path e duracao.
2. Converta panic do handler em `500`.
3. Exija `Bearer` para rotas protegidas.
4. Responda `OPTIONS` antes da autenticacao.
5. Componha a ordem `CORS -> Auth -> Recoverer -> Logger -> mux`.

## Pronto quando

- GET publico continua acessivel.
- Rota protegida sem token retorna `401`.
- Preflight retorna `200` com headers CORS.
- Panic nao derruba o servidor.

## Responda

- Por que CORS precisa executar antes de Auth?
- Qual middleware deve ser mais externo para capturar panic?

> Confianca: [1-5]
