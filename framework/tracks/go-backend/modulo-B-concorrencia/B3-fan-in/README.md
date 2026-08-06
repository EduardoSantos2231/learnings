# B3 — Fan-in

> Implementacao | 60 min | Go stdlib

## Objetivo

Combine resultados de varios produtores em um unico canal e encerre com seguranca.

## Faca

1. Receba uma quantidade variavel de canais de entrada.
2. Encaminhe valores para um canal de saida.
3. Feche a saida quando todos os produtores terminarem.
4. Interrompa o fluxo quando o contexto for cancelado.

## Pronto quando

- Nenhum valor e perdido no encerramento normal.
- O consumidor pode usar `range` na saida.
- Cancelamento nao deixa goroutines bloqueadas.

## Responda

- Onde deve existir o `WaitGroup`?
- Por que o consumidor nao deve fechar a saida?

> Confianca: [1-5]
