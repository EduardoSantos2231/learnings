# select-sem-default

Fixar de vez o padrão `select` com send no case (não no default) para responder a cancelamento.

## Tarefas

### 1 — `sendOrCancel`

Escreva uma função genérica:

```go
func sendOrCancel[T any](ch chan<- T, val T, ctx context.Context) error
```

- Tenta enviar `val` para `ch`.
- Se `ctx` for cancelado antes do send, retorna `ctx.Err()`.
- **Sem `default`** — o send deve ser um `case` do `select`.

### 2 — `fanInCancelable`

Escreva um merger que lê de N canais e publica num canal de saída:

```go
func fanIn[T any](ctx context.Context, chans ...<-chan T) <-chan T
```

- Cria um canal de saída.
- Para cada canal de entrada, dispara uma goroutine que faz `select` entre `ctx.Done()` e ler do canal.
- Se `ctx` cancelar, todas as goroutines terminam.
- Use `sync.WaitGroup` + goroutine separada para fechar o canal de saída.

Depois:

3. **Compare** com uma versão `fanInNonCancelable` (sem `ctx`). Rode as duas com 2 canais, buffer pequeno, e um `main` que cancela o contexto depois de 1ms. O que acontece com a versão non-cancelable?
