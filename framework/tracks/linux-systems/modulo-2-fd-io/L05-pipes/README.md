# L5 — pipes

> Template: Implementação | Scaffolding: alto

Crie pipes entre processos pai e filho.

## Tarefas

1. Use `os.Pipe()` para criar um pipe
2. Pai escreve `"hello from parent"` no pipe, filho lê e imprime
3. Filho escreve `"hello from child"` no pipe, pai lê e imprime
4. Trate fechamento do pipe: leitor recebe `io.EOF` quando escritor fecha

## Validação

Execute o programa e veja ambas as mensagens impressas.
