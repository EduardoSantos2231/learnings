# Correções — Go

## 14 — nil interface gotcha (B.3)

### Q1 — Representação interna da interface

**Sua resposta:** disse que `w` armazena o valor de `buf` (ponteiro nil), então `w` não é nil.

**Correção:** Descreveu o *sintoma* mas não o *mecanismo*. Uma interface em Go é um par `(type, value)`:
- Antes de `w = buf`: `w = (nil, nil)` → interface é nil
- Depois de `w = buf`: `w = (*bytes.Buffer, nil)` → type não é nil, então interface **não** é nil

Sem esse par na cabeça, o gotcha parece mágica. Com ele, é óbvio.

### Q3 — Casos reais

**Sua resposta:** não soube citar, só um palpite sobre "descuido de inicialização".

**Correção:** Dois casos clássicos que já pegaram todo mundo (Netflix, Uber, etc.):

1. **`error` com ponteiro nil:** Função que retorna `*MyError` — quando não há erro, retorna `(*MyError)(nil)` em vez de `nil`. O `if err != nil` vira `true` mesmo sem erro, porque o `type` do par não é nil.

2. **`io.Reader`/`io.Writer` nil:** Um `*bytes.Buffer` nil passado pra função que espera `io.Reader`. A interface não é nil, mas o valor interno é nil → panic ao tentar ler.

### Código: `safeWrite` com `IsNil()`

**Problema:** `reflect.ValueOf(w).IsNil()` **panica** se o tipo interno não for nilable (ex: struct, int, string). É seguro só para: `Ptr`, `Chan`, `Map`, `Slice`, `Func`, `Interface`.

**Correção:** checar o `Kind` antes:
```go
v := reflect.ValueOf(w)
switch v.Kind() {
case reflect.Ptr, reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface:
    return v.IsNil()
}
return false
```

### Código: Comentário impreciso

**Linha 25-26:** "buffer possui sim o método Write" — quem implementa `Write` é `*bytes.Buffer` (pointer receiver), não `bytes.Buffer` struct.
