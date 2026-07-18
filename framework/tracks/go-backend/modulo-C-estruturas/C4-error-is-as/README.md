# error-is-as

Consolidar a diferença entre `errors.Is` (valor sentinela) e `errors.As` (tipo).

## Tarefas

### 1 — Defina erros

```go
var ErrNotFound = errors.New("not found")

type ValidationError struct {
    Field string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed on field: %s", e.Field)
}
```

### 2 — Função `processItem`

```go
func processItem(id string) error
```

- Se `id == ""`: retorna `ErrNotFound`
- Se `id == "0"`: retorna `ValidationError{Field: "id"}`
- Se `id == "x"`: retorna `fmt.Errorf("db error: %w", ErrNotFound)`
- Senão: retorna `nil`

### 3 — Função `handleError`

```go
func handleError(err error)
```

- Usa `errors.Is` para detectar `ErrNotFound` e printa "not found"
- Usa `errors.As` para detectar `*ValidationError` e printa o campo
- Mostra que `==` não funciona com o wrapped error (`"db error: %w"`)
- Mostra que `errors.As` precisa de `*ValidationError` (ponteiro), não `ValidationError`

### 4 — main

Crie um `main.go` que chama `processItem` com `""`, `"0"`, `"x"` e `"42"`, passa cada resultado para `handleError`, e printa os resultados.
