# L9 — proc

> Template: Explicação | Scaffolding: médio

Explore e explique o filesystem `/proc`.

## Tarefas

1. Liste todos os processos rodando lendo `/proc` (filtre diretórios numéricos)
2. Para cada processo: leia `/proc/[pid]/status` para obter nome, PID, PPID, memória (VmRSS)
3. Para o processo atual: leia `/proc/self/fd/` para listar os file descriptors abertos
4. **Explique (máx 300 palavras):** o que é `/proc`? Como ele difere de um filesystem regular? Por que os arquivos não têm tamanho real?
5. **Bônus:** leia `/proc/self/maps` e explique os mapeamentos de memória

## Validação

Compare sua listagem de processos com `ps aux`.
