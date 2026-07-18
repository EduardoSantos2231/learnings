# Módulo 3 — Filesystem

> Scaffolding: médio | Go `os`, `io/fs`, `syscall`

**Objetivo:** Explorar o filesystem Linux programaticamente: percorrer diretórios,
ler /proc, inspecionar inodes, entender hardlinks e symlinks.

## Desafios

| # | Desafio | O que você constrói |
|---|---------|-------------------|
| L8 | walk | Walk recursivo: listar arquivos, permissões, tamanhos, tipos |
| L9 | proc | Parser de /proc: listar processos, memória, FDs abertos |
| L10 | inode | Inspecionar inodes: número, links, diferença hardlink vs symlink |

## Conceitos ao final do módulo

- `os.ReadDir`, `filepath.WalkDir` — percorrer diretórios
- `os.FileInfo` — Name, Size, Mode, ModTime, IsDir, Sys()
- Permissões Unix: rwx para owner/group/other
- `/proc` — filesystem virtual com info de processos e kernel
- `/proc/<pid>/status`, `/proc/<pid>/fd/`, `/proc/<pid>/maps`
- Inode: identificador único do arquivo no filesystem
- Hardlink: mesmo inode, mesmo arquivo, não funciona entre filesystems
- Symlink: atalho para caminho, inode diferente, quebra se alvo deletado
- `syscall.Stat_t` — campos: Ino (inode number), Nlink (hardlink count)
