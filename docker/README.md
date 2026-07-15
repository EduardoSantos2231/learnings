# Estudos Docker

Roteiro completo em [roteiro.md](./roteiro.md).

## Desafios

| # | Projeto | Conceitos |
|---|---------|-----------|
| 01 | hello-world | `docker run`, imagem vs container, `docker ps`, `docker rm` |
| 02 | interactive-shell | `-it` flags, `docker exec`, PID 1, isolamento de processos |
| 03 | my-first-dockerfile | Dockerfile, FROM, COPY, RUN, CMD, layers, cache de build |
| 03b | entrypoint-vs-cmd | ENTRYPOINT, argumentos em runtime, CMD como default args |
| 04 | dockerignore-layers | `.dockerignore`, ordem de layers, cache de build, redução de tamanho |
| 05 | volumes-bind | Bind mount vs volume, `docker volume`, persistência de dados |
