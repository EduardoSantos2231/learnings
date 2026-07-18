# L1 — spawn-wait

> Template: Implementação | Scaffolding: alto

Spawne um processo e capture sua saída.

## Tarefas

1. Use `os/exec` para executar `ls -la`, capture stdout e stderr
2. Obtenha o exit code do processo
3. Spawne com variáveis de ambiente customizadas
4. Defina um timeout: se o processo não terminar em 5s, mate-o e retorne erro

## Validação

- Rode com `ls` (sucesso)
- Rode com `ls /nonexistent` (erro + stderr)
- Rode com `sleep 10` (timeout)
