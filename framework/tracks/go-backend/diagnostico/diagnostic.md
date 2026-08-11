# Diagnostico pratico — Go Backend

> Execute sem consulta. O objetivo e escolher o proximo bloco, nao obter nota.

## Faca

Implemente um pequeno servico HTTP em Go que:

1. Aceite `POST /items` com `{"name":"x"}`.
2. Retorne `400` para JSON invalido ou nome vazio.
3. Retorne `201` com um ID crescente para cada item valido.
4. Use um mutex para proteger o armazenamento.
5. Processe em paralelo uma lista de tarefas de validacao, com limite de 3 e cancelamento por `context.Context`.

## Restricoes

- Use apenas a biblioteca padrao.
- Escreva testes para entradas validas, invalidas e concorrentes.
- Rode `go test -race ./...`.

## Pronto quando

- O servidor atende os tres cenarios HTTP.
- O teste com `-race` passa.
- Voce consegue explicar onde aparecem erros, interfaces, mutex e contexto.

## Checkpoint Docker

- Crie uma imagem com multi-stage.
- Rode a API com `docker run -p`.
- Verifique o endpoint pelo host e veja os logs.
- Explique qual processo e o PID 1.

## Responda

- Qual parte voce implementaria diferente em producao?
- Qual conceito foi mais fraco durante a implementacao?

> Confianca: [1-5]
