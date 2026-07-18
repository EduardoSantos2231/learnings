# 02 — Interactive Shell no Ubuntu

## Objetivo

Rodar um container Linux interativo e entender isolamento, ciclo de vida e `docker exec`.

## Tarefas

### 1. Primeiro container interativo
```bash
docker run -it ubuntu bash
```
- O que `-i` e `-t` fazem separadamente?
- Rode `ls`, `cat /etc/os-release`, `whoami` dentro do container.
- Crie um arquivo: `touch /tmp/eu-estou-aqui.txt`
- Saia com `exit`.

### 2. O container morreu?
- `docker ps -a` — qual o status do container?
- Rode `docker start -ai <container_id_ou_nome>` — o arquivo `/tmp/eu-estou-aqui.txt` ainda existe?

### 3. docker exec vs docker run
- Inicie o Ubuntu em background: `docker run -dit --name meu-ubuntu ubuntu bash`
- Entre nele sem criar novo container: `docker exec -it meu-ubuntu bash`
- Rode `ps aux` dentro do `exec` — o que aparece?
- Saia com `exit`. O container `meu-ubuntu` ainda está rodando? (`docker ps`)

### 4. Comparação
- Rode `docker run -it ubuntu bash` em outro terminal.
- Dentro dele, rode `ps aux`. Compare com a saída do `exec`.

## Desafio extra
Descubra o que acontece se você rodar `docker exec -it meu-ubuntu bash` e depois matar o processo bash com `kill -9 $$`. O container morre junto?
