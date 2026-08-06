# A1 — Currency Converter

> Implementacao | 45 min | Go stdlib

## Objetivo

Implemente uma CLI que converta um valor entre moedas conhecidas.

## Faca

1. Leia valor, origem e destino pelos argumentos.
2. Valide valor positivo e moedas suportadas.
3. Separe validacao, conversao e apresentacao.
4. Use erros identificaveis para entradas invalidas.

## Restricoes

- Use apenas a biblioteca padrao.
- Nao use `panic` para entrada invalida.

## Pronto quando

- `go test ./...` passa.
- Entrada invalida retorna erro sem resultado falso.
- O resultado e impresso com duas casas decimais.

## Responda

- Quando `errors.Is` e melhor que `==`?
- Qual regra pertence ao dominio e qual pertence a CLI?

> Confianca: [1-5]
