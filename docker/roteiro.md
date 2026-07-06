# Roteiro de Estudos — Docker

> **Modo de uso:** Este documento serve para retomar a sessão quando a janela de contexto for esgotada.
> Basta abri-lo e informar ao seu "professor" (opencode) qual Desafio deseja continuar.

---

## 1. Perfil do Estudante

| Campo | Valor |
|---|---|
| **Background** | TypeScript / frontend |
| **Nível atual** | Zero absoluto em Docker |
| **Estilo de aprendizado** | Guiado por conceitos — prefere direcionamento e correção a receber comandos prontos |
| **Foco** | Domínio prático de Docker para desenvolvimento e deploy |
| **Regra de ouro** | Nada de sudo cego — entender cada flag antes de rodar |
| **Regra do professor** | Nunca entregar comandos prontos. Conduzir com perguntas, apontar erros, sugerir caminhos. |
| **Regra do aluno** | Executar, explicar a saída e responder às perguntas conceituais. |

---

## 2. Roadmap

### Módulo 1 — Fundamentos (do zero)

| # | Desafio | Conceitos | Status |
|---|---------|-----------|--------|
| 01 | hello-world | `docker run`, imagem vs container, `docker ps`, `docker rm`, `docker start` | ✅ |
| 02 | interactive-shell | `-i`/`-t`, `docker exec`, PID 1, isolamento de processos, `docker start -ai` | ✅ |
| 03 | my-first-dockerfile | `Dockerfile` (`FROM`, `WORKDIR`, `COPY`, `RUN`, `CMD`), `docker build`, layers | ✅ (pende desafio extra ENTRYPOINT) |
| 03b | entrypoint-vs-cmd | `ENTRYPOINT`, CMD como default args, argumentos em runtime | ⬜ |
| 04 | dockerignore-layers | `.dockerignore`, ordem de layers, cache de build, redução de tamanho | ⬜ |
| 05 | volumes-bind | Bind mount vs volume, persistência de dados, `docker volume` | ⬜ |
| 06 | networking | `docker network`, bridge, comunicação entre containers, DNS interno | ⬜ |
| 07 | compose-simple | `docker-compose.yml` com 2 serviços, `depends_on`, variáveis de ambiente | ⬜ |

### Módulo 2 — Aplicação real

| # | Desafio | Conceitos | Status |
|---|---------|-----------|--------|
| 08 | compose-fullstack | Dockerizar app Go + banco + cache, `.env`, `healthcheck` | ⬜ |
| 09 | healthchecks-restart | `HEALTHCHECK`, restart policies (`always`, `on-failure`), resource limits | ⬜ |
| 10 | production-ready | Multi-stage build, usuário não-root, `docker scout`, imagem mínima | ⬜ |

---

## 3. Perguntas Teóricas por Desafio

### 01 — hello-world
- Qual a diferença entre imagem e container?
- `docker ps` vs `docker ps -a`?
- `docker run` vs `docker start`?
- Dá pra entrar num container depois que ele finaliza?

### 02 — interactive-shell
- `-i` vs `-t` — o que cada um faz? E os dois juntos?
- `docker run -it` vs `docker exec -it`?
- O que acontece com o PID 1 no `exit` do `run` vs `exec`?

---

## 4. Erros Recorrentes

| # | Erro | Desafio | Por que acontece |
|---|---|---|---|
| 1 | `/temp/` em vez de `/tmp/` | 02 | Confundir diretório temporário do Linux |
| 2 | ... | ... | ... |
