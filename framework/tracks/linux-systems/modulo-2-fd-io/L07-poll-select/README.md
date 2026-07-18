# L7 — poll-select

> Template: Design | Scaffolding: alto

Faça polling de múltiplos file descriptors sem bloquear.

## Tarefas

1. **Design:** como esperar por dados em múltiplos FDs simultaneamente sem spawnar uma goroutine por FD?
2. **Pesquise:** a syscall `select()` e seu equivalente em Go (`syscall.Select`)
3. **Implemente:** crie 2 pipes, escreva neles com delays, use select para ler de quem estiver pronto primeiro
4. **Adicione timeout:** select deve retornar após 5s se nenhum FD estiver pronto
