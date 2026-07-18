# Mixed Practice 2 — Padrão de Concorrência

> Interleaving: para cada cenário, escolha o padrão correto e justifique.

## Cenário 1: Processamento de Lote

Você tem 10.000 imagens para redimensionar. Cada imagem leva ~100ms.
O servidor tem 8 CPUs. Você quer processar tudo o mais rápido possível.

**Pergunta:** Worker pool com N workers fixos, fan-out ilimitado, ou outra abordagem?
Implemente com a abordagem escolhida.

## Cenário 2: API Gateway

Sua API recebe 1000 req/s de clientes externos. Cada requisição consulta
um serviço interno que suporta no máximo 50 req/s.

**Pergunta:** Como proteger o serviço interno? Worker pool, rate limiter, semáforo?
Implemente com a abordagem escolhida.

## Cenário 3: Agregador de APIs

Você precisa consultar 3 APIs externas diferentes e retornar o resultado
combinado. Cada API tem latência variável (50ms a 2s). Você quer responder
em no máximo 1.5s.

**Pergunta:** Fan-in com timeout, select sequencial, ou outra abordagem?
Implemente com a abordagem escolhida.
