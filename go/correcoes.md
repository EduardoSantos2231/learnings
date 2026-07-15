# Correções — Go

## 14 — nil interface gotcha (B.3)

### Q1 — Representação interna da interface

**Sua resposta:** disse que `w` armazena o valor de `buf` (ponteiro nil), então `w` não é nil.

**Correção:** Descreveu o *sintoma* mas não o *mecanismo*. Uma interface em Go é um par `(type, value)`:
- Antes de `w = buf`: `w = (nil, nil)` → interface é nil
- Depois de `w = buf`: `w = (*bytes.Buffer, nil)` → type não é nil, então interface **não** é nil

Sem esse par na cabeça, o gotcha parece mágica. Com ele, é óbvio.

### Q3 — Casos reais

**Sua resposta:** não soube citar, só um palpite sobre "descuido de inicialização".

**Correção:** Dois casos clássicos que já pegaram todo mundo (Netflix, Uber, etc.):

1. **`error` com ponteiro nil:** Função que retorna `*MyError` — quando não há erro, retorna `(*MyError)(nil)` em vez de `nil`. O `if err != nil` vira `true` mesmo sem erro, porque o `type` do par não é nil.

2. **`io.Reader`/`io.Writer` nil:** Um `*bytes.Buffer` nil passado pra função que espera `io.Reader`. A interface não é nil, mas o valor interno é nil → panic ao tentar ler.

### Código: `safeWrite` com `IsNil()`

**Problema:** `reflect.ValueOf(w).IsNil()` **panica** se o tipo interno não for nilable (ex: struct, int, string). É seguro só para: `Ptr`, `Chan`, `Map`, `Slice`, `Func`, `Interface`.

**Correção:** checar o `Kind` antes:
```go
v := reflect.ValueOf(w)
switch v.Kind() {
case reflect.Ptr, reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface:
    return v.IsNil()
}
return false
```

### Código: Comentário impreciso

**Linha 25-26:** "buffer possui sim o método Write" — quem implementa `Write` é `*bytes.Buffer` (pointer receiver), não `bytes.Buffer` struct.

---

## 15 — Middleware Chain (D.2)

### Código: `chain()` off-by-one

**Problema:** `for i := len(middlewares); i > 0; i--` — `len(middlewares)` é 5, índices válidos 0-4. Acessar `middlewares[5]` dá panic.

**Correção:** `for i := len(middlewares) - 1; i >= 0; i--` — começa do último índice válido e inclui o índice 0.

**Tipo:** Falta de atenção — off-by-one clássico. Precisou de 2 iterações pra acertar.

### Código: `chain()` ignorado no `ListenAndServe`

**Problema:** Criou `finalHandler` com `chain(mux, middlewares...)` mas passou `loggerMiddleware(jsonMiddleware(finalHandler))` pro servidor — `chain()` foi descartado, middlewares rodavam duplicados ou não rodavam.

**Correção:** `http.ListenAndServe(":8080", finalHandler)` — usa o handler montado pelo `chain()` direto.

**Tipo:** Falta de atenção — codificou a solução (`chain`) e depois a ignorou.

### Código: `recovererMiddleware` escrevia 500 em toda request

**Problema:** O `w.WriteHeader(500)` + `Encode` estavam **fora** do `if recover() != nil` — executavam em toda requisição, inclusive nas bem-sucedidas.

**Correção:** Mover `WriteHeader` + `Encode` pra **dentro** do `if err := recover(); err != nil { ... }`.

**Tipo:** Lógica condicional — não percebeu que o código depois do `if` roda sempre.

### Código: Auth não validava token

**Problema 1:** `strings.Split(token, "")` separa cada caractere individualmente. `bearerTkn[2]` era o 3º caractere, não "admin123".

**Correção 1:** `strings.Split(token, " ")` e `bearerTkn[1]`.

**Problema 2:** `token == ""` só dava `return` sem escrever 401 — cliente recebia resposta vazia.

**Correção 2:** `w.WriteHeader(http.StatusUnauthorized)` antes do `return`.

**Tipo:** Falta de atenção — confundiu separador vazio com espaço e esqueceu de retornar status.

### Código: `strconv.Atoi` erro como 500 (repetido do D.1)

**Problema:** `http.StatusInternalServerError` em parsing de ID inválido — erro é do cliente.

**Correção:** `http.StatusBadRequest`.

**Tipo:** Erro recorrente — mesmo ponto já anotado no D.1. Indica que a correção anterior não fixou.

### Código: Campo JSON não exportado no recoverer (repetido do D.1)

**Problema:** `message string` (minúsculo) — `json.Encode` ignora campos não exportados.

**Correção:** `Message string`.

**Tipo:** Erro recorrente — mesmo erro do D.1, repetido 3x naquele desafio.

### Código: CORS sem `Access-Control-Allow-Headers`

**Problema:** Preflight `OPTIONS` não informava quais headers eram permitidos — navegador bloqueava requisições com `Authorization`.

**Correção:** Adicionar `w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")`.

**Tipo:** Conhecimento de protocolo — faltou saber que o preflight precisa desse header.

### Respostas conceituais

#### Q1 — Handler vs Middleware

**Problema:** Respondeu sobre `HandleFunc` typecasting em vez da pergunta.

**Correção:** A diferença é:
- `func(w, r)` = handler, processa **uma** requisição
- `func(next) Handler` = middleware, **envolve** um handler adicionando comportamento antes/depois

#### Q2 — Escrever depois do `next.ServeHTTP`

**Problema:** Resposta vaga sobre "desempilhar". Não explicou o mecanismo.

**Correção:** Se o handler interno chamou `w.Write()`, o status code e headers já foram enviados (congelados). Escrever headers depois do `next` é silenciosamente ignorado pelo `net/http`. Se o handler não escreveu nada, funciona.

#### Q3 — Múltiplos `defer recover()` na chain

**Problema:** Não respondeu "e se mais de um middleware tiver defer recover?".

**Correção:** A pilha de defers é LIFO. O recover **mais interno** (mais próximo do handler que panica) roda primeiro. Se ele recuperar, os externos não veem o panic. Se ele não recuperar (`recover()` retorna nil), o próximo da pilha (mais externo) tenta.

#### Q4 — Auth antes de Logger (resposta ausente na 1ª versão)

**Problema:** Não respondeu na primeira versão.

**Correção:** Auth antes de Logger economiza recursos (não loga requisições rejeitadas) e reduz superfície de ataque (usuário não autenticado não passa do auth).

#### Q6 — `Vary: Origin` com `Access-Control-Allow-Origin: *`

**Problema:** Resposta explicou Vary corretamente mas não concluiu se é recomendado.

**Correção:** Não é recomendado. `Vary: Origin` serve para variar o cache por origem quando o servidor reflete o valor do `Origin` request header (ex: `Access-Control-Allow-Origin: https://example.com`). Com `*` (permite todas as origens), a resposta não varia — `Vary: Origin` é desnecessário e prejudica o cache.

---

## 17 — error-is-as (Reparo 3.2)

### Código: Type param errado no `errors.AsType`

**Problema:** `errors.AsType[ValidationError](err)` — a type assertion `err.(ValidationError)` falha porque `processItem` retorna `&ValidationError{}` (ponteiro).

**Correção:** `errors.AsType[*ValidationError](err)` — o tipo precisa casar: `*ValidationError` em vez de `ValidationError`.

### Respostas conceituais

#### Q3 — `errors.Is` vs `errors.As`

**Problema inicial:** Disse que `Is` serve para "saber se houve um tipo de erro" — isso é o papel do `As`.

**Correção:** `errors.Is` verifica **valores** sentinela (`ErrNotFound`, `io.EOF`). `errors.As` extrai **tipos** (`*ValidationError`).

---

## 18 — slice-leak (Reparo 3.3)

### Código: `unsafePop[cap(unsafePop)]` em vez de `unsafePop[:cap(unsafePop)]`

**Problema:** Tentou acessar índice `cap(unsafePop)` (=5) diretamente, mas o slice tem `len=3` — acesso fora do comprimento causa panic.

**Correção:** `s[:cap(s)]` re-expande a janela do slice até a capacidade do backing array. É diferente de `s[cap(s)]` (índice único, estoura o len).

### Código: `dequeue` implementado como `popSafe`

**Problema:** Removeu o último elemento em vez do primeiro — copiou a função `popSafe` sem adaptar a lógica.

**Correção:** `dequeue` deve fazer `s[1:]` e retornar `s[0]`.

### Código: `popSafe` retornava valor zerado

**Problema:** `s[lastIndex] = ""` era executado **antes** de salvar o valor, então o retorno era sempre `""`.

**Correção:** Salvar o valor primeiro: `lastVal := s[lastIndex]`, depois zerar, depois retornar.

### Respostas conceituais

#### Q3 — GC e backing array

**Problema:** Achou que o GC coleta elementos individuais do backing array.

**Correção:** O GC coleta o **array inteiro** quando ninguém mais o referencia. Enquanto a slice existir, todo o backing array (incluindo elementos "removidos") permanece alocado. O `append` que estoura a capacidade aloca um novo array — só então o antigo pode ser coletado.
