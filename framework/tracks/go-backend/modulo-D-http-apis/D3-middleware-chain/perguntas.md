# D.2 — Perguntas Conceituais

Responda em `respostas.md`.

---

1. Qual a diferença entre `func(http.Handler) http.Handler` (middleware) e `func(http.ResponseWriter, *http.Request)` (handler comum)? Quando usar cada um?

2. O que acontece se um middleware chama `next.ServeHTTP(w, r)` e *depois* escreve algo no `w` (ex: adiciona um header)? A escrita depois do `next` ainda funciona? Por quê?

3. Como o `defer recover()` se comporta em uma chain de middlewares? Se o handler interno panica, qual middleware captura? E se mais de um middleware tiver `defer recover()`?

4. Por que o `Auth` costuma vir *antes* de `Logger` na chain? O que muda se a ordem for invertida?

5. No middleware `CORS`, por que precisamos tratar `OPTIONS` separadamente e retornar `200` sem chamar `next`?

6. O header `Vary: Origin` é recomendado junto com `Access-Control-Allow-Origin: *`? Explique.
