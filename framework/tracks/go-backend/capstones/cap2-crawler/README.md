# Capstone 2 — Crawler Web Concorrente

> Template: Capstone | Síntese do Módulo B | Sem scaffolding

## Contexto

Você vai construir um crawler web que, dado um domínio, extrai todos os links
internos e reporta o status HTTP de cada página. O crawler deve ser eficiente,
respeitar limites de concorrência e suportar cancelamento graceful.

## Requisitos

### Funcionalidades obrigatórias

1. **Worker pool** (B1): N workers fixos que consomem URLs de um channel e
   fazem HTTP GET. N é configurável via flag.
2. **Rate limiter** (B4): Máximo de M requisições por segundo ao mesmo domínio.
   Use token bucket com `sync.Mutex`.
3. **Fan-in** (B2, B3): Cada worker envia resultados para um channel de saída.
   Um coletor faz merge e dedup (URLs já visitadas).
4. **Context com timeout** (B1, B5): Se o crawl demorar mais que T segundos,
   cancela tudo e reporta resultados parciais.
5. **Graceful shutdown**: Ao receber SIGINT (Ctrl+C), para de enfileirar novas
   URLs, drena os workers e reporta resultados parciais.

### Requisitos não-funcionais

- Sem bibliotecas externas (só `net/http`, `golang.org/x/net/html` para parse)
- Dedup eficiente de URLs (sync.Map ou map com mutex)
- Relatório final: total de URLs, status codes, erros, duração

## Tarefas

### Fase 1: Design (sem código)

Desenhe a arquitetura:
- Quais goroutines? Quais channels? (buffered/unbuffered?)
- Como o rate limiter se integra ao worker pool?
- Como o cancelamento (context + SIGINT) se propaga?
- Fluxo de uma URL: entrada → rate limiter → worker → resultado → coletor

### Fase 2: Implementação

Implemente seguindo seu design.

### Fase 3: Testes

- Teste com um domínio pequeno (ex: servidor HTTP local com 20 páginas)
- Teste de cancelamento: timeout curto → resultados parciais sem panic
- Teste de rate limiting: verifique que o throughput não excede M req/s

### Fase 4: Retrospectiva

- Qual padrão de concorrência foi mais difícil de integrar?
- O design inicial sobreviveu? O que mudou?

## Conceitos envolvidos

- Worker pool, WaitGroup, channels — B1 (worker-pool)
- Fan-in simples — B2 (parallel-query)
- Select multi-canal, graceful shutdown — B3 (fan-in)
- Token bucket, sync.Mutex — B4 (rate-limiter)
- Context cancelamento, select sem default — B5 (select-sem-default)
