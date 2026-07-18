# D7 — Healthcheck

> Template: Implementação | Scaffolding: alto

## Contexto

Containers em produção precisam informar se estão saudáveis. O Docker tem uma
instrução `HEALTHCHECK` que testa a saúde do container periodicamente e expõe
o status via `docker ps`.

## Tarefas

### Tarefa 1: Adicione HEALTHCHECK

Pegue o Dockerfile do desafio D6 (multi-stage) e adicione:

```dockerfile
HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
  CMD wget -qO- http://localhost:8080/ || exit 1
```

Faça build, rode o container e verifique:
```bash
docker ps  # Deve mostrar "(healthy)" após alguns segundos
```

### Tarefa 2: Simule uma falha

Com o container rodando, mate o processo da aplicação:
```bash
docker exec <id> pkill server
```

Observe o que acontece com:
- `docker ps` (status)
- Os logs: `docker logs <id>`
- O container é reiniciado? Por que (não)?

### Tarefa 3: Restart policy

Rode o mesmo container com `--restart=on-failure`:
```bash
docker run -d --restart=on-failure --name health-test sua-imagem
```

Repita a tarefa 2. O container reinicia agora? Por quê?

Teste também `--restart=always` e `--restart=unless-stopped`.
Qual a diferença prática entre eles?

## Validação

```bash
docker run -d --name hc-test sua-imagem
sleep 10
docker ps --filter name=hc-test --format '{{.Status}}'
# Deve conter "(healthy)"

docker exec hc-test pkill server
sleep 10
docker ps -a --filter name=hc-test --format '{{.Status}}'
# Deve conter "(unhealthy)" ou "Exited"
```
