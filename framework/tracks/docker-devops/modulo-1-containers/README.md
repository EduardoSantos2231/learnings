# Módulo 1 — Contêineres e Imagens

> Scaffolding: alto

**Objetivo:** Entender o que é um container, como interagir com ele e como criar suas próprias imagens.

## Desafios

| # | Desafio | O que você fez |
|---|---------|---------------|
| D1 | hello-world | Rodou o hello-world, analisou output, entendeu imagem vs container, docker ps -a, docker rm |
| D2 | interactive-shell | Rodou container interativo, entendeu -i vs -t, docker exec, PID 1, isolamento de processos |
| D3 | first-dockerfile | Escreveu seu primeiro Dockerfile, entendeu FROM, COPY, RUN, CMD, layers e cache |
| D4 | entrypoint-cmd | Diferenciou ENTRYPOINT de CMD, argumentos em runtime, CMD como default args |

## Conceitos dominados ao final do módulo

- **Imagem**: template imutável com camadas em stack (union filesystem)
- **Container**: instância runtime de uma imagem com camada writable
- **`docker run`**: cria e inicia um container. Flags: `--name`, `-d`, `-p`, `-e`, `--rm`
- **`docker ps`**: lista containers rodando. `-a` inclui parados
- **`docker rm`**: remove containers parados. `-f` força remoção em execução
- **`-i` (interactive)**: mantém STDIN aberto
- **`-t` (tty)**: aloca pseudo-terminal (formatação, cores, sinais)
- **`docker exec`**: roda comando em container já em execução
- **PID 1**: primeiro processo do container. Se ele morre, o container morre
- **Dockerfile**: instruções em ordem → cada uma gera uma layer
- **FROM**: imagem base
- **COPY**: copia arquivos do host para a imagem
- **RUN**: executa comando durante o build
- **CMD**: comando padrão na inicialização do container
- **ENTRYPOINT**: comando fixo; CMD vira argumento padrão
- **ENTRYPOINT + CMD**: `ENTRYPOINT ["echo"]` + `CMD ["hello"]` → `docker run img world` executa `echo world`

## Para revisitar

Consulte `spaced-repetition/schedule.md`.
