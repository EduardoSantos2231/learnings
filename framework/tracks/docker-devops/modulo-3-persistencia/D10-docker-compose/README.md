# D10 — Docker Compose Básico

> Template: Implementação | Scaffolding: alto

## Contexto

Gerenciar containers individualmente com `docker run` funciona para 1-2 serviços,
mas fica inviável com 3+. Docker Compose permite descrever múltiplos serviços
em um arquivo YAML e gerenciá-los com um comando.

## Tarefas

### Tarefa 1: Dois serviços manuais (baseline)

Suba manualmente (sem Compose) uma aplicação web Go + Redis:

```bash
docker network create app-net
docker run -d --name redis --network app-net redis:alpine
docker run -d --name app --network app-net -p 8080:8080 \
  -e REDIS_ADDR=redis:6379 minha-app
```

Teste a comunicação. Depois limpe tudo: `docker rm -f redis app && docker network rm app-net`.

### Tarefa 2: docker-compose.yml

Crie um `docker-compose.yml` que faça exatamente a mesma coisa:

```yaml
version: "3.8"
services:
  redis:
    image: redis:alpine
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - REDIS_ADDR=redis:6379
    depends_on:
      - redis
```

Submeta com `docker compose up -d`. Verifique que a aplicação acessa o Redis.

### Tarefa 3: Ciclo de vida

Teste os comandos:
```bash
docker compose ps        # Lista serviços
docker compose logs app  # Logs de um serviço
docker compose stop      # Para sem remover
docker compose start     # Reinicia
docker compose down      # Para e remove containers + network
```

O que `docker compose down` NÃO remove por padrão? Como remover também?

## Restrições

- A aplicação Go deve ser uma API simples que usa Redis para contador de visitas
- Apenas stdlib + `redis` client mínimo (ou HTTP puro para o Redis)
- O Compose file deve criar a rede automaticamente (não declarar networks explícitas)

## Validação

```bash
docker compose up -d
curl localhost:8080  # Deve retornar contador de visitas
curl localhost:8080  # Contador incrementa
docker compose down
docker compose up -d
curl localhost:8080  # Contador resetou (Redis sem volume)
```
