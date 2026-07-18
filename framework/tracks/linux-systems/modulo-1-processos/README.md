# Módulo 1 — Processos & Sinais

> Scaffolding: alto | Go `os/exec`, `os/signal`

**Objetivo:** Entender como o SO gerencia processos: criação, espera, sinais, e o ciclo de vida de um processo Unix.

## Desafios

| # | Desafio | O que você constrói |
|---|---------|-------------------|
| L1 | spawn-wait | Spawn processo, capturar stdout/stderr, wait pelo exit code |
| L2 | signals | Handler de SIGTERM/SIGINT, graceful shutdown com cleanup |
| L3 | daemon | Processo daemon: fork, setsid, chdir, desligar stdin/stdout/stderr |
| L4 | zombies | Criar zumbis de propósito, observar com `ps`, evitar com wait |

## Conceitos ao final do módulo

- `os/exec.Cmd` — criar e executar processos
- fork + exec: o que acontece em cada etapa
- wait/waitpid: coletar exit status, evitar zumbis
- Exit code: 0 = sucesso, não-zero = erro
- Sinais: SIGTERM (terminação graciosa), SIGINT (Ctrl+C), SIGKILL (não capturável), SIGHUP (hangup)
- `os/signal.Notify` — capturar sinais em Go
- Processo daemon: desacoplar do terminal, nova sessão, diretório raiz
- Zumbi: processo filho que terminou mas o pai ainda não coletou o exit status
