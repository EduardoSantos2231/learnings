# Capstone 1 — tee + tail -f

> Síntese dos Módulos 1-3 | Sem scaffolding

## Contexto

Implemente versões production-ready de `tee` e `tail -f` em Go,
usando todos os conceitos de processos, FDs e filesystem.

## Requisitos — tee

1. Lê stdin, escreve em stdout + N arquivos simultaneamente
2. Flag `-a` para append em vez de truncate
3. Flag `-i` para ignorar SIGINT (não interromper com Ctrl+C)
4. Tratamento de erro: se um arquivo falhar, continua nos outros
5. Performance: use buffer adequado, não escreva byte a byte

## Requisitos — tail -f

1. Monitora arquivo e imprime novas linhas em tempo real
2. Detecta rotação de log (arquivo deletado e recriado) — segue o novo
3. Flag `-n N` para mostrar as últimas N linhas antes de começar a seguir
4. Graceful shutdown com SIGTERM
5. Usa inotify via syscall (Linux) — sem polling

## Conceitos envolvidos

- FDs stdin/stdout — L6 (redirection)
- Abertura/escrita de arquivos — L8 (walk)
- Sinais — L2 (signals)
- Inodes — L10 (inode): detectar rotação por inode number
- /proc — L9: verificar FDs abertos
