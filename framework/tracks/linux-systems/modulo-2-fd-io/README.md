# Módulo 2 — File Descriptors & I/O

> Scaffolding: alto | Go `os`, `syscall`

**Objetivo:** Dominar file descriptors, pipes entre processos e redirecionamento — o fundamento de como o shell conecta comandos.

## Desafios

| # | Desafio | O que você constrói |
|---|---------|-------------------|
| L5 | pipes | Pipe entre pai e filho: escrever em um lado, ler do outro |
| L6 | redirection | Redirecionar stdin/stdout/stderr com dup2 |
| L7 | poll-select | Polling de múltiplos FDs com select — sem bloqueio |

## Conceitos ao final do módulo

- File descriptor: inteiro que referencia arquivo/pipe/socket aberto
- FDs padrão: 0 = stdin, 1 = stdout, 2 = stderr
- `os.Pipe()` — cria pipe (reader, writer)
- `syscall.Dup2(oldfd, newfd)` — duplica FD para um número específico
- Redirecionamento: `cmd.Stdout = file` em Go
- `Cmd.ExtraFiles` — passar FDs extras para processo filho
- select/epoll: esperar por eventos em múltiplos FDs sem bloquear
- I/O não blocante: `syscall.SetNonblock(fd, true)`
