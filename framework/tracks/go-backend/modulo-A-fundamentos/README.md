# Módulo A — Fundamentos da Linguagem

> Scaffolding: alto

**Objetivo:** Dominar erros, ponteiros e interfaces — a base de tudo em Go.

## Desafios

| # | Desafio | O que você construiu |
|---|---------|---------------------|
| A1 | currency-conversor | CLI com taxas de câmbio fixas, tratamento de erros com `errors.Is` e sentinel errors |
| A2 | bank-account | Sistema de conta bancária com pointer receivers, erro tipado e `errors.As` |
| A3 | linked-list | Lista ligada com operações Add, Remove, Reverse — segurança contra nil |
| A4 | shape-interface | Cálculo de área/perímetro com interface implícita e type switch |

## Conceitos dominados ao final do módulo

- Erros: sentinel errors (`var ErrX = errors.New(...)`) vs erro tipado (struct com `Error()`)
- `errors.Is` vs `errors.As` — quando usar cada um
- Pointer receiver (`func (a *Account)`) — quando usar e por quê
- Ponteiros em structs autorreferenciais (`*Node`)
- Nil safety — verificar `node != nil` antes de acessar campos
- Interface implícita — qualquer tipo que implementa os métodos satisfaz a interface
- Type switch — `switch v := x.(type)` para lógica condicional por tipo

## Para revisitar

Consulte o `spaced-repetition/schedule.md` para as datas de revisão agendadas.
