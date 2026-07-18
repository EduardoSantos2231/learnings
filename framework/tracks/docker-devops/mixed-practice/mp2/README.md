# Mixed Practice 2 — Otimização de Dockerfile

> Interleaving: você recebe Dockerfiles reais. Identifique problemas e otimize.

## Cenário 1: Imagem Inchada

```dockerfile
FROM ubuntu:22.04
RUN apt-get update
RUN apt-get install -y python3
RUN apt-get install -y curl
COPY . /app
RUN rm -rf /var/lib/apt/lists/*
WORKDIR /app
CMD ["python3", "app.py"]
```

**Pergunta:** Liste TODOS os problemas deste Dockerfile. Corrija.
Qual o impacto de cada problema no tamanho final e no cache?

## Cenário 2: Build Lento

Você tem uma aplicação Go. Toda vez que altera um arquivo `.go`, o build baixa
dependências novamente (2 minutos). O `COPY . .` está antes do `RUN go mod download`.

**Pergunta:** Como reordenar para cachear dependências? Reescreva o Dockerfile.
Qual a economia de tempo esperada?

## Cenário 3: Multi-stage Desnecessário?

Um colega diz que multi-stage builds são sempre melhores. Você discorda.

**Pergunta:** Dê um exemplo concreto onde single-stage é melhor que multi-stage.
Justifique com números: tamanho da imagem, tempo de build, complexidade.
