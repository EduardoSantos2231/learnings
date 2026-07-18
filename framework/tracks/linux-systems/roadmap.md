# Roadmap — Linux Systems

> Track de sistemas Linux implementando conceitos do zero em Go com `os`, `os/exec`, `syscall`.
> Siga a ordem. O professor anuncia o próximo automaticamente.

## Posição atual

| Campo | Valor |
|-------|-------|
| Módulo atual | 1 — Processos & Sinais |
| Último concluído | — |
| Próximo desafio | L1-spawn-wait |
| Próximo formato | Implementação |

---

## Módulo 1 — Processos & Sinais

> Scaffolding: alto

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| L1 | spawn-wait | Implementação | Spawn processo, wait, exit code | ⬜ |
| L2 | signals | Implementação | SIGTERM/SIGINT handler, graceful shutdown | ⬜ |
| L3 | daemon | Design | fork, setsid, chdir, desligar stdin | ⬜ |
| L4 | zombies | Debug | Processos zumbis: criar, observar, evitar | ⬜ |

## Mixed Practice 1

| # | Desafio | Template | Status |
|---|---------|----------|--------|
| MP1 | gerenciamento-processos | Mixed Practice | ⬜ |

---

## Módulo 2 — File Descriptors & I/O

> Scaffolding: alto

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| L5 | pipes | Implementação | Pipe entre processos, pai↔filho | ⬜ |
| L6 | redirection | Implementação | dup2, stdin/stdout/stderr | ⬜ |
| L7 | poll-select | Design | Polling com select/epoll | ⬜ |

## Mixed Practice 2

| # | Desafio | Template | Status |
|---|---------|----------|--------|
| MP2 | tee-tail | Mixed Practice | ⬜ |

---

## Módulo 3 — Filesystem

> Scaffolding: médio

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| L8 | walk | Implementação | Walk recursivo, permissões, tipos | ⬜ |
| L9 | proc | Explicação | Leitura de /proc: PID, mem, fds | ⬜ |
| L10 | inode | Debug | Inodes, hardlink vs symlink | ⬜ |

## Mixed Practice 3

| # | Desafio | Template | Status |
|---|---------|----------|--------|
| MP3 | ferramenta-find | Mixed Practice | ⬜ |

## Capstone 1 — tee + tail -f

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| C1 | tee-tail | Capstone | L1-L10 | ⬜ |

---

## Módulo 4 — Mini Shell

> Scaffolding: baixo

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| L11 | parser | Implementação | Tokenizer, AST simples | ⬜ |
| L12 | execution | Implementação | Execução com PATH lookup | ⬜ |
| L13 | pipes-shell | Implementação | Pipes no shell (`cmd1 \| cmd2`) | ⬜ |
| L14 | redirect-shell | Implementação | Redirecionamento (`>`, `<`, `2>`) | ⬜ |

## Capstone 2 — Shell Completo

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| C2 | shell | Capstone | L1-L14 | ⬜ |
