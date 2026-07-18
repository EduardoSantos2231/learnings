# L8 — walk

> Template: Implementação | Scaffolding: médio

Percorra recursivamente uma árvore de diretórios e reporte informações.

## Tarefas

1. Use `filepath.WalkDir` para percorrer um diretório
2. Para cada arquivo: imprima caminho, tamanho, permissões (formato rwx), data de modificação
3. Classifique: arquivo regular, diretório, symlink, socket, pipe, dispositivo
4. Adicione filtros: pule arquivos ocultos (começam com `.`), pule diretórios que batem com padrão, profundidade máxima N

## Validação

Rode em `/tmp`, compare com `find /tmp -ls`.
