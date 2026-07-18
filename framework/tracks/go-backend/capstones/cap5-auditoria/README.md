# Capstone 5 — Auditoria de Código Go

> Template: Capstone | Síntese do Módulo E + Revisão Geral | Sem scaffolding

## Contexto

Você recebeu um repositório Go de 500 linhas escrito por um dev júnior.
Ele "funciona" (passa nos testes), mas está cheio de armadilhas sutis que
causariam bugs em produção. Seu trabalho é auditar e corrigir.

## O código (resumo dos problemas)

O repositório contém uma API de gerenciamento de tarefas. Os problemas incluem:

1. **Nil interface gotcha**: função retorna `*TaskService` nil como `TaskRepository`
   interface. O caller verifica `repo != nil` que é true e chama método → panic.
2. **Slice leak**: função `GetRecent(n int)` retorna `allTasks[len(allTasks)-n:]`
   mantendo referência ao array inteiro de tasks.
3. **Select bloqueante**: função `WaitForTask(id string)` faz select com timeout,
   mas sem case para ctx.Done() — se o contexto for cancelado, vaza goroutine.
4. **errors.Is/As incorreto**: usa `==` para comparar erros wrappeados com `%w`.
   Usa `errors.Is` onde deveria usar `errors.As` para extrair campo do erro.
5. **Middlewares fora de ordem**: CORS depois de Auth → preflight falha.
6. **Mutex copiado por valor**: struct com mutex passada por valor em função.
7. **Goroutine sem cleanup**: goroutine de background sem mecanismo de stop.

## Tarefas

### Fase 1: Identificação (auditoria)

Para cada problema acima:
- Localize no código
- Explique por que é um bug (mecanismo interno)
- Classifique a severidade (panic? memory leak? comportamento incorreto?)
- Proponha a correção

### Fase 2: Correção

Implemente as correções. Cada correção deve ser mínima e focada —
não reescreva o módulo inteiro.

### Fase 3: Prevenção

Para cada classe de bug, sugira uma prática ou ferramenta que o evitaria:
- Linter? Qual regra?
- Padrão de código? Qual?
- Teste? Qual tipo?

### Fase 4: Retrospectiva

Quantos desses bugs você teria cometido há 19 desafios atrás?
O que mudou na sua percepção de código Go?

## Conceitos envolvidos

- Nil interface (type, value) — E1, E2 (nil-interface)
- Slice backing array — C5 (slice-leak)
- Select sem default, ctx.Done() — B5 (select-sem-default)
- errors.Is vs errors.As, %w — C4 (error-is-as)
- Middleware order — D3 (middleware-chain)
- Mutex por valor — B4 (rate-limiter)
- Cleanup goroutine — D1 (cache-ttl)
