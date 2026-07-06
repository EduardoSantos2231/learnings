# 03 — Meu Primeiro Dockerfile

## Objetivo

Escrever seu próprio `Dockerfile`, fazer build da imagem e rodar um container a partir dela.

## Tarefas

### 1. Crie um `Dockerfile` simples

Crie o arquivo `Dockerfile` dentro de `03-my-first-dockerfile/` com:

- **`FROM`** — use `ubuntu:latest` como imagem base
- **`RUN`** — atualize os pacotes (`apt-get update`) e instale `curl`
- **`WORKDIR`** — defina `/app` como diretório de trabalho
- **`COPY`** — copie um arquivo `index.html` (você cria) para `/app/`
- **`CMD`** — use `["cat", "index.html"]` para exibir o conteúdo

Crie também um `index.html` simples ao lado do `Dockerfile`.

### 2. Build e execute

```bash
docker build -t meu-site .
docker run --rm meu-site
```

- O que a flag `--rm` faz?
- O que acontece se você rodar `docker run -it meu-site bash`? Você consegue acessar o bash? O curl está instalado?

### 3. Explore as layers

- Rode `docker history meu-site` — o que cada linha representa?
- Faça uma pequena mudança no `index.html` e rebuild. O que o Docker **reaproveita** do cache? O que ele **executa de novo**? Por quê?

## Perguntas

1. Qual a diferença entre `CMD` e `ENTRYPOINT`?
2. O que acontece se você tiver **dois** `CMD` no mesmo `Dockerfile`?
3. Para que serve o `WORKDIR`? O que acontece se você não definir um?
4. Por que a ordem das instruções no `Dockerfile` importa para o cache de build?
5. O que a flag `--rm` faz e por que ela é útil para testes?

## 🔄 Desafio extra (pendente)

Modifique o `Dockerfile` para usar `ENTRYPOINT` em vez de `CMD`. Rode:

```bash
docker run meu-site
docker run meu-site /etc/os-release
```

O que mudou no comportamento?
