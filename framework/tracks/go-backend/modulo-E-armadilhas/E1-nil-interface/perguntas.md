## Perguntas

Responda em `respostas.md`.

### 1. Representação interna

Por que `w != nil` depois de `w = buf`, mesmo com `buf == nil`? O que exatamente
muda na representação interna da interface?

### 2. Type assertion vs reflexão

Qual a diferença entre usar `reflect.ValueOf(w).IsNil()` e fazer uma type
assertion `w.(*bytes.Buffer)` para checar se o valor interno é nil? Quando cada
um é apropriado?

### 3. Casos reais

Cite pelo menos **2 situações reais** em código Go onde esse gotcha aparece
com frequência (dica: funções que retornam `error` ou `io.Writer`).
