# Mixed Practice 2 — Cliente HTTP

> Interleaving: use seu servidor HTTP do Módulo 2. Agora escreva o cliente.

## Cenário 1: Download com Progresso

Escreva um cliente HTTP que baixa um arquivo grande e mostra progresso
(bytes baixados / total). Use apenas `net.Dial` + seu parser HTTP.

## Cenário 2: Cliente Keep-Alive

Escreva um cliente que reusa a conexão TCP para fazer 5 requisições
sequenciais ao mesmo servidor. Meça o tempo com e sem keep-alive.

## Cenário 3: Cliente Pipeline

HTTP/1.1 permite pipelining: enviar múltiplas requisições sem esperar
resposta. Implemente um cliente que envia 3 requisições em pipeline.
O servidor do R6 consegue responder? O que acontece?
