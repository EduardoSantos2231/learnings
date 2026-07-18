# D6 — Multi-stage Builds

> Template: Implementação | Scaffolding: alto

## Contexto

Você tem uma aplicação Go que compila para um binário de ~10MB, mas a imagem
Docker atual tem ~800MB porque inclui o toolchain Go inteiro e o Ubuntu.

Seu objetivo: reduzir essa imagem para <20MB usando multi-stage builds.

## Especificação

O build acontece em dois estágios:
1. **Stage 1 (builder)**: imagem `golang:1.22-alpine`, compila o binário
2. **Stage 2 (runtime)**: imagem `alpine:3.20`, copia SÓ o binário do stage 1

## Tarefas

### Tarefa 1: Single-stage (baseline)

Crie um Dockerfile single-stage simples que:
- Usa `golang:1.22-alpine`
- Copia o código
- Compila com `go build -o /app/server .`
- Roda com `CMD ["/app/server"]`

Faça build e anote o tamanho: `docker images | grep <nome>`.

### Tarefa 2: Multi-stage

Reescreva o Dockerfile como multi-stage:
- Stage 1 (`AS builder`): compila
- Stage 2: `FROM alpine:3.20`, copia binário com `COPY --from=builder`
- (Opcional) Adicione certificados CA: `RUN apk add --no-cache ca-certificates`

Faça build e compare o tamanho com a versão single-stage.

### Tarefa 3: Análise

Responda:
- Quanto menor ficou a imagem? (em MB e %)
- Por que a imagem single-stage é tão grande?
- O que exatamente o `COPY --from=builder` copia?
- O stage 1 (builder) aparece na imagem final? Como provar?

## Restrições

- A aplicação Go deve compilar e rodar em ambas as versões
- Use `alpine`, não `scratch` (para facilitar debugging)
- Apenas Dockerfile — sem docker-compose

## Validação

```bash
docker build -t app-single -f Dockerfile.single .
docker build -t app-multi -f Dockerfile.multi .
docker images | grep app-
# app-multi deve ser pelo menos 90% menor que app-single
docker run --rm app-multi
# Deve imprimir o output esperado da aplicação
```

## Aplicação Go de exemplo

```go
// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from multi-stage build!")
	})
	http.ListenAndServe(":8080", nil)
}
```
