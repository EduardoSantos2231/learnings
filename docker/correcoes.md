# Correções — Docker

## 01 — hello-world

### Imagem vs container
**Sua resposta:** não respondeu explicitamente.
**Correção:** Imagem é o template read-only com sistema de arquivos e instruções (`Dockerfile`). Container é a instância executável da imagem — tem uma camada read-write, ciclo de vida próprio, isolamento de processos, rede e filesystem.

### `docker run` vs `docker start`
**Sua resposta:** `docker run` cria nova instância, `docker start` retoma a aplicação anterior reaproveitando o container.
**Correção:** Quase certo. Detalhe: `docker start` não "retoma o processo de onde parou" — ele roda o entrypoint **do zero**, mas no mesmo filesystem que o container tinha ao ser parado. O arquivo que você criar vai sobreviver porque fica na camada read-write do container, que persiste entre `stop` e `start`.

### `docker ps` vs `docker ps -a`
**Sua resposta:** `-a` lista todos os containers, inclusive os que já pararam.
**Correção:** Correto. `docker ps` só mostra containers **running**. `docker ps -a` mostra todos os estados: running, exited, created, paused.

---

## 02 — interactive-shell

### `-i` vs `-t`
**Sua resposta:** `-i` é interactive (mantém stdin aberto), `-t` é TTY (aloca pseudo-terminal).
**Correção:** Correto. E você testou na prática: `-i` sozinho funciona mas perde formatação (PS1, cores). `-t` sozinho aloca o pseudo-terminal mas sem stdin conectado o bash recebe EOF e morre na hora.

### `docker run -it` vs `docker exec -it`
**Sua resposta:** `run` instancia novo container do zero, `exec` executa comando em container já ativo.
**Correção:** Correto. `docker exec` nunca cria um novo container.

### Caminho errado: `/temp/` em vez de `/tmp/`
**Sua resposta:** Tentou `touch /temp/eu-estou-aqui.txt` e recebeu "No such file or directory".
**Correção:** O diretório temporário no Linux é `/tmp/`, não `/temp/`. Basta trocar: `touch /tmp/eu-estou-aqui.txt`.

### PID 1 — `docker run` vs `docker exec`
**Sua resposta:** Não soube responder — achou que seria a mesma coisa.
**Correção:** Esse é o conceito mais importante do desafio:
- `docker run -it ubuntu bash` → o **bash é o PID 1** do container. Quando você dá `exit`, o PID 1 morre → **container morre junto** (status `exited`).
- `docker exec -it <container> bash` → o bash é um **processo filho** do PID 1. Quando você dá `exit`, só o bash morre. O PID 1 e o container **continuam rodando**.
- Por isso `docker start -ai` funciona depois de um `exit` no `docker run`: você está "religando" o container que morreu. E por isso `docker exec` permite entrar e sair quantas vezes quiser sem matar nada.
