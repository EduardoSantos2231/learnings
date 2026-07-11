## Perguntas

1. Depois de `dequeue` (que faz `s = s[1:]`), por que `s[:cap(s)]` ainda mostra o elemento removido?

2. Por que o `Pop` vaza memória (referência retida) mas o `Dequeue` não precisa de limpeza explícita no mesmo sentido? Os dois vazam?

3. Se você fizer `s = s[1:]` e depois `s = append(s, "x")`, o que acontece com o valor que estava em `s[0]` antes do shift?
