# L10 — inode

> Template: Debug | Scaffolding: médio

Inspecione inodes, hardlinks e symlinks.

## Tarefas

1. Crie um arquivo, obtenha seu número de inode (`syscall.Stat_t.Ino`)
2. Crie um hardlink para o arquivo, verifique mesmo número de inode e nlink incrementado
3. Delete o original, verifique que o hardlink ainda funciona (dados não são deletados até nlink chegar a 0)
4. Crie um symlink, verifique número de inode diferente, delete o original, verifique que o symlink quebra
5. Escreva um programa que encontra todos os hardlinks de um mesmo arquivo escaneando um diretório por inodes iguais

## Verificação

Use `ls -li` e `stat` para confirmar cada etapa.
