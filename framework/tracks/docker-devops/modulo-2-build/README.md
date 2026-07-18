# Módulo 2 — Build e Otimização

> Scaffolding: alto → médio

**Objetivo:** Dominar cache de build, reduzir tamanho de imagens e garantir resiliência.

## Desafios

| # | Desafio | O que você fez / fará |
|---|---------|----------------------|
| D5 | dockerignore-layers | Entendeu `.dockerignore`, ordem de layers, cache de build, hashing |
| D6 | multi-stage | Separar build de runtime com multi-stage, reduzir imagem Go de 800MB para <20MB |
| D7 | healthcheck | HEALTHCHECK, restart policies, status do container no `docker ps` |

## Conceitos dominados ao final do módulo

- **`.dockerignore`**: filtra arquivos ANTES do contexto ir para o daemon
- **Cache de build**: cada instrução gera hash da instrução + conteúdo. Se o hash bate, reusa layer cacheada. Alterar uma linha invalida TODAS as layers abaixo
- **Ordem de layers**: instruções que mudam pouco primeiro (ex: `apt-get update`); código que muda muito por último (ex: `COPY . .`)
- **Multi-stage builds**: múltiplos `FROM` no mesmo Dockerfile. `COPY --from=builder` copia artefatos entre stages. Imagem final só tem o runtime, não o toolchain
- **HEALTHCHECK**: instrução no Dockerfile que define como verificar saúde do container. `docker ps` mostra `(healthy)` ou `(unhealthy)`
- **Restart policies**: `--restart=always|unless-stopped|on-failure` controla o que acontece quando o container para
- **Limpeza de cache**: `rm -rf /var/lib/apt/lists/*` na mesma layer do `apt-get install` para não deixar lixo em layers separadas

## Para revisitar

Consulte `spaced-repetition/schedule.md`.
