# nil-interface-revisao (reativado)

Retomar o exercício abandonado B.3 — nil interface gotcha.

## Tarefas

### 1 — Crie um programa que demonstra o gotcha

```go
var w io.Writer          // (type=nil, value=nil) → w == nil: true
var buf *bytes.Buffer    // nil pointer
w = buf                  // (type=*bytes.Buffer, value=nil) → w != nil: true!
w.Write([]byte("hello")) // panic!
```

### 2 — Implemente `safeWrite`

```go
func safeWrite(w io.Writer, data []byte) (int, error)
```

- Verifica se `w` é nil (tanto interface nil quanto pointer nil dentro da interface).
- Usa `reflect.ValueOf(w).IsNil()` — mas **antes** verifica o Kind para não panicar em tipos não nilable.

### 3 — Explique

Escreva em `respostas.md`: por que `w != nil` mesmo com `buf` sendo nil?

Referência: exercício original em `14-nil-interface-gotcha/`.
