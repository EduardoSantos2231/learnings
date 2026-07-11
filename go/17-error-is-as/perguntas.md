## Perguntas

1. Por que `errors.Is(err, ErrNotFound)` funciona com `fmt.Errorf("db error: %w", ErrNotFound)` mas `err == ErrNotFound` não?

2. Por que `errors.As` precisa de um ponteiro (`*ValidationError`) e não do tipo direto (`ValidationError`)? O que acontece se você passar `ValidationError` em vez de `*ValidationError`?

3. Quando você usaria `errors.Is` vs `errors.As`? Dê um exemplo de cada.
