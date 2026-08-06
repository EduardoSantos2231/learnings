# C5 — Slice e backing array

> Debug | 45 min | Go stdlib

## Objetivo

Demonstre e corrija referencias mantidas por uma slice depois de remover itens.

## Faca

1. Crie uma slice com capacidade maior que o tamanho.
2. Compare `pop` que nao zera com `popSafe` que zera.
3. Remova o primeiro item e observe o backing array.
4. Teste o comportamento depois de `append`.

## Pronto quando

- O programa mostra a referencia retida e a correcao.
- `popSafe` zera o elemento antes de cortar a slice.
- O teste cobre capacidade extra e append posterior.

## Responda

- Por que o tamanho da slice nao descreve toda a memoria referenciada?
- Qual operacao exige zerar referencias?

> Confianca: [1-5]
