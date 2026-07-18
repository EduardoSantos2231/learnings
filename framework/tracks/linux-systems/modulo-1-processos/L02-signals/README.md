# L2 — signals

> Template: Implementação | Scaffolding: alto

Manipule sinais do SO para graceful shutdown.

## Tarefas

1. Crie um programa que faz trabalho em loop (ex: contar segundos)
2. Capture SIGINT (Ctrl+C) e SIGTERM com `os/signal.Notify`
3. Ao receber sinal: pare de aceitar trabalho novo, termine o atual, imprima resumo, saia
4. Demonstre: execute o programa, envie `kill -TERM` de outro terminal, observe o graceful shutdown

## Teste

- Inicie → envie SIGTERM → cleanup → saia
- Inicie → envie SIGKILL → morte imediata (sem cleanup)
