# slice-leak

Visualizar o backing array da slice e entender por que `Pop` precisa zerar o elemento removido.

## Tarefas

### 1 — Crie uma slice com capacidade extra

```go
original := make([]string, 0, 5)
original = append(original, "a", "b", "c", "d")
```

### 2 — Função `pop`

```go
func pop(s []string) ([]string, string)
```

- Remove o último elemento.
- **Não** zera o elemento removido.

### 3 — Função `popSafe`

```go
func popSafe(s []string) ([]string, string)
```

- Remove o último elemento.
- **Zera** o elemento removido (`s[last] = ""`) antes de cortar.

### 4 — Função `dequeue`

```go
func dequeue(s []string) ([]string, string)
```

- Remove o primeiro elemento.

### 5 — main

No `main.go`:
- Crie a slice, chame `pop`, depois print `s[:cap(s)]` — o elemento "removido" ainda está lá.
- Chame `popSafe`, print `s[:cap(s)]` — o elemento foi zerado.
- Chame `dequeue`, print `s[:cap(s)]` — o primeiro elemento ainda está no backing array, mas fora da janela da slice.
- Demonstre que `append` após `dequeue` pode **sobrescrever** o elemento "perdido" (ou reexpor o antigo se houver capacidade).
