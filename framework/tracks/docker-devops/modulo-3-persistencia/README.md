# Módulo 3 — Persistência, Redes e Compose

> Scaffolding: médio

**Objetivo:** Dominar persistência de dados, comunicação entre containers e orquestração com Compose.

## Desafios

| # | Desafio | O que você fez / fará |
|---|---------|----------------------|
| D8 | volumes-bind | Bind mount vs named volume, `docker volume`, persistência entre containers |
| D9 | docker-networks | Bridge, comunicação entre containers por nome, DNS interno do Docker |
| D10 | docker-compose | `docker-compose.yml`, services, depends_on, variáveis de ambiente |
| D11 | compose-multi | Stack multi-serviço: nginx + app + postgres, volumes, networks |

## Conceitos dominados ao final do módulo

- **Bind mount**: monta diretório/arquivo do host no container. Caminho absoluto. Bom para dev (live reload)
- **Named volume**: gerenciado pelo Docker (`docker volume create`). Portável, backups fáceis. Bom para produção
- **`docker volume`**: `create`, `ls`, `inspect`, `rm`, `prune`
- **Bridge network**: rede privada entre containers no mesmo host. DNS interno resolve nome do container → IP
- **Comunicação entre containers**: containers na mesma network se alcançam pelo nome do container (não por IP)
- **`docker-compose.yml`**: descreve múltiplos serviços, volumes, networks em um arquivo
- **services**: cada serviço é um container. Pode ter build ou image, ports, volumes, environment
- **depends_on**: ordem de inicialização (não garante que o serviço está pronto — usar healthcheck + condition)
- **`docker compose up`**: sobe tudo. `-d` para detached. `--build` para rebuild
- **`docker compose down`**: para e remove tudo. `-v` remove volumes também

## Para revisitar

Consulte `spaced-repetition/schedule.md`.
