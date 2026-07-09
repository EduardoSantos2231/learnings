# 03b — ENTRYPOINT vs CMD

## Objetivo

Entender na prática a diferença entre `ENTRYPOINT` e `CMD` no Dockerfile, e como argumentos em runtime interagem com cada um.

## Pré-requisito

Você vai usar o mesmo `Dockerfile` e `index.html` do desafio 03.
Se não existir mais, recrie rapidinho:

- `FROM ubuntu:latest`
- `RUN apt-get update && apt-get install -y curl`
- `WORKDIR /app`
- `COPY index.html .`
- `CMD ["cat", "index.html"]`

Crie também um `index.html` simples ao lado do `Dockerfile`.

---

## Tarefas

### 1. Build com CMD

Garanta que o Dockerfile original (com `CMD`) está pronto e build:

```bash
docker build -t meu-site .
docker run meu-site
docker run meu-site /etc/os-release
```

Anote: o que aconteceu na segunda execução? Por quê?

### 2. Troque CMD por ENTRYPOINT

Altere o `Dockerfile`: substitua `CMD` por `ENTRYPOINT`, mantendo o mesmo comando.

Exemplo:
```dockerfile
ENTRYPOINT ["cat", "index.html"]
```

```bash
docker build -t meu-site .
docker run meu-site
docker run meu-site /etc/os-release
```

O que mudou? Por quê?

### 3. ENTRYPOINT + CMD como argumento padrão

Agora separe as responsabilidades:

```dockerfile
ENTRYPOINT ["cat"]
CMD ["index.html"]
```

```bash
docker build -t meu-site .
docker run meu-site
docker run meu-site /etc/os-release
```

Compare com os resultados anteriores. Qual a diferença?

### 4. Teste com argumento inválido

```bash
docker run meu-site --flag-inexistente
```

O que acontece? Por quê?

---

## Resumo esperado

Preencha uma tabela mental como esta:

| Dockerfile | `docker run meu-site` | `docker run meu-site /etc/os-release` |
|------------|----------------------|---------------------------------------|
| `CMD ["cat", "index.html"]` | cat index.html | ? |
| `ENTRYPOINT ["cat", "index.html"]` | cat index.html | ? |
| `ENTRYPOINT ["cat"]` + `CMD ["index.html"]` | cat index.html | ? |

