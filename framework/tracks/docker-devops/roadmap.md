# Roadmap — Docker & DevOps

> Track de Docker do zero a stacks production-ready.
> Siga a ordem. O professor anuncia o próximo automaticamente.

## Posição atual

| Campo | Valor |
|-------|-------|
| Módulo atual | 1 — Contêineres e Imagens |
| Último concluído | D4-entrypoint-cmd |
| Próximo desafio | MP1 — Diagnóstico de Containers |
| Próximo formato | Mixed Practice |

---

## Módulo 1 — Contêineres e Imagens

> Scaffolding: alto (comandos sugeridos, dicas de flags)

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| D1 | hello-world | Implementação | `docker run`, imagem vs container, `docker ps`, `docker rm` | ✅ |
| D2 | interactive-shell | Implementação | `-it` flags, `docker exec`, PID 1, isolamento de processos | ✅ |
| D3 | first-dockerfile | Implementação | Dockerfile, FROM, COPY, RUN, CMD, layers, cache | ✅ |
| D4 | entrypoint-cmd | Implementação | ENTRYPOINT, argumentos em runtime, CMD como default args | ✅ |

## Mixed Practice 1

| # | Desafio | Template | Status |
|---|---------|----------|--------|
| MP1 | diagnostico-containers | Mixed Practice | ⬜ |

## Capstone 1 — Container Multi-serviço

| # | Desafio | Template | Conceitos envolvidos | Status |
|---|---------|----------|---------------------|--------|
| C1 | container-multi | Capstone | D1, D2, D3, D4 | ⬜ |

---

## Módulo 2 — Build e Otimização

> Scaffolding: alto → médio

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| D5 | dockerignore-layers | Otimização | `.dockerignore`, ordem de layers, cache de build | ✅ |
| D6 | multi-stage | Implementação | Multi-stage builds, builder vs runtime, redução de tamanho | ⬜ |
| D7 | healthcheck | Implementação | HEALTHCHECK, restart policies, `docker ps` status | ⬜ |

## Mixed Practice 2

| # | Desafio | Template | Status |
|---|---------|----------|--------|
| MP2 | otimizacao-dockerfile | Mixed Practice | ⬜ |

## Capstone 2 — Imagem Otimizada para Produção

| # | Desafio | Template | Conceitos envolvidos | Status |
|---|---------|----------|---------------------|--------|
| C2 | imagem-producao | Capstone | D5, D6, D7 | ⬜ |

---

## Módulo 3 — Persistência, Redes e Compose

> Scaffolding: médio

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| D8 | volumes-bind | Implementação | Bind mount vs volume, `docker volume`, persistência | ✅ |
| D9 | docker-networks | Implementação | Bridge, comunicação entre containers, DNS interno | ⬜ |
| D10 | docker-compose | Implementação | `docker-compose.yml`, services, depends_on | ⬜ |
| D11 | compose-multi | Design | Múltiplos serviços, volumes, networks no compose | ⬜ |

## Capstone 3 — Stack Completa com Compose

| # | Desafio | Template | Conceitos envolvidos | Status |
|---|---------|----------|---------------------|--------|
| C3 | stack-completa | Capstone | D8, D9, D10, D11 | ⬜ |

---

## Legenda

- ✅ Concluído
- ⬜ Pendente
- 🔄 Revisão pendente (spaced repetition)
