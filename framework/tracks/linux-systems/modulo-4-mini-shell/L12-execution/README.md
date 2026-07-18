# L12 — execution

> Template: Implementação | Scaffolding: baixo

Execute comandos com PATH lookup e builtins.

## Tarefas

1. Implemente PATH lookup: dado um nome de comando, procure o executável em cada diretório do `$PATH` (use `os/exec.LookPath`)
2. Implemente builtins que não executam binário externo:
   - `cd <dir>` — muda o diretório do shell
   - `pwd` — imprime o diretório atual
   - `exit [code]` — sai do shell com exit code opcional
3. Para comandos externos: execute usando `os/exec`, passando os argumentos parseados
4. Trate erros: comando não encontrado, permissão negada, diretório inexistente

## Validação

- `ls -la` → mesma saída do shell real
- `cd /tmp && pwd` → `/tmp`
- `nonexistent` → mensagem de erro "command not found"
- `exit 42` → shell termina com código 42
