# Mixed Practice 3 — Ferramenta `find` Simplificada

> Interleaving: use os conceitos de filesystem para construir um `find`.

## Cenário: find

Implemente uma ferramenta que busca arquivos no filesystem:
- `-name "*.go"` — filtra por nome (glob)
- `-type f|d|l` — filtra por tipo
- `-mtime -7` — arquivos modificados nos últimos 7 dias
- `-size +1M` — arquivos maiores que 1MB
- `-exec <cmd> {} \;` — executa comando para cada arquivo encontrado

Comece com -name e -type. Adicione os outros progressivamente.
Compare com o `find` real: falta algo?
