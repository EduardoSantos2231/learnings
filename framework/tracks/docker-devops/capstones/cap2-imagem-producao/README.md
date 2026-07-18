# Capstone 2 — Imagem Otimizada para Produção

> Template: Capstone | Síntese do Módulo 2 | Sem scaffolding

## Contexto

Você tem uma aplicação Go (servidor HTTP simples) e precisa criar uma imagem
Docker otimizada para produção: mínima, com healthcheck e build eficiente.

## Requisitos

### Funcionalidades obrigatórias

1. Multi-stage build: stage 1 compila o binário Go (`golang:1.22-alpine`),
   stage 2 roda só o binário (`alpine:3.20`)
2. `.dockerignore` otimizado: exclui `.git`, `*.md`, `node_modules`, `.cache`
3. Ordem de layers otimizada: `go.mod`/`go.sum` primeiro, depois source
4. HEALTHCHECK: endpoint `/health` que retorna 200
5. Usuário não-root: `USER 1000:1000` no stage final
6. Limpeza de cache: `rm -rf /var/cache/apk/*` na mesma layer da instalação

### Requisitos não-funcionais

- Imagem final < 15MB
- Build cacheado: alterar `.go` não re-baixa dependências
- Testar healthcheck: matar o processo do app e ver `docker ps` mostrar unhealthy

## Tarefas

### Fase 1: Design

### Fase 2: Implementação

### Fase 3: Métricas

Compare sua imagem com uma versão single-stage (ubuntu + Go toolchain):
- Tamanho
- Tempo de build (cold vs cached)
- Número de layers
- Vulnerabilidades (`docker scout` ou inspeção manual)

### Fase 4: Retrospectiva

## Conceitos envolvidos

- `.dockerignore`, ordem de layers, cache — D5
- Multi-stage builds — D6
- HEALTHCHECK, restart policies — D7
