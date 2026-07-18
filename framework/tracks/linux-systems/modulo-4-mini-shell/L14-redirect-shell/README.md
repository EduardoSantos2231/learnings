# L14 — redirect-shell

> Template: Implementação | Scaffolding: baixo

Implemente redirecionamento no shell: `>`, `<`, `>>`, `2>`, `2>&1`.

## Tarefas

1. Parseie operadores de redirecionamento na linha de comando:
   - `>` — redireciona stdout para arquivo (trunca)
   - `>>` — redireciona stdout para arquivo (append)
   - `<` — redireciona stdin a partir de arquivo
   - `2>` — redireciona stderr para arquivo
   - `2>&1` — redireciona stderr para stdout
2. Use `os.OpenFile` com as flags corretas: `O_WRONLY|O_CREAT|O_TRUNC` para `>`, `O_WRONLY|O_CREAT|O_APPEND` para `>>`, `O_RDONLY` para `<`
3. Atribua os arquivos abertos aos campos `cmd.Stdin`, `cmd.Stdout`, `cmd.Stderr`
4. Combine pipes com redirecionamento: `ls -la | grep go > out.txt`

## Exemplos

| Comando | Comportamento |
|---------|--------------|
| `echo hello > file.txt` | Escreve "hello" em file.txt |
| `echo world >> file.txt` | Adiciona "world" ao final |
| `cat < file.txt` | Lê de file.txt como stdin |
| `ls /err 2> errors.txt` | Stderr vai para errors.txt |
| `ls /tmp /err > out.txt 2>&1` | Ambos stdout e stderr em out.txt |

## Validação

Verifique conteúdo de arquivos após cada operação e compare com o comportamento do shell real.
