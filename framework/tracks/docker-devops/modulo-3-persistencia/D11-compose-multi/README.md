# D11 — Docker Compose Multi-serviço

> Template: Design | Scaffolding: médio

## Contexto

Você vai projetar (e depois implementar) uma stack completa com Docker Compose:
API Go + PostgreSQL + Redis + Nginx como proxy reverso. Diferente do D10,
aqui você precisa pensar na arquitetura antes de codar.

## Tarefas

### Tarefa 1: Design da arquitetura

Desenhe (texto ou diagrama ASCII) a arquitetura da stack:

- Quais serviços?
- Quais redes? (back-end para app+db+redis, front-end para nginx+app?)
- Quais volumes? (PostgreSQL, Redis precisam persistir?)
- Como o Nginx descobre o endereço da API?
- Como a API descobre o endereço do PostgreSQL e Redis?

### Tarefa 2: docker-compose.yml

Escreva o `docker-compose.yml` completo com:
- **app**: build da API Go, depende de db e redis (com healthcheck condition)
- **db**: PostgreSQL 16-alpine, volume nomeado, variáveis de ambiente para user/pass/db
- **cache**: Redis 7-alpine, volume nomeado
- **proxy**: Nginx alpine, porta 80 mapeada, config customizada como volume bind mount
- **networks**: back-end (app+db+cache), front-end (proxy+app)

### Tarefa 3: Config do Nginx

Crie `nginx.conf` que faz proxy reverso para `app:8080`:

```nginx
events {}
http {
    server {
        listen 80;
        location / {
            proxy_pass http://app:8080;
        }
    }
}
```

### Tarefa 4: Inicialização do banco

Crie `init.sql` que roda na primeira inicialização do PostgreSQL:
```sql
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

Monte como volume em `/docker-entrypoint-initdb.d/`.

## Restrições

- API Go em stdlib (como nos desafios anteriores)
- Imagem do app usa multi-stage build (D6)
- PostgreSQL com volume para persistência (D8)
- Redis com volume para persistência (D8)
- Nginx config via bind mount (D8)
- Comunicação por nome de serviço (D9)

## Validação

```bash
docker compose up -d
# Esperar o PostgreSQL iniciar
sleep 5
curl localhost/items       # Lista vazia
curl -X POST localhost/items -H 'Content-Type: application/json' -d '{"name":"teste"}'
curl localhost/items       # Lista com 1 item
docker compose down
docker compose up -d
curl localhost/items       # Lista ainda tem 1 item (volume persistiu)
docker compose down -v     # Remove volumes também
```
