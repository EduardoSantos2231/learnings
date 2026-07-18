# L13 — pipes-shell

> Template: Implementação | Scaffolding: baixo

Implemente pipes no shell: conecte stdout de um comando ao stdin do próximo.

## Tarefas

1. Parseie o operador `|` — divida a linha em múltiplos comandos
2. Para cada par de comandos consecutivos: crie um pipe (`os.Pipe()`), conecte `cmd1.Stdout` ao pipe writer, `cmd2.Stdin` ao pipe reader
3. Execute todos os comandos do pipeline concorrentemente
4. Aguarde todos terminarem e retorne o exit code do último comando

## Exemplo

```
ls -la | grep .go | wc -l
```

Isso deve executar `ls` → pipe → `grep` → pipe → `wc`, com a saída final indo para stdout do shell.

## Validação

- `echo "hello\nworld" | grep hello` → "hello"
- `ls /nonexistent | wc -l` → stderr mostra erro do ls, stdin do wc vazio
