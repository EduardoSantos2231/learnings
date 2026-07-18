# 04 — Dockerignore e Ordem de Layers

## Objetivo

Entender como o `.dockerignore` reduz o tamanho da imagem e como a ordem das instruções no `Dockerfile` afeta o cache de build.

## Setup inicial

```
04-dockerignore-layers/
├── node_modules/       ← já criado (lixo simulado)
│   ├── heavy-dep/
│   │   └── index.js
│   └── .cache/
│       └── cachefile
├── Dockerfile          ← você cria
├── .dockerignore       ← você cria
└── index.html          ← você cria
```

A pasta `node_modules/` já existe com arquivos grandes simulando dependências desnecessárias no build.

---

## Tarefas

### 1. Crie `index.html`

Um HTML simples qualquer (1 parágrafo basta).

### 2. Crie um `Dockerfile`

Com:

- `FROM ubuntu:latest`
- `RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*`
- `WORKDIR /app`
- `COPY . .`
- `CMD ["cat", "index.html"]`

### 3. Build sem `.dockerignore`

```
docker build -t meu-site .
docker images meu-site
```

Anote o tamanho.

### 4. Crie `.dockerignore`

Ignore `node_modules/`, `*.log`, `.git`.

Rebuild e compare:

```
docker build -t meu-site .
docker images meu-site
```

O que mudou no tamanho e no tempo de build?

### 5. Explore as layers

```
docker history meu-site
```

Veja o tamanho de cada layer. A layer do `COPY . .` diminuiu depois do `.dockerignore`?

### 6. Mexa na ordem

O `COPY . .` copia TUDO — se qualquer arquivo mudar, essa layer e todas as seguintes perdem o cache.

Reorganize o `Dockerfile`:

1. Instalar dependências do sistema (`apt`) primeiro
2. Copiar apenas arquivos que não mudam muito
3. Só no final copiar o que muda sempre (`index.html`)

Dica: use `COPY` seletivo em vez de `COPY . .`.

Faça uma mudança no `index.html` e rebuild. Quantas layers foram reaproveitadas?

---

## Desafio extra (opcional)

Adicione um script `script.sh` no container que usa `curl` para baixar algo. Coloque o `COPY script.sh` na ordem certa para não invalidar o cache do `COPY index.html`.

---

## Perguntas

Responda em `respostas.md`.
