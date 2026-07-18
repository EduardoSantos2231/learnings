# Capstone 3 — Stack Completa com Docker Compose

> Template: Capstone | Síntese do Módulo 3 | Sem scaffolding

## Contexto

Você vai criar uma stack completa com Docker Compose: uma API Go, PostgreSQL
para dados, Redis para cache, e Nginx como proxy reverso. Tudo com volumes
persistentes e rede interna.

## Requisitos

### Funcionalidades obrigatórias

1. **API Go** (D8, D9): servidor HTTP que aceita POST/GET em `/items`.
   Persiste no PostgreSQL, cacheia no Redis.
2. **PostgreSQL** (D10, D11): banco de dados com volume nomeado para persistência.
   Script de inicialização (`init.sql`) que cria a tabela.
3. **Redis** (D10): cache com volume nomeado. API verifica Redis antes do PostgreSQL.
4. **Nginx** (D11): proxy reverso na porta 80 que encaminha para a API na porta 8080.
5. **Network** (D9): todos os serviços na mesma bridge network. Comunicação por
   nome de serviço (não IP).
6. **Volumes** (D8): PostgreSQL e Redis usam named volumes. Dados sobrevivem a
   `docker compose down` (sem `-v`).

### Requisitos não-funcionais

- `docker compose up` sobe tudo com um comando
- `depends_on` com healthcheck condition (PostgreSQL pronto antes da API)
- `.env` para variáveis de ambiente (senhas, portas)

## Tarefas

### Fase 1: Design

Desenhe a arquitetura:
- Serviços, redes, volumes
- Fluxo de uma requisição: Nginx → API → Redis (cache) → PostgreSQL
- `docker-compose.yml` com todos os serviços

### Fase 2: Implementação

Implemente a API Go (simples, stdlib), o Dockerfile multi-stage, o compose file,
o init.sql, e a config do Nginx.

### Fase 3: Testes

- `docker compose up -d` → todos os serviços saudáveis
- POST `/items` → persistido no PostgreSQL
- GET `/items/{id}` → segunda chamada vem do Redis (cache hit)
- `docker compose down` → containers removidos
- `docker compose up -d` novamente → dados intactos (volumes)
- `docker compose down -v` → dados removidos

### Fase 4: Retrospectiva

## Conceitos envolvidos

- Bind mount vs named volume, persistência — D8
- Bridge network, DNS interno — D9
- docker-compose.yml, services — D10
- Múltiplos serviços, depends_on, healthcheck — D11
