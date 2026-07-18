# Módulo 4 — Mini Shell

> Scaffolding: baixo | Go `os`, `os/exec`, `syscall`

**Objetivo:** Construir um shell funcional do zero: parser de comandos, execução,
pipes e redirecionamento. O capstone final da track.

## Desafios

| # | Desafio | O que você constrói |
|---|---------|-------------------|
| L11 | parser | Tokenizer + parser de linha de comando (argumentos, quotes, escapes) |
| L12 | execution | Executar comandos com PATH lookup, builtins (cd, exit, pwd) |
| L13 | pipes-shell | Pipes: `cmd1 | cmd2 | cmd3` conectando stdout→stdin |
| L14 | redirect-shell | Redirecionamento: `>`, `<`, `>>`, `2>`, `2>&1` |

## Conceitos ao final do módulo

- Tokenizer: split por espaços, respeitando aspas simples/duplas
- AST simples: Command (nome, args, stdin, stdout, stderr)
- Pipeline: sequência de comandos conectados por pipes
- PATH lookup: buscar executável em cada diretório do $PATH
- Builtins: comandos implementados pelo shell (não executáveis externos)
- `os/exec.LookPath` — buscar executável no PATH
- `Cmd.StdinPipe`, `Cmd.StdoutPipe` — conectar pipes entre comandos
- `os.OpenFile` com flags: O_RDONLY, O_WRONLY, O_CREAT, O_TRUNC, O_APPEND
- Redirecionamento: abrir arquivo e atribuir a cmd.Stdin/Stdout/Stderr
