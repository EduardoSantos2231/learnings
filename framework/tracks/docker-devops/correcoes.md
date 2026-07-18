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

---

## 03b — entrypoint-vs-cmd

### Q3 — CMD sem ENTRYPOINT: argumento substitui, não concatena

**Sua resposta inicial:** "o argumento é inserido após o padrão"

**Correção:** Sem ENTRYPOINT, o CMD é o comando inteiro. Passar argumentos no `docker run` **substitui** o CMD completamente, não concatena. Exemplo:

```dockerfile
CMD ["cat", "index.html"]
```

- `docker run meu-site` → executa `cat index.html`
- `docker run meu-site /etc/os-release` → tenta executar `/etc/os-release` como comando (falha com permissão negada)

A concatenação só ocorre quando há `ENTRYPOINT` + `CMD`.

### Q5 — ENTRYPOINT para alvo fixo + flags flexíveis

**Sua resposta inicial:** `CMD ["ping", "google.com"]`

**Correção:** Com `CMD` puro, o usuário não consegue passar flags — o argumento substituiria o comando inteiro. Solução correta:

```dockerfile
ENTRYPOINT ["ping"]
CMD ["google.com"]
```

Assim o alvo `google.com` é default substituível, e flags como `-c 4` são injetadas no ping.

### ✅ Acertos

- Entendeu `ENTRYPOINT` como programa fixo e `CMD` como argumento default
- Testou na prática com `cat` e confirmou os comportamentos
- Q1, Q2, Q4 respondidas corretamente de primeira

---

## 04 — dockerignore-layers

### Q1 — "veria mas seriam ignorados"

**Sua resposta:** "o docker não veria os arquivos que inserimos no `.dockerignore` para serem copiados - veria mas seriam ignorados"

**Correção:** Contradição nos termos. O `.dockerignore` não funciona como "vê e ignora" — ele **filtra os arquivos do build context antes** de enviar ao Docker daemon. O Docker nunca "vê" esses arquivos. Mais preciso: o `.dockerignore` remove arquivos do tarball do build context, então o `COPY . .` nem encontra os arquivos ignorados para copiar.

### Q3 — explicação genérica sobre layers

**Sua resposta:** "camadas são construídas uma em cima da outra... Se a camada 3 foi alterada o docker força a re-execução da instrução 4"

**Correção:** A intuição está correta mas genérica. O mecanismo: cada instrução do `Dockerfile` gera uma camada (diff layer). O Docker calcula um **hash** do conteúdo da camada. Se o conteúdo muda (ex: `COPY` traz arquivo modificado), o hash muda → aquela camada e **todas as posteriores** são reconstruídas. É por isso que a ordem importa: colocar instruções estáveis (apt, curl) antes evita rebuild caro.

### Q4 — ENTRYPOINT

**Sua resposta:** "ENTRYPOINT -> defina o ponto de alterada"

**Correção:** "ponto de alterada" é typo de "ponto de entrada". Mais importante: não explicou a função do `ENTRYPOINT`. O `ENTRYPOINT` define o **executável fixo** do container — é o programa que roda quando o container inicia. O `CMD` vira argumento default para esse executável (quando combinados). Sem `ENTRYPOINT`, o `CMD` é o comando inteiro e é **substituído** por argumentos do `docker run`, não concatenado.

### ✅ Acertos

- Dockerfile com `rm -rf /var/lib/apt/lists/*` no mesmo `RUN` do `apt-get install`
- `.dockerignore` com `node_modules/`, `*.log`, `.git` — cobertura correta
- Respostas Q5 e Q6 corretas
- Percebeu sozinho a diferença de ~100KB com e sem `.dockerignore`
- Reorganizou o `COPY` de `COPY . .` para `COPY ["./index.html", "."]` — copia seletivo
- Resposta Q7 correta: mesmo `RUN` evita layers separadas

---

## 05 — volumes-bind

### Dockerfile desnecessário e com problemas

**Problema:** Criou Dockerfile com `apt-get update` sem `apt-get install` — a instrução atualiza o índice de pacotes mas não instala nada, gerando uma camada inútil. Usou `ubuntu:latest` em vez de uma versão específica.

**Correção:** O desafio não exigia Dockerfile — podia usar `docker run -it ubuntu bash` direto. Se fosse criar um, evitar `latest` e só rodar `apt-get` quando necessário.

### Tarefa 1 incompleta

**Problema:** Mostrou criar o arquivo dentro do container mas não provou que ele **some** após remover o container.

**Correção:** O passo que faltou: sair do container → `docker rm` → novo container → mostrar que o arquivo não existe mais. Sem isso, a tarefa não demonstra que dados morrem com o container.

### ✅ Acertos

- Respostas Q2, Q3, Q4 corretas de primeira
- Testou bind mount na prática (ler e escrever do host + container)
- Testou volume nomeado com `docker volume create`
- Q1 após revisão: diferença entre bind mount e volume + cenários corretos
