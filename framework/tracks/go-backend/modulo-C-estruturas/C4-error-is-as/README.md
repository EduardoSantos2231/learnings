# C4 — errors.Is e errors.As

> Debug | 45 min | Go stdlib

## Objetivo

Corrija o handler que confunde erro sentinela, erro tipado e erro wrappeado.

## Faca

1. Reproduza os casos `not found`, validacao e banco.
2. Use `errors.Is` para valor sentinela.
3. Use `errors.As` para extrair o erro tipado.
4. Preserve contexto com `%w`.

## Pronto quando

- Os tres erros sao classificados corretamente.
- Comparar o erro wrappeado com `==` deixa de ser necessario.
- `go test ./...` passa.

## Responda

- O que se perde ao usar `%v` no lugar de `%w`?
- Quando `errors.As` seria a escolha errada?

> Confianca: [1-5]
