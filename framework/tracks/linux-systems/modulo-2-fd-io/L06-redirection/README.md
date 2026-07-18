# L6 — redirection

> Template: Implementação | Scaffolding: alto

Redirecione stdin, stdout, stderr de um processo filho.

## Tarefas

1. Execute `cat` com stdin vindo de uma string em vez do terminal
2. Execute um comando com stdout redirecionado para arquivo (equivalente a `cmd > file.txt`)
3. Execute um comando com stderr redirecionado para stdout (equivalente a `cmd 2>&1`)
4. Combine: execute `ls /nonexistent /tmp` capturando stdout e stderr no mesmo buffer

Use os campos `cmd.Stdin`, `cmd.Stdout`, `cmd.Stderr`.
