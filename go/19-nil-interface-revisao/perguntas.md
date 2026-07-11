## Perguntas

1. O que exatamente significa "interface value = (type, value) pair"? Dê o par nos 3 momentos do gotcha: `var w io.Writer`, `w = buf`, `w.Write()`.

2. `safeWrite` com `reflect.ValueOf(w).IsNil()` pode panicar — em que caso? Como evitar?

3. Cite 2 situações reais em que esse gotcha aparece (dica: uma é com `error`, outra com `io.Writer`/`io.Reader`).
