# Perguntas — ENTRYPOINT vs CMD

Responda no `respostas.md` (você cria).

1. Qual a diferença fundamental entre `ENTRYPOINT` e `CMD`?

2. O que acontece se você tiver **apenas** `ENTRYPOINT` (sem `CMD`) e rodar `docker run imagem sem parametro algum`? E se rodar `docker run imagem --help`?

3. O que acontece se você tiver **apenas** `CMD` (sem `ENTRYPOINT`) e rodar `docker run imagem`? E `docker run imagem echo oi`?

4. Existe algum jeito de sobrescrever o `ENTRYPOINT` na linha de comando? (Dica: `docker run --entrypoint`)

5. (**Extra**) Você precisa de um container que sempre rode `ping google.com`. O usuário pode passar flags como `-c 4`, mas o alvo (`google.com`) não pode ser mudado. Use `ENTRYPOINT` ou `CMD`? Escreva o comando do Dockerfile.
