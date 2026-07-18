# Mixed Practice 3 — A Estrutura Certa

> Interleaving: para cada cenário, escolha a estrutura de dados ou padrão I/O correto.

## Cenário 1: Histórico de Navegação

Você está implementando um navegador. Precisa das funcionalidades:
- Avançar para a próxima página
- Voltar para a página anterior
- Ao navegar para uma nova página, o histórico "futuro" é descartado

**Pergunta:** Stack, Queue, Doubly LinkedList, ou slice? Por quê?
Implemente com `Push`, `Pop`, `Back`, `Forward`.

## Cenário 2: Processador de Logs Gigantes

Você tem um arquivo de log de 10 GB. Precisa transformar cada linha para
maiúsculas e escrever em outro arquivo. Não cabe na memória.

**Pergunta:** Como processar? `io.Copy`, `bufio.Scanner`, `io.Pipe`, ou memory-map?
Implemente com a abordagem escolhida.

## Cenário 3: Ranking em Tempo Real

Você recebe um stream de pontuações de jogadores. Precisa responder
"top 10" a qualquer momento, com pontuações sendo atualizadas constantemente.

**Pergunta:** BST, heap, slice ordenado, ou map + ordenação periódica?
Implemente com a abordagem escolhida.
