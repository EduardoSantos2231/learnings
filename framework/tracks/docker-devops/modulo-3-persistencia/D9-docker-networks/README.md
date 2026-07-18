# D9 — Docker Networks

> Template: Implementação | Scaffolding: alto

## Contexto

Containers isolados não se comunicam. O Docker oferece redes virtuais (bridge)
para conectar containers. Você vai criar uma rede, conectar containers e fazê-los
se comunicar usando nomes de serviço (DNS interno do Docker).

## Tarefas

### Tarefa 1: Rede isolada

Crie uma rede bridge e comprove que containers fora dela não se alcançam:

```bash
docker network create minha-rede
docker run -d --name container-a --network minha-rede alpine sleep 3600
docker run -d --name container-b alpine sleep 3600
```

Do container-a, tente pingar container-b. Funciona? Por quê?

### Tarefa 2: Comunicação por nome

Recrie container-b na mesma rede:

```bash
docker rm -f container-b
docker run -d --name container-b --network minha-rede alpine sleep 3600
```

Agora tente do container-a: `docker exec container-a ping container-b`.
Funciona? Por quê? Como o container-a descobre o IP do container-b?

### Tarefa 3: Portas expostas vs rede interna

Rode um nginx na rede:

```bash
docker run -d --name web --network minha-rede nginx:alpine
```

Do container-a: `docker exec container-a wget -qO- web`.
Funciona? O nginx está respondendo? Em qual porta?

Agora tente do host: `curl localhost`. Funciona? Por quê?
O que você precisa fazer para o host acessar o nginx?

### Tarefa 4: Inspeção

```bash
docker network inspect minha-rede
```

Identifique no output:
- Os containers conectados
- Os IPs de cada container
- O driver da rede
- A subnet

## Validação

```bash
docker network create teste
docker run -d --name a --network teste alpine sleep 300
docker run -d --name b --network teste alpine sleep 300
docker exec a ping -c 2 b  # Deve funcionar
docker exec b ping -c 2 a  # Deve funcionar
docker network rm teste     # Remove rede e desconecta containers
```
