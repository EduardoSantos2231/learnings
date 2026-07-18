# Mixed Practice 1 — Gerenciamento de Processos

> Interleaving: cenários reais de gerenciamento de processos.

## Cenário 1: Process Manager

Você precisa de um programa que mantenha outro processo sempre rodando.
Se o processo filho morrer, reinicia. Se morrer 5x em 10s, desiste.

**Pergunta:** Como implementar? Sinais, wait, ou polling?
Implemente.

## Cenário 2: Graceful Shutdown em Cadeia

Dois processos: worker e logger. O logger depende do worker.
No shutdown: pare o worker primeiro (espere ele terminar),
depois pare o logger. Tudo via sinais.

**Pergunta:** Como orquestrar? Quem envia sinal para quem?
Implemente.

## Cenário 3: Isolamento de Crash

Processo A e B rodam no mesmo processo manager. Se A crashar,
B deve continuar rodando. O manager deve detectar e logar o crash de A.

**Pergunta:** Como isolar? Goroutines ou subprocessos?
Qual a diferença prática?
