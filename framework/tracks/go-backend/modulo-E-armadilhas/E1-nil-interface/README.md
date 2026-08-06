# E1 — Nil interface

> Debug | 45 min | Go stdlib

## Objetivo

Explique e corrija o caso em que uma interface contem um ponteiro nil.

## Faca

1. Demonstre a diferenca entre interface nil e ponteiro nil.
2. Reproduza o panic ao chamar o metodo.
3. Implemente `safeWrite(io.Writer, []byte)`.
4. Verifique o `Kind` antes de usar reflexao nilavel.

## Pronto quando

- O programa mostra os dois estados de nil.
- `safeWrite` nao causa panic para interface nil ou ponteiro nil.
- Tipos nao nilaveis nao quebram a verificacao.

## Responda

- Por que `w != nil` mesmo com o ponteiro interno nil?
- Quando seria melhor corrigir o contrato em vez de usar reflexao?

> Confianca: [1-5]
