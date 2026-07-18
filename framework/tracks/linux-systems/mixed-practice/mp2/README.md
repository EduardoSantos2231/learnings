# Mixed Practice 2 — Pipe e Redirecionamento

> Interleaving: construa um `tee` simplificado.

## Cenário: tee

Implemente um programa que lê stdin e escreve simultaneamente em stdout
e em um arquivo. Igual ao comando `tee arquivo.txt`.

Depois estenda: suporte a `tee -a` (append em vez de truncate).
Depois: suporte a múltiplos arquivos: `tee file1.txt file2.txt`.

## Bônus: tail -f

Implemente `tail -f arquivo.txt`: monitora o arquivo e imprime novas linhas
conforme são adicionadas. Sem polling (use inotify ou semelhante).
