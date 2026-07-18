# Mixed Practice 1 — Diagnóstico de Containers

> Interleaving: você recebe cenários problemáticos. Diagnostique e corrija sem ajuda.

## Cenário 1: Container Some Misteriosamente

Você roda `docker run --name web nginx` e ele para imediatamente.
`docker ps` não mostra nada. `docker ps -a` mostra `Exited (0)`.

**Pergunta:** O que aconteceu? Como investigar? Como corrigir?
Execute os comandos para provar sua hipótese.

## Cenário 2: PID 1 Trap

Você cria um Dockerfile com `CMD ["/app/start.sh"]`. O script inicia o nginx em
background (`nginx &`) e termina. O container morre logo depois.

**Pergunta:** Por que o container morreu? O que você mudaria no script ou no
Dockerfile para o container continuar vivo?

## Cenário 3: Porta não Exposta

```dockerfile
FROM nginx
COPY index.html /usr/share/nginx/html/
```

Você faz build, roda com `docker run -d -p 8080:80 meu-nginx`, mas `curl localhost:8080`
retorna "connection refused".

**Pergunta:** Liste 3 possíveis causas e como verificar cada uma.
