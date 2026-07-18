# Capstone 2 — Shell Completo

> Síntese de todos os módulos | Sem scaffolding

## Contexto

Construa um shell Unix funcional: prompt, parser, execução, pipes,
redirecionamento, builtins, variáveis de ambiente e controle de jobs.

## Requisitos

1. **Prompt**: `$ ` com readline básico (suporte a setas ← → e histórico)
2. **Parser**: suporte a aspas simples, duplas, escape `\`, variáveis `$HOME`
3. **Execução**: busca no PATH, builtins (`cd`, `exit`, `pwd`, `export`, `fg`, `bg`, `jobs`)
4. **Pipes**: `cmd1 | cmd2 | cmd3` (múltiplos estágios)
5. **Redirecionamento**: `cmd > file`, `cmd < file`, `cmd >> file`, `cmd 2>&1`
6. **Background**: `cmd &` (rodar em background, imprimir PID)
7. **Controle de jobs**: `jobs` lista processos em background, `fg %1` traz para foreground
8. **Sinais**: Ctrl+C mata foreground, não mata shell; Ctrl+Z suspende foreground
9. **Graceful exit**: `exit` ou Ctrl+D limpa jobs pendentes

## Conceitos envolvidos

- Spawn, wait, exit code — L1 (spawn-wait)
- Sinais — L2 (signals)
- Pipes — L5 (pipes)
- Redirecionamento, dup2 — L6 (redirection)
- Filesystem, PATH — L8 (walk)
- /proc, /proc/self/fd — L9 (proc)
- Tokenizer e AST — L11 (parser)
- Execução com LookPath — L12 (execution)
- Pipeline multi-estágio — L13 (pipes-shell)
- Redirecionamento avançado — L14 (redirect-shell)
