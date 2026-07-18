# Capstone 1 — Container Multi-serviço

> Template: Capstone | Síntese do Módulo 1 | Sem scaffolding

## Contexto

Você vai criar um container que roda Nginx servindo uma página HTML customizada
e um script de healthcheck que monitora o Nginx. Tudo em uma imagem otimizada.

## Requisitos

### Funcionalidades obrigatórias

1. Dockerfile que parte de `nginx:alpine` (imagem base otimizada)
2. Copia um `index.html` customizado (tema: página de status do servidor)
3. ENTRYPOINT como `nginx` e CMD como `-g daemon off;`
4. Suporte a argumentos em runtime: `docker run ... meu-nginx -t` para testar config
5. Script de entrada que valida configuração antes de iniciar

### Requisitos não-funcionais

- Imagem final < 50MB
- Máximo 5 layers além do FROM
- `.dockerignore` filtrando arquivos desnecessários

## Tarefas

### Fase 1: Design

Desenhe a estrutura: Dockerfile, index.html, script de entrada, `.dockerignore`.
Explique cada decisão de layer order.

### Fase 2: Implementação

Construa a imagem e teste:
- `docker run -d -p 8080:80 meu-nginx` → página carrega
- `docker exec <id> nginx -t` → config válida
- `docker run --rm meu-nginx -t` → testa config e sai

### Fase 3: Publicação (simulada)

- Liste as camadas da imagem: `docker history`
- Meça o tamanho: `docker images`
- Documente como publicaria no Docker Hub

### Fase 4: Retrospectiva

- O que o Módulo 1 te preparou bem?
- Onde você sentiu lacunas?

## Conceitos envolvidos

- Imagem vs container, `docker run`, `docker ps`, `docker rm` — D1
- `docker exec`, PID 1 — D2
- Dockerfile, FROM, COPY, CMD, layers — D3
- ENTRYPOINT, argumentos em runtime — D4
