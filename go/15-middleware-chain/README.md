# D.2 — Middleware Chain

Continuação do D.1 (Products API). Vamos adicionar uma cadeia de middlewares usando o padrão `func(http.Handler) http.Handler`.

---

## Tarefa 1 — Base + Logger + Recoverer

1. Copie a estrutura do D.1 (ou recrie do zero): `POST /products`, `GET /products`, `GET /products/{id}`, `PUT /products/{id}`, `DELETE /products/{id}`.
2. Crie um middleware `Logger` que loga no stdout: método, path, e duração da requisição.
3. Crie um middleware `Recoverer` que captura panic com `defer recover()`, loga o erro e retorna `500 Internal Server Error`.
4. Aplique `Recoverer(Logger(mux))` — o recoverer deve ser o mais externo.

## Tarefa 2 — Auth + CORS

1. Crie um middleware `Auth` que lê o header `Authorization: Bearer <token>`.
   - Token válido: `"admin123"` (hardcoded).
   - Se inválido ou ausente → `401 Unauthorized`.
   - Pule a verificação para `GET /products` (lista pública).
2. Crie um middleware `CORS` que adiciona os headers:
   - `Access-Control-Allow-Origin: *`
   - `Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS`
   - `Access-Control-Allow-Headers: Content-Type, Authorization`
   - Para `OPTIONS` (preflight), retorne `200 OK` sem chamar o próximo handler.

## Tarefa 3 — Chain final

Componha a chain nesta ordem:

```
CORS → Auth → Recoverer → Logger → mux
```

Teste com `curl` ou seu HTTP client favorito — garanta que:
- panic não derruba o servidor
- requisição sem token recebe 401 (exceto GET /products)
- preflight OPTIONS recebe 200 com os headers CORS
- toda requisição é logada com método, path e duração

---

## Dicas

- A assinatura padrão: `func next(next http.Handler) http.Handler { return http.HandlerFunc(...) }`
- Duração: `time.Since(start)` ou `time.Now().Sub(start)`
- `defer recover()` só funciona na mesma goroutine — o middleware que envolve o handler é o lugar certo
- A ordem importa: o middleware mais externo executa primeiro no request e último no response
