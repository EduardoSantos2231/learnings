# Módulo 3 — DNS Resolver

> Scaffolding: médio | Go `net` + encoding binário

**Objetivo:** Construir um resolver DNS do zero — montar queries no wire format,
fazer resolução iterativa a partir dos root servers, cachear respostas.

## Desafios

| # | Desafio | O que você constrói |
|---|---------|-------------------|
| R9 | dns-wire | Construir query DNS manualmente: header, question, encoding de nomes |
| R10 | dns-resolver | Resolver DNS iterativo: root → TLD → authoritative nameserver |
| R11 | dns-cache | Cache de respostas com TTL do record, invalidação |

## Conceitos ao final do módulo

- DNS wire format: header (12 bytes), question, answer, authority, additional
- Encoding de nomes DNS: labels com length prefix, compressão de ponteiros
- Tipos de record: A (IPv4), AAAA (IPv6), NS (nameserver), CNAME (alias), MX (mail)
- Resolução iterativa vs recursiva
- Root servers (a.root-servers.net ... m.root-servers.net)
- TTL: tempo de vida do record, cache expira após TTL
- `encoding/binary` — BigEndian para fields do header DNS
