# L11 — parser

> Template: Implementação | Scaffolding: baixo

Tokenize e parseie uma linha de comando no estilo shell.

## Tarefas

1. Implemente um tokenizer que separa a linha de comando em tokens respeitando espaços, aspas simples e aspas duplas
2. Trate escapes: `\n`, `\t`, `\\`, `\"`, `\'`
3. Construa uma AST simples: `type Command struct { Name string; Args []string }`
4. Trate edge cases: aspas não fechadas (erro), strings vazias, múltiplos espaços consecutivos

## Exemplos

| Input | Tokens |
|-------|--------|
| `ls -la /tmp` | `["ls", "-la", "/tmp"]` |
| `echo "hello world"` | `["echo", "hello world"]` |
| `echo 'it\'s fine'` | `["echo", "it's fine"]` |
| `grep "a b" -r` | `["grep", "a b", "-r"]` |

## Validação

Crie uma tabela de casos de teste com input → output esperado e valide cada um.
