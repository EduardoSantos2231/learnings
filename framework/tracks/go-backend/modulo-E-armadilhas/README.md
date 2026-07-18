# Módulo E — Armadilhas da Linguagem

> Scaffolding: baixo

**Objetivo:** Internalizar os gotchas clássicos de Go que aparecem em entrevistas e code reviews.

## Desafios

| # | Desafio | O que você fez |
|---|---------|---------------|
| E1 | nil-interface | Debug do gotcha: `var w io.Writer; var buf *bytes.Buffer; w = buf; w != nil` |
| E2 | nil-interface-revisao | Revisão aprofundada: interface (type, value) pair, reflect, prevenção |

## Conceitos dominados ao final do módulo

- Interface em Go é um par (type, value) — ambos nil para interface ser nil
- Ponteiro nil atribuído a interface → type não-nil → interface não-nil
- `reflect.ValueOf(w).IsNil()` — só funciona se w.Kind() é nilable (ptr, chan, func, interface, map, slice)
- Prevenção: retornar interface, não ponteiro concreto
- Prevenção: verificar nil antes de atribuir a interface

## Para revisitar

Consulte o `spaced-repetition/schedule.md` para as datas de revisão agendadas.
