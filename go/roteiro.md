# Roteiro de Estudos — Go para Entrevistas

> **Modo de uso:** Este documento serve para retomar a sessao quando a janela de contexto for esgotada.
> Basta abri-lo e informar ao seu "professor" (opencode) qual Modulo/Exercicio deseja continuar.

---

## 1. Perfil do Estudante

| Campo | Valor |
|---|---|
| **Background** | TypeScript / frontend |
| **Nivel atual** | Intermediario-inicial (ja domina goroutines, channels, `select`, `sync.WaitGroup`, `context`) |
| **Estilo de aprendizado** | Guiado por conceitos — prefere direcionamento e correcao a receber codigo pronto |
| **Foco** | Entrevistas para vagas Go (junior/pleno) |
| **Regra de ouro** | Apenas biblioteca padrao (sem libs externas) |
| **Regra do professor** | Nunca entregar código pronto. Conduzir com perguntas, apontar contradições, sugerir caminhos. A excelência do aluno é o objetivo — ele quem escreve, explica e decide entre alternativas. |
| **Regra do aluno** | Codar, explicar a própria implementação e escolher entre A ou B quando o professor apresentar caminhos. |

---

## 2. Desafios Completos

### 2.1 — Currency Converter (CLI + struct + sentinel errors)

**Arquivos:** `currency_conversor/main.go`, `currency_conversor/utils/typos.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `flag`, struct `Converter` com `map[string]float64`, taxas baseadas em USD |
| **Erros** | Sentinel errors com `errors.Is`, `fmt.Errorf("...: %w")`, stderr, exit codes 1 e 2 |
| **Padrao final** | `NewConverter()` retorna struct com rates pre-carregado; `Convert(from, to, amount)` faz a aritmetica e valida moedas |

**✅ O que funcionou bem:**
- Separacao clara: `main.go` (orquestracao + flags) vs `utils/typos.go` (dominio)
- `errors.Is` usado corretamente para sentinel errors
- Writer correto: `fmt.Fprintf(os.Stderr, ...)` para erros, `fmt.Printf` para saida normal

**⚠️ Pontos a melhorar:**
- Variavel `outputErr` nao segue convencao Go (`err` basta)
- `validateFlags` agrupa `From` e `To` num unico `errEmptyCurrency` — nao distingue qual campo veio vazio
- `Convert` so informa qual moeda e invalida no caso do `To`; o `From` invalido so diz "You must provide a valid currency"
- Nome do arquivo `typos.go` nao reflete o conteudo

**🔁 Comportamento observado:**
- Tendencias de modelar dominio com tipos ricos (ex: `type USD struct`, `type BRL struct`) — padrao OO do TS
- Variavel `flagsError` vs `outputErr` — esquecimento ao renomear (ocorreu uma vez)
- `flag.Parse()` esquecido na primeira tentativa

**📌 Sugestao para novos desafios:** Nomear variaveis como a comunidade (simples, curtas) e evitar construir hierarquias de tipos quando `map` resolve.

---

### 2.2 — BankAccount (pointer receiver + erros tipados)

**Arquivos:** `bank_account/main.go`, `bank_account/utils/definitions.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | Pointer receiver, struct privada (`bankAccount`), constructor `NewBankAccount`, `errors.As`, erro tipado |
| **Erros** | `ErrorInsufficientFunds{Balance, Amount}` implementa `Error()`, `ErrInvalidOwner`, `ErrInvalidDeposit`, `ErrInvalidAmount` |

**✅ O que funcionou bem:**
- `errors.As` para inspecionar `*ErrorInsufficientFunds` e acessar campos `Balance` e `Amount`
- Struct privada + constructor exposto — encapsulamento Go
- Methos consistentes: todos usam pointer receiver (ate os que nao precisariam, para consistencia)

**⚠️ Pontos a melhorar:**
- `NewBankAccount` retorna struct + erro simultaneamente (cria a conta mesmo com erro) — nao e idomatico (o caller so deve usar o valor se `err == nil`)
- `Deposit` valida `amount < 0`, `Withdraw` valida `amount <= 0` — inconsistente (0 e deposito valido mas saque invalido?)
- `ErrorInsufficientFunds.Error()` retorna mensagem generica sem incluir os campos `Balance`/`Amount`
- `ParseFloat` checa `n != 0` — nome enganoso (nao tem relacao com `strconv.ParseFloat`)
- `collectFlags` retorna `*flags` mas nao precisaria (so 1 chamador)

**🔁 Comportamento observado:**
- Primeiro contato com `*sync.WaitGroup` nil — tendencia de pensar que `var wg *sync.WaitGroup` cria instancia (como em TS/JS)

**📌 Sugestao para novos desafios:** Atencao a "zero value idiom" — em Go, valor zero deve ser util. E sempre retornar `nil` no valor quando retorna erro.

---

### 2.3 — Worker Pool (goroutines + channels + WaitGroup + context)

**Arquivos:** `worker_pool/main.go`, `worker_pool/pool/pool.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | N workers lendo do mesmo channel, `close()` sinalizando fim, `sync.WaitGroup`, `context.WithTimeout` |
| **Padrao chave** | `GenJobs` gera jobs e fecha `JobsChan` com `defer close()`; workers fazem `for task := range JobsChan` |

**✅ O que funcionou bem:**
- Buffer sizes corretas: `JobsChan` e `ResultChan` com capacidade = total de jobs (evita deadlock com worker pool)
- `select { case <-ctx.Done(): return; case JobsChan <- job: }` no `GenJobs` — pattern correto
- `for Id := range workersAmount` — uso de Go 1.22+ range sobre int

**⚠️ Pontos a melhorar:**
- `syncGroup.Wait()` + `close(ResultChan)` direto na `main` (nao em goroutine separada) — funciona neste caso porque workers sao as unicas goroutines no WG, mas e fragil se adicionar mais goroutines depois
- `GenWorker` tem `select` interno com `ctx.Done()` dentro de `for range` — se ctx cancela, o worker sai sem drenar o canal, mas os jobs restantes no buffer ficam orfaos (GC limpa)
- `time.Sleep(50ms)` no `GenJobs` so existe para "ver os jobs sendo processados" — em producao, nao se coloca sleep artificial em geradores
- Nomes excessivamente longos: `timeoutAmount`, `workersAmount`, `JobsAmount`
- Typo: `colletctFlags`

**🔁 Comportamento observado:**
- `defer close(chan)` dentro de cada goroutine individual — padrao incorreto que apareceu 2x (aqui e no parallel query)
- `len(chan)` para contar resultados processados — confusao entre elementos no buffer e elementos ja consumidos
- `var wg *sync.WaitGroup` nil — repeticao do erro do BankAccount

**📌 Sugestao para novos desafios:** Sempre fechar canal em goroutine exclusiva ou apos `Wait()` em goroutine separada. Revisar quando `close(chan)` deve estar na goroutine produtora vs consumidora.

---

### 2.4 — Parallel Query (fan-in simples + time.Duration)

**Arquivos:** `paralell_query/main.go`, `paralell_query/sources/source.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | N goroutines (fontes), 1 coletor, `time.Since()`, `rand`, `Result` struct com `ElapsedTime` |
| **Timeout** | `context.WithTimeout` configurado via flag `--timeout` |

**✅ O que funcionou bem:**
- `gracefulShutdown` padrao: goroutine separada faz `syncGroup.Wait()` + `close(resultChan)` — aprendido do worker pool
- `select { case <-parentCtx.Done(); case resultChan <- Result{...} }` — permite cancelamento mesmo se canal estiver cheio
- `Result.Source` usa `fmt.Sprintf("Source %d", callId)` — identificacao clara

**⚠️ Pontos a melhorar:**
- `time.Duration(time.Duration(randomSleep) * time.Millisecond)` — `time.Duration` aplicado duas vezes; `time.Duration(randomSleep) * time.Millisecond` e suficiente
- `startTime := time.Now()` antes do calculo de `randomSleep` — inclui no `ElapsedTime` o tempo gasto gerando numeros aleatorios (desprezivel, mas conceitualmente errado)
- `genContext` retorna `func() {}` como noop cancel — padrao aceitavel, mas a variavel descartada `_` seria mais limpa

**🔁 Comportamento observado:**
- Tentativa inicial de lancar goroutines e esperar resultado dentro do `for` (sequencial disfarcado) — instinto de "await" do TS
- Esquecer `flag.Parse()` novamente

**📌 Sugestao para novos desafios:** Lancar goroutines em lote PRIMEIRO, consumir resultados DEPOIS. Separacao mental clara entre "disparo" e "coleta".

---

### 2.5 — Fan-In com Merger (select multi-canal)

**Arquivos:** `fan-in/main.go`, `fan-in/producers/producerA.go`, `fan-in/producers/producerB.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `select` com 3 cases (chanA, chanB, ctx.Done), dois produtores com velocidades diferentes, merger, graceful shutdown |
| **Fluxo** | `ProducerA` → `chanA` → `merger` → `dataChan` → `main` / `ProducerB` → `chanB` → `merger` → `dataChan` → `main` |

**✅ O que funcionou bem:**
- Estrutura de 3 goroutines (A, B, merger) + 1 gracefulShutdown + main leitora — correta e sem deadlock
- Buffer nos 3 canais evita bloqueios (chanA=10, chanB=10, dataChan=100)
- `select` no merger com 3 cases: ler A, ler B, ou ctx.Done() — pattern correto

**⚠️ Pontos a melhorar:**
- Produtores usam `select { default: ... }` — o send dentro do `default` bloqueia sem escutar `ctx.Done()`. Se o buffer de `chanA`/`chanB` encher, o produtor trava no send e nao responde ao cancelamento
- Sleep ANTES do send (a cadencia inclui tempo de envio)
- Sem estatistica final (contagem de mensagens de A vs B)
- Nome do parametro `dataChan` nos produtores — na verdade recebem `chanA` e `chanB`

**🔁 Comportamento observado:**
- Bug inicial: produtores escreviam direto em `dataChan` (bypassando `chanA`/`chanB`), merger lia canais vazios — codigo morto
- `select` com `default` tratado como "sempre executa" sem entender que o conteudo do `default` bloqueia

**📌 Sugestao para novos desafios:** `select` com send dentro do case (nao no default) para responder a cancelamento. Sempre testar: "se eu der ctrl+C, quanto tempo esse codigo demora pra parar?"

---

### 2.6 — Rate Limiter Token Bucket (sync.Mutex + goroutine background) `[A.1]`

**Arquivos:** `rate-limiter/main.go`, `rate-limiter/bucket/bucket.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `sync.Mutex`, goroutine em background (`time.NewTicker`), channel de sinalizacao `stopChan`, Token Bucket algorithm |
| **Fluxo** | `NewBucket(N, interval)` → dispara goroutine que recarrega 1 token a cada `interval` ate `capacity`; `Allow()` consome token |

**✅ O que funcionou bem:**
- `Allow()` com `Lock/Unlock` correto e curto — sem segurar o lock por tempo desnecessario
- `increaseTokens` com `for { select { case <-timer.C, case <-stopChan } }` — sem double-consumption do timer
- `main.go` com loop de demonstracao (30 iteracoes, 100ms sleep, print pass/block)
- `Stop()` via canal buffered sinalizando a goroutine de recarga

**⚠️ Pontos a melhorar:**
- `stopChan` com buffer 1 e fragil — segundo `Stop()` travaria (bloqueio no send)
- `main.go` usa `NewBucket(25, 5s)` com loop de 3s — refill nunca dispara durante a demonstracao (so burst e visivel)
- Sem teste concorrente com `go test -race` (100 goroutines disputando tokens)

**🔁 Comportamento observado (iteracoes do bug):**
- **1ª tentativa:** `increaseTokens` fazia `b.mu.Lock()` fora do loop e segurava pra sempre → `Allow()` bloqueava em `Lock()` (deadlock)
- **2ª tentativa:** Moveu `Lock()` pra dentro do loop, mas `Unlock()` so dentro do `if b.tokens < b.capacity` → quando bucket enchia, `Unlock` nunca chamado, goroutine tentava `Lock()` de novo e deadlockava a si mesma (mutex não reentrante)
- **3ª tentativa:** `for range timer.C` + `select { case <-timer.C }` interno → double-consumption (refill a cada 10s em vez de 5s)
- **4ª tentativa (final):** `for { select { case <-timer.C: lock → inc → unlock; case <-stopChan: return } }` — correto

**📌 Sugestao para novos desafios:** Mutex deve ser segurado pelo menor tempo possivel — sempre Lock/Unlock no mesmo bloco. Cuidado com `select` aninhado dentro de `for range` sobre o mesmo canal.

---

### 2.7 — Linked List (ponteiros + nil seguro) `[C.1]`

**Arquivos:** `linked-list/main.go`, `linked-list/list/list.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `type Node struct { Val int; NextNode *Node }`, ponteiro de ponteiro implicito, iteracao com `for current != nil`, `Remove` com prev tracking |
| **Metodos** | `Insert`, `Remove`, `Reverse`, `Find`, `Print`, `Len` |

**✅ O que funcionou bem:**
- `Reverse()` com padrao de 3 ponteiros (prev, current, next) — estrutura correta
- `Remove()` trata cabeca separadamente e depois itera com prev tracking
- `Find()` retorna `nil` para nao encontrado (nil-safe)
- `Len()` iterativo sem campo extra

**⚠️ Pontos a melhorar:**
- `Insert()` demorou a corrigir: `if head == nil { head = node; return }` — o `return` foi esquecido, causando duplicata na lista vazia
- `Print()` imprime valores em linhas separadas (nao no formato `[1 -> 2 -> 3 -> nil]`)
- `main.go` cria cabeca manualmente com `Node{Val: 1}` em vez de usar `Insert` — funciona mas inconsistente

**🔁 Comportamento observado:**
- **Muita dificuldade com `Reverse()`** — manipulacao de ponteiros (prev, current, next) causou confusao. Estudante pediu para revisitar no futuro.

**📌 Sugestao para novos desafios:** Revisitar `Reverse` apos mais pratica com ponteiros. Desenhar no papel os 3 ponteiros antes de codificar ajuda.

---

### 2.8 — Shape Interface (polimorfismo + type switch) `[B.1]`

**Arquivos:** `interfaces/main.go`, `interfaces/shapes/shapes.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | Interface implicita (`Shape` com `Area()`/`Perimeter()`), type switch (`switch f := format.(type)`), duck typing, value vs pointer receiver |
| **Structs** | `Circle{R}`, `Rectangle{W, H}`, `Triangle{A, B, C}` com value receivers |

**✅ O que funcionou bem:**
- Interface definida e implementada automaticamente (sem `implements`, sem heranca)
- Type switch no `main.go` — estrutura correta
- `Triangle.Perimeter()` e `Triangle.Area()` com formula de Heron (quase correta)
- Slice heterogeneo: `[]Shape = append(forms, triangle, circle, rectangle)` — todos implementam a interface
- Respondeu corretamente que slice vazio nao tem elementos pra acessar e que interface exige todos os metodos

**⚠️ Pontos a melhorar:**
- **Formula de Heron usou perimetro total em vez de semi-perimetro** (`s := t.Perimeter() / 2` esquecido)
- Typo: `perimter` em vez de `perimeter`
- Nao implementou `TotalArea` (opcional)
- Nome da variavel de loop `format` atipico (mais comum: `shape` ou `s`)

**🔁 Comportamento observado:**
- Duvida: "se metodo e pointer receiver, preciso declarar `&Circle{}`?" — respondido: sim, se QUALQUER metodo for pointer receiver, so `*T` satisfaz a interface

**📌 Sugestao para novos desafios:** Proximo exercicio de interface (B.2) vai reforcar `io.Reader`/`io.Writer` — interfaces da stdlib.

---

### 2.9 — Stack & Queue (thread-safe + pub/priv split) `[C.2]`

**Arquivos:** `stack_queue/main.go`, `stack_queue/structures/stack.go`, `stack_queue/structures/queue.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `sync.Mutex`, padrao pub/priv (`isEmpty` vs `IsEmpty`), memory leak em slice, `Push`/`Pop`/`Peek`/`Len`/`IsEmpty` |
| **Structs** | `Stack{items []string, mu sync.Mutex}` com `Push`, `Pop`, `Peek`, `Len`, `IsEmpty`; `Queue` analogo com `Enqueue`/`Dequeue` |

**✅ O que funcionou bem:**
- Stack e Queue com todos os metodos thread-safe (lock em cada metodo publico)
- Padrao pub/priv split: `IsEmpty()` (publico, locked) → `isEmpty()` (privado, sem lock) — chamado internamente por `Pop`/`Dequeue`/`Peek` que ja seguram o lock. Evita lock recursivo
- `Pop` zera `s.items[lastIndex] = ""` antes de cortar — previne memory leak no backing array
- Sentinelas `EmptyStack`/`EmptyQueue` com `errors.New`
- `Dequeue` usa `q.items = q.items[1:]` — nao precisa zerar pois o slice passa a começar no índice 1 do backing array

**⚠️ Pontos a melhorar:**
- `main.go` so demonstra `Push`/`Peek`/`IsEmpty` — nao testa `Pop`/`Dequeue`/`Len`
- `Dequeue` perde os valores do backing array antigo sem zera-los — discutido e decidido que nao e necessario (slice começa de indice 1)

**🔁 Comportamento observado:**
- Duvida sobre `Lock()` em `IsEmpty()` e `Len()`: o estudante percebeu que adicionar lock nesses metodos criaria lock-recursivo quando chamados de `Pop()`/`Peek()` (que ja seguram o lock). Aprendizado: pub/priv split resolve.
- Inicialmente `items` era exportado (`Items` maiusculo) — corrigido para privado

**📌 Sugestao para novos desafios:** Pub/priv split e o padrao Go idiomatico para estruturas thread-safe com metodos auxiliares internos.

---

### 2.10 — countingWriter (wrapper io.Writer + contagem) `[B.2.5]`

**Arquivos:** `10-io-reader-writer/countingReader.go`, `10-io-reader-writer/upperWritter.go`, `10-io-reader-writer/main.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `io.Writer`, delegacao (wrapper/decorator), `io.Copy`, composicao de reader + writer |
| **Fluxo** | `io.Copy(upperWriter, countingReader)` → le do reader, escreve no writer, contagem acumulada |

**✅ O que funcionou bem:**
- `countingReader.Read` com delegacao simples: `c.reader.Read(p)` + `c.count += n` + retorno direto
- `upperWriter.Write` com `bytes.ToUpper` + delegacao pra `u.writer.Write`
- Entendimento do fluxo `io.Copy`: nao precisa de loop no `Read` — `io.Copy` faz o loop externo
- Entendimento de que `os.Stdout` e `strings.NewReader` ja implementam as interfaces — o wrapper so adiciona comportamento

**⚠️ Pontos a melhorar:**
- Dificuldade inicial com `countingReader.Read`: loop infinito dentro do `Read` foi a principal barreira mental
- Demorou a entender que `c.reader.Read(p)` chama o `Read` do `strings.Reader`, nao o proprio `Read` do wrapper

**🔁 Comportamento observado:**
- Confusao inicial: "por que chamo `Read` dentro do `Read`?" — resposta: o `Read` de fora e do `countingReader`, o `Read` de dentro e do `strings.Reader` (outro tipo)
- Depois do `countingWriter` (mais simples), o `countingReader` fez sentido

**📌 Sugestao para novos desafios:** Wrappers de `io.Reader`/`io.Writer` sao o padrao mais comum em Go stdlib. Praticar com `io.LimitReader` caseiro e `io.MultiWriter` caseiro consolida.

---

### 2.11 — BST (recursão + delete 3 casos) `[C.3]`

**Arquivos:** `11-bst/bst.go`, `11-bst/main.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `treeNode` struct, recursão em Go, delete 3 casos (folha, 1 filho, 2 filhos), `traverseLeft` como sucessor inorder |
| **Métodos** | `insert`, `delete`, `search`, `inOrder`, `min`, `max` |

**✅ O que funcionou bem:**
- `delete` com `findNodeToRemove` retornando `*treeNode` — padrão recursivo de religação da subárvore no desenrolar da pilha (mesmo pattern do `insertRecursive`)
- `search()` consertado (estava vazio, passou a delegar pra `searchAndAdvance`)
- Renomeação `advanceToInsert` → `insertRecursive` — nomenclatura mais clara

**⚠️ Pontos a melhorar:**
- `if node.left != nil && node.right != nil` no delete é redundante (caso 2 filhos é o único que sobra quando chega ali) — estudante deixou explícito de propósito
- `if/else` na navegação (`> val` / `< val`) vs `if/if` simétrico — frágil se alguém mexer na ordem dos blocos e pular o `return`

**🔁 Comportamento observado:**
- **Maior dificuldade:** entender como as chamadas recursivas se empilham na call stack, principalmente no delete — onde o retorno religa os ponteiros da subárvore modificada
- **Condições de parada:** `node == nil` (árvore vazia / não encontrou) e `node.value == val` (achou o alvo) — internalizado como padrão
- **`return node` no fim da recursão:** padrão que religa os ponteiros corretamente no desenrolar — foi o conceito mais difícil de fixar
- Caso 2 filhos exigiu visualizar: achar sucessor → copiar valor → deletar sucessor (deleção recursiva no mesmo nó)

**📌 Sugestão para novos desafios:** Desenhar a pilha de chamadas no papel antes de codificar recursão. O delete da BST é o melhor exercício pra fixar "return node" como reconstrução.

---

### 2.12 — Cache TTL (RWMutex + lazy eviction + cleanup goroutine) `[A.2]`

**Arquivos:** `12-cache-ttl/cache/cache.go`, `12-cache-ttl/main.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `sync.RWMutex` (RLock para leitura, Lock para escrita), lazy eviction no `Get`, goroutine de cleanup com `time.NewTicker`, sinalização com `close(chan struct{})` |
| **Métodos** | `NewCache`, `Set`, `Get`, `Delete`, `Stop` |

**✅ O que funcionou bem:**
- `RLock` em `Get`, `Lock` em `Set`/`Delete` — semântica correta de leitura concorrente vs escrita exclusiva
- Lazy eviction: TTL verificado no `Get` — sem depender da goroutine de cleanup pra funcionar
- `Stop()` com `ticker.Stop()` + `close(quit)` — ordem correta (parar ticker antes de fechar o canal evita tick perdido)
- `for { select { case <-ticker.C: lock → clean → unlock; case <-quit: return } }` — padrão idêntico ao Rate Limiter, aplicado corretamente
- `NewCache` evoluiu de "sem parâmetro" para `NewCache(interval time.Duration)` — decisão própria

**⚠️ Pontos a melhorar:**
- `defer c.mu.Lock()` em vez de `Unlock` — quase deadlock (corrigido)
- `NewCache` com `defer result.ticker.Stop()` — matava o ticker antes do primeiro cleanup (corrigido)
- `Stop()` inicialmente só fechava `quit` sem parar o ticker — goroutine fazia o `Stop` dentro do `case <-quit`

**🔁 Comportamento observado:**
- **Maior dificuldade:** manipulação de tempo em Go (`time.Time` vs `time.Duration`, `time.Now().Add(ttl)`, `time.Now().After(t)`)
- **`chan struct{}`:** estranhou o conceito de canal vazio pra sinalização — perguntou "por que não `chan bool`?" — internalizou que `struct{}` é o padrão Go por ocupar 0 bytes
- **`ticker.Stop()` não fecha o canal:** descobriu na prática que `for range ticker.C` trava mesmo após `Stop()` — levou à adição do `quit`
- **`close(nil)` → panic:** esqueceu de inicializar `quit` no construtor — aprendeu que canal zero value é nil
- Decidiu sozinho que `NewCache` sem parâmetro era suficiente (só lazy eviction), depois evoluiu para aceitar `interval` e adicionar cleanup periódico

**📌 Sugestão para novos desafios:** Sempre testar `close()` em canais não inicializados antes de confiar. `ticker.Stop()` não fecha o canal do ticker — se for usar `for range`, precisa de um canal de quit separado.

---

### 2.13 — Products API (HTTP CRUD + middleware) `[D.1]`

**Arquivos:** `13-products-api/main.go`, `13-products-api/product/products.go`, `13-products-api/product/store.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `net/http`, `http.ServeMux` com patterns de método + path param, `json.NewEncoder`/`json.NewDecoder`, `r.PathValue`, middleware pattern (`func(http.Handler) http.Handler`), `interface` no `main.go` para desacoplar store dos handlers |
| **Métodos/Rotas** | `GET /products`, `POST /products`, `GET /products/{id}`, `PUT /products/{id}`, `DELETE /products/{id}` |

**✅ O que funcionou bem:**
- Interface `database` no `main.go` — desacoplou os handlers da implementação concreta da store
- Middleware `jsonMiddleware` — centralizou `Content-Type` e removeu repetição dos handlers
- Store com `NewStore()` — manteve `items` privado com encapsulamento real
- ID auto-incremento com `sync.Mutex` — thread-safe e sem expor ponteiro interno
- Pub/priv split no `Store` — método privado `getById` + público `GetById` com lock resolveu deadlock

**⚠️ Pontos a melhorar:**
- `strconv.Atoi` erro usando `500 InternalServerError` em vez de `400 BadRequest` — erro é do cliente, não do servidor
- `userData.ID` vs `userData.Product.ID` em struct embutido — confusão com embedding de struct
- `ticker.Stop()` não fecha o canal — descoberto na prática (mesmo erro do Rate Limiter revisitado)

**🔁 Comportamento observado:**
- **Dificuldade principal:** struct fields minúsculos ignorados pelo `json.Encode` — apareceu 3x (no `dealError` original, no `getProducts` e nos handlers de POST/PUT/DELETE) até fixar que campo JSON precisa ser exportado
- **Middleware:** entendeu o padrão `func(next http.Handler) http.Handler` e aplicou corretamente — removeu `Content-Type` repetido de todos os handlers
- **JSON:** internalizou `json.NewDecoder(r.Body).Decode(&val)` para leitura e `json.NewEncoder(w).Encode(val)` para escrita
- **Interface no main:** entendeu que interface perto de quem usa (main.go) permite trocar implementação sem mexer nos handlers
- **Raciocínio sobre locks:** decidiu sozinho que `Delete` e `Update` deveriam ter `Lock` (não `RLock`) por serem operações de escrita — correto

**📌 Sugestão para novos desafios:** Separar handlers em arquivo próprio (`handlers.go`) e manter `main.go` só com servidor + middlewares. Próximo passo natural é D.2 (cadeia de middlewares).

---

### 2.14 — nil interface gotcha (B.3)

**Arquivos:** `14-nil-interface-gotcha/main.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | Interface como par `(type, value)`, `nil` interface vs `nil` pointer dentro de interface, `reflect.ValueOf(w).IsNil()`, `recover()` |
| **Fluxo** | `var w io.Writer` (nil) → `var buf *bytes.Buffer` (nil) → `w = buf` → `w != nil` (gotcha!) → `w.Write()` → panic |

**✅ O que funcionou bem:**
- Fluxo de demonstração correto: printa `w == nil: true`, atribui, printa `false`, chama `Write`, panica, `recover()` captura
- `safeWrite` com `reflect.ValueOf(w).IsNil()` — implementação funcional

**⚠️ Pontos a melhorar:**
- Resposta Q1 descreveu sintoma mas não o mecanismo (par type, value)
- Resposta Q3 não soube citar casos reais (error com ponteiro nil, io.Reader nil)
- `safeWrite` com `IsNil()` pode panica se kind não for nilable — precisa checar `Kind` antes
- Comentário: `bytes.Buffer` vs `*bytes.Buffer` — quem implementa `Write` é o pointer receiver

**🔁 Comportamento observado:**
- Entendimento prático do gotcha ok, mas conceito de `(type, value)` pair ainda não internalizado

**📌 Sugestão:** Revisar o conceito de interface value = (type, value) pair antes de avançar para B.4 (Generic Stack) ou D.2 (Middleware Chain).

---

### 2.15 — Middleware Chain (D.2) ✅

**Arquivos:** `15-middleware-chain/main.go`, `15-middleware-chain/product/products.go`, `15-middleware-chain/product/store.go`

| Item | Detalhe |
|---|---|
| **Conceitos** | `func(http.Handler) http.Handler`, chain de middlewares, `defer recover()`, `bearer token` auth, CORS (`OPTIONS` preflight) |
| **Middlewares** | `Logger`, `Recoverer`, `Auth`, `CORS`, `JSON` (Content-Type) |
| **Chain** | `CORS → Auth → Recoverer → Logger → JSON → mux` |

**✅ O que funcionou bem:**
- `chain()` helper com variadic `...middleware` — pattern reutilizável
- `Logger` com tempo antes e depois do `next` — duração correta
- `Recoverer` capturando panic e retornando 500 JSON — sem crash
- `CORS` tratando `OPTIONS` sem chamar `next` — preflight correto
- `jsonMiddleware` centralizando `Content-Type` — sem repetição nos handlers
- Store + Product copiados do D.1 sem erros novos — padrão consolidado
- `authMiddleware` com bypass para `GET /products` (rota pública)

**⚠️ Erros encontrados (ver correcoes.md para detalhes):**
- `chain()` off-by-one — `len(middlewares)` em vez de `len-1`
- `chain()` ignorado no `ListenAndServe` — middlewares duplicados
- `recoverer` escrevia 500 em toda request (fora do `if recover()`)
- Auth: `Split(token, "")` em vez de `Split(token, " ")` e `bearerTkn[2]` em vez de `[1]`
- Auth: token ausente não retornava 401
- `strconv.Atoi` erro como 500 (repetido do D.1)
- Campo JSON não exportado no recoverer (repetido do D.1)
- CORS sem `Access-Control-Allow-Headers`

**🔁 Comportamento observado:**
- Erros do D.1 se repetiram (500 em parsing, campo JSON privado) — indica que a correção anterior não fixou o hábito
- Off-by-one no loop do `chain()` — padrão que já havia aparecido em desafios anteriores
- Criou `chain()` mas não usou — codificou a solução e ignorou no passo seguinte (falta de revisão antes de rodar)
- Respostas conceituais Q1, Q2, Q3 incompletas ou tangenciais — dificuldade em responder exatamente o que foi perguntado

**📌 Sugestão:** Revisar o código inteiro antes de dar como pronto — os 3 bugs mais graves (chain off-by-one, recoverer incondicional, auth quebrado) seriam pegos por uma execução sequer com `go run`. Rodar o servidor e testar com `curl` antes de considerar pronto.

---

### 2.16 — Nil Interface Revisão (B.3) ❌ ABANDONADO

**Arquivos:** `16-nil-interface-revisao/README.md`, `16-nil-interface-revisao/perguntas.md`

| Item | Detalhe |
|------|---------|
| **Conceitos** | Interface par (type, value), falso positivo err != nil, reflect.ValueOf.IsNil, type assertion |
| **Motivo** | Conceito ainda não internalizado. Retornar quando houver disposição para revisitar com calma. |
| **Próximo passo sugerido** | Praticar outros desafios de Go e voltar a este quando o `par (type, value)` soar mais natural |

### Ranking por frequencia

| # | Erro | Exercicios onde apareceu | Por que acontece |
|---|---|---|---|
| 1 | **`var wg *sync.WaitGroup` (ponteiro nulo)** | 1.2, 1.3, 1.4 | Vicio de linguagens onde `var x *T` cria instancia; em Go cria nil |
| 2 | **`defer close(chan)` dentro de cada goroutine** | 1.3, 1.4 | Pensar que cada goroutine e dona do canal e pode fecha-lo |
| 3 | **`select { default: ... send ... }` sem escutar cancelamento** | 1.5 | Confundir `default` executavel com `default` cancelavel |
| 4 | **Sequencial disfarcado de paralelo** (lancar goroutine + esperar resultado no mesmo loop) | 1.4 | Intuicao de `await` do TypeScript |
| 5 | **Esquecer `flag.Parse()`** | 1.1, 1.4 | Registrar flags com `Var` mas nao chamar `Parse()` |
| 6 | **`time.Duration(int)` = nanossegundos** | varios | Esquecer de multiplicar por `time.Millisecond`/`time.Second` |
| 7 | **`len(channel)` para contar processados** | 1.3 | Confundir buffer occupancy com total consumido |
| 8 | **`Lock()` sem `Unlock()` em todos os caminhos** | A.1 | `if b.tokens < b.capacity { ... b.mu.Unlock() }` — `Unlock` so chamado dentro do if, nunca fora |
| 9 | **Mutex segurado por toda goroutine** (Lock antes do loop) | A.1 | Bloqueio permanente que impede qualquer outra goroutine de adquirir o lock |
| 10 | **`for range timer.C` + `select { case <-timer.C }`** (double consumption) | A.1 | `for range` ja consome o tick; o `select` interno le o timer.C vazio e espera o proximo |
| 11 | **`stopChan` com buffer 1 — segundo `Stop()` trava** | A.1 | Buffer cheio, ninguem le, send bloqueia pra sempre |
| 12 | **Off-by-one em loop (`len` vs `len-1`)** | D.2 | Usar `len(middlewares)` em vez de `len-1` como último índice válido |
| 13 | **Código condicional fora do `if`** (executa sempre) | D.2 | `WriteHeader(500)` fora do `if recover() != nil` — sempre executava |
| 14 | **`Split` com separador errado** | D.2 | `Split(token, "")` em vez de `Split(token, " ")` — espaço vs vazio |
| 15 | **Status code errado p/ erro do cliente (500 → 400)** | D.1, D.2 | `strconv.Atoi` parsing invalido nao e erro do servidor — repetiu o erro do D.1 |
| 16 | **Campo JSON nao exportado (letra minuscula)** | D.1, D.2 | `message string` — `json.Encode` ignora campos privados — repetiu o erro do D.1 |
| 17 | **Codificar solucao e nao usa-la** (chain ignorado) | D.2 | Criou `chain()` mas passou wrapping manual pro `ListenAndServe` |
| 18 | **errors.Is vs errors.As invertido** (confundiu qual faz o quê) | Reparo 3.2 | Disse que `Is` é "confirmar se erro é de um tipo" (é `As`) |
| 19 | **Slice memory leak** (não sabia por que Pop zera mas Dequeue não) | Reparo 3.3 | Não entendeu que backing array retém referências mesmo após `s = s[1:]` |
| 20 | **Nil interface par (type, value) impreciso** | Reparo 3.4 | Sabe o fenômeno mas linguagem confusa sobre type vs value |
| 21 | **`s[cap(s)]` em vez de `s[:cap(s)]`** | 18 | Confundir acesso por índice com re-slice para expandir janela do backing array |
| 22 | **`dequeue` implementado como `popSafe`** | 18 | Removeu o último elemento em vez do primeiro — copiou código sem adaptar |
| 23 | **`popSafe` retornava valor zerado** | 18 | Salvou o valor DEPOIS de zerar, em vez de antes — ordem errada das operações |
| 24 | **`errors.AsType[V]` em vez de `[*V]`** | 17 | Type assertion `err.(ValidationError)` falha se err contém `*ValidationError` |

### Padroes conceituais ja internalizados

| Conceito | Evidencia |
|---|---|
| `errors.Is` vs `errors.As` (sentinel vs struct) | Usa corretamente em todos os exercicios |
| Pointer receiver para mutacao | `Deposit`, `Withdraw`, `Convert` |
| Buffer de channel = protecao contra deadlock | `make(chan T, N)` com N >= total |
| `gracefulShutdown` com goroutine separada | Padrao repetido em 1.3, 1.4, 1.5 |
| `select` com `ctx.Done()` como case | Todos os exercicios com timeout |
| Struct private + constructor publico | `bankAccount` + `NewBankAccount` |

---

## 3. Exercícios de Reparo Planejados

Exercícios criados após avaliação conceitual para lacunas identificadas nos tópicos abaixo.

### 3.1 — select-sem-default (Reparo: Q2/Q9) ✅

**Arquivos:** `16-select-sem-default/`

| Item | Detalhe |
|------|---------|
| **Conceitos** | `select` sem `default`, send como `case`, cancelamento com `ctx.Done()`, fan-in cancelável |
| **Lacuna** | Erro #3 do ranking (select com default contendo send bloqueante) — curado. |
| **Objetivo** | Escrever `sendOrCancel[T]` e `fanInCancelable[T]` usando select **sem default**, com send como case. Comparar com versão non-cancelável. |
| **Status** | ✅ Concluído. Respostas conceituais corretas. |

### 3.2 — error-is-as (Reparo: Q4) ✅

**Arquivos:** `17-error-is-as/`

| Item | Detalhe |
|------|---------|
| **Conceitos** | `errors.Is` (valor sentinela), `errors.As` (tipo), `%w` wrapping |
| **Lacuna** | Aluno inverteu os papéis — disse que `Is` é "confirmar se erro é de um tipo" (isso é `As`) — curado. |
| **Objetivo** | Implementar `processItem` com erros sentinela + tipado + wrapped, e `handleError` com `Is` e `As` lado a lado. |
| **Status** | ✅ Concluído. Usou `errors.AsType[*ValidationError]` corretamente (Go 1.26). |

### 3.3 — slice-leak (Reparo: Q6) ✅

**Arquivos:** `18-slice-leak/`

| Item | Detalhe |
|------|---------|
| **Conceitos** | Slice backing array, capacidade, memory leak em Pop vs Dequeue |
| **Lacuna** | Aluno não sabia responder; buscou ajuda externa. Conceito de backing array + data pointer ainda não internalizado — curado. |
| **Objetivo** | Demonstrar com `s[:cap(s)]` que elementos "removidos" ainda estão no backing array. Implementar `pop`, `popSafe`, `dequeue`. |
| **Status** | ✅ Concluído. Código funcional + respostas corretas. |

### 3.4 — nil-interface-revisao (Reativado: Q3)

**Arquivos:** `19-nil-interface-revisao/`

| Item | Detalhe |
|------|---------|
| **Conceitos** | Interface (type, value) pair, nil pointer vs nil interface, `reflect.ValueOf.IsNil` |
| **Lacuna** | Exercício 16 abandonado. Linguagem imprecisa sobre o par (type, value). |
| **Objetivo** | Re-escrever o programa do gotcha e explicar por escrito os 3 momentos do par (type, value). |

### Ordem de execução sugerida

```
1. 16-select-sem-default  (fixa o padrão mais crítico)
2. 17-error-is-as         (rápido, consolida)
3. 18-slice-leak          (visual, experimental)
4. 19-nil-interface-revisao (fechamento conceitual)
```

Após concluir os 4 reparos, voltar ao roteiro progressivo normal: A.3 → B.4 → C.4 → D.3/D.4.

---

## 4. Roteiro Progressivo

### Modulo A — Concorrencia Avancada (🔥🔥🔥🔥🔥) `[3-5 dias]`

Peso maximo em entrevistas. Foco em `select` com send, `sync.Mutex`/`RWMutex`, `sync.Once`, goroutine leak.

#### A.1 — Rate Limiter Token Bucket ✅ CONCLUIDO
- `type Bucket struct { tokens int; capacity int; mu sync.Mutex; interval time.Duration }`
- `Allow() bool` — consome token se disponivel
- Goroutine em background recarrega tokens periodicamente
- **Teste:** 100 goroutines tentando consumir — garantir que nao ultrapassa `capacity`

#### A.2 — Cache Concorrente com TTL ✅ CONCLUIDO
- `type Cache struct { data map[string]cacheEntry; mu sync.RWMutex }`
- `type cacheEntry struct { value any; expiresAt time.Time }`
- `Get`, `Set`, `Delete` com `RWMutex` (leituras concorrentes, escrita exclusiva)
- **Pegadinha:** limpar entradas expiradas no `Get` ou com goroutine de limpeza periodica

#### A.3 — Scheduler com Fan-Out/Fan-In
- N workers processam tarefas de uma fila
- Cada worker envia resultado para um canal unico
- Merger coleta todos os resultados
- **Novo conceito:** `sync.Once` para fechar o canal de resultados exatamente uma vez

---

### Modulo B — Interfaces & Type System (🔥🔥🔥🔥) `[3-4 dias]`

O que mais diferencia Go de outras linguagens. Interfaces sao implicitas (duck typing).

#### B.1 — Polimorfismo com `Shape` interface ✅ CONCLUIDO
- `type Shape interface { Area() float64; Perimeter() float64 }`
- Implementar `Circle`, `Rectangle`, `Triangle`
- Funcao `PrintShapeInfo(s Shape)` — polimorfismo
- Funcao `TotalArea(shapes []Shape) float64`
- **Novo conceito:** type assertion `c, ok := s.(*Circle)` e type switch `switch v := s.(type)`

#### B.2 — Reader/Writer do zero ✅ CONCLUIDO
- `type UpperWriter struct { writer io.Writer }` — implementa `io.Writer` transformando em maiusculas
- `type CountingReader struct { reader io.Reader; count int64 }` — implementa `io.Reader` contando bytes
- `countingWriter` (extra) — delegacao + contagem para `io.Writer`
- **Novo conceito:** middleware pattern com interfaces (wrapper/decorator)

#### B.3 — nil interface gotcha ✅ CONCLUIDO
- `var w io.Writer` (nil interface — valor E tipo sao nil)
- `var buf *bytes.Buffer` (nil pointer) → atribuir a `w = buf` → `w != nil` (o tipo nao e nil!)
- Escrever codigo que demonstra o gotcha e explica POR QUE acontece
- **Novo conceito:** interface value = (type, value) pair

#### B.4 — Generic Stack (Go 1.18+)
- `type Stack[T any] struct { items []T }`
- `Push`, `Pop`, `Peek`, `IsEmpty`
- Comparar com interface vazia `interface{}` / `any` — quando usar generics vs `any`

---

### Modulo C — Estruturas de Dados (🔥🔥🔥) `[4-5 dias]`

Estruturas classicas em Go. Atencao a ponteiros, nil seguro, e slices.

#### C.1 — Linked List ✅ CONCLUIDO
- `type Node struct { Value int; Next *Node }`
- Metodos: `Insert`, `Remove`, `Reverse`, `Find`, `Len`
- **Novo conceito:** `**Node` para modificar a raiz sem retornar (c idiom)

#### C.2 — Stack & Queue (thread-safe) ✅ CONCLUIDO
- `type Stack struct { items []string }` → `Push`, `Pop (string, error)`
- `type Queue struct { items []string }` → `Enqueue`, `Dequeue`
- Versao thread-safe com `sync.Mutex`
- **Pegadinha:** memory leak em Pop — o slice subjacente mantem referencia ao elemento removido. Fazer `s.items[len-1] = ""` antes de cortar.

#### C.3 — Binary Search Tree ✅ CONCLUIDO
- `type Node struct { Value int; Left, Right *Node }`
- `Insert`, `Search`, `InOrder`, `PreOrder`, `PostOrder` (retornam `[]int`)
- `Min`, `Max`, `Delete` (o Delete e o mais dificil — 3 casos)
- **Novo conceito:** recursao em Go (nao tem tail-call optimization, pilha pode estourar em arvore profunda)

#### C.4 — LRU Cache
- `type LRUCache struct { capacity int; cache map[int]*listNode; head, tail *listNode }`
- Usar lista duplamente encadeada + hash map
- `Get(key)`, `Put(key, value)`
- **Novo conceito:** combinar duas estruturas, mover node para frente (O(1))

---

### Modulo D — HTTP & APIs (🔥🔥🔥) `[4-5 dias]`

Foco em `net/http`, `encoding/json`, middlewares, graceful shutdown.

#### D.1 — CRUD de Produtos ✅ CONCLUIDO
- `type Product struct { ID int; Name string; Price float64 }`
- Rotas: `GET /products`, `POST /products`, `GET /products/{id}`, `PUT /products/{id}`, `DELETE /products/{id}`
- Usar `http.ServeMux` (Go 1.22+) com patterns como `GET /products/{id}`
- Store em memoria com `sync.RWMutex`
- `Content-Type: application/json`

#### D.2 — Middleware Chain ✅ CONCLUIDO
- `func Logger(next http.Handler) http.Handler` — loga metodo, path, duracao
- `func Recoverer(next http.Handler) http.Handler` — `defer recover()` evita crash
- `func Auth(next http.Handler) http.Handler` — header `Authorization: Bearer <token>`
- `func CORS(next http.Handler) http.Handler` — CORS headers + OPTIONS preflight
- `func chain(h http.Handler, m ...middleware) http.Handler` — composição variádica
- Chain final: `CORS(Auth(Recoverer(Logger(JSON(mux)))))`

#### D.3 — JSON Errors padronizados
- `type APIError struct { Code int; Message string; Details any }`
- `func WriteJSON(w, status, data)` e `func WriteError(w, status, msg)`
- Validar body (campos obrigatorios, tipos), retornar 400 com erro estruturado

#### D.4 — Graceful Shutdown
- `signal.NotifyContext` para SIGINT/SIGTERM
- `srv.Shutdown(ctx)` com timeout
- Drenar conexoes ativas antes de encerrar

---

### Modulo E — Testes (🔥🔥🔥) `[3-4 dias]`

Table-driven tests, mocks via interface, testes concorrentes com `-race`.

#### E.1 — Table-Driven Tests
- Para todos os handlers do CRUD (D.1)
- `tests := []struct { name, method, path, body string; expectedStatus int; expectedBody string }`
- `for _, tt := range tests { t.Run(tt.name, func(t *testing.T) { ... }) }`

#### E.2 — Mock via Interface
- Separar storage em interface: `type Storage interface { ... }`
- `InMemoryStorage` para producao, `MockStorage` com funcoes injetaveis para testes
- Testar handlers sem storage real

#### E.3 — Testes Concorrentes
- Disparar 10 goroutines fazendo GET/POST no mesmo server de teste
- Rodar com `go test -race`
- Identificar e corrigir data races

#### E.4 — Table-Driven para estruturas de dados
- Testar Linked List, Stack, Queue, BST com tabelas de casos
- Casos de borda: vazio, 1 elemento, muitos elementos, nulos

---

### Modulo F — Projetos Integradores (🔥🔥🔥) `[5-7 dias]`

Projetos que juntam tudo: HTTP + concorrencia + estruturas de dados + testes.

#### F.1 — URL Shortener
- `POST /shorten` (body: `{"url": "..."}`) → `{"short": "abc123"}`
- `GET /{short}` → redirect 301
- Gerar short code com `crypto/rand`
- Store em memoria + `sync.RWMutex`
- Testes de integracao

#### F.2 — Todo List API com persistencia JSON
- CRUD de tasks (`pending`/`done`)
- Persistir em arquivo `tasks.json` a cada mutacao
- Carregar do arquivo na inicializacao
- Endpoints RESTful + testes

#### F.3 — Chat via SSE (Server-Sent Events)
- `GET /events` — streaming de eventos
- `POST /message` — broadcast para todos os clients
- Channel de broadcast + `http.Flusher`
- Desafio: gerenciar clients que desconectam (evitar leak de goroutine e channel)

#### F.4 — API Gateway / Proxy Reverso simples
- Roteia requests para backends configurados
- Round-robin load balancing
- Health check nos backends (goroutine periodica)
- Rate limiting por IP

---

## 5. Ordem Sugerida de Execucao

```
Semana R:  16-select-sem-default ✅ + 17-error-is-as ✅ + 18-slice-leak ✅ + 19-nil-interface-revisao ⬜  ← REPAROS (3/4 concluídos)
Semana 1:  A.1 (Rate Limiter) + C.1 (Linked List)
Semana 2:  B.1 (Shape interface) + C.2 (Stack/Queue)
Semana 3:  B.2 (Reader/Writer) + C.3 (BST) ✅
Semana 4:  A.2 (Cache TTL) + B.3 (nil interface gotcha)
Semana 5:  D.1 (CRUD API) + D.2 (Middlewares) ✅
Semana 6:  D.3 (JSON errors) + D.4 (Graceful Shutdown)
Semana 7:  E.1 (Table tests) + E.2 (Mock via interface)
Semana 8:  F.1 (URL Shortener)
Semana 9:  A.3 (Fan-Out/Fan-In) + C.4 (LRU Cache)
Semana 10: F.2 (Todo API) ou F.3 (Chat SSE)
```

> Intercale modulos para nao enjoar: sempre junte 1 exercicio de estruturas + 1 de concorrencia + 1 de HTTP por semana.
>
> **Semana R:** 3/4 concluídos. Finalizar `19-nil-interface-revisao` antes de prosseguir.

---

## 6. Flashcards / Go-tchas

Lista de pegadinhas classicas para revisar antes de entrevistas:

| # | Gotcha | Explicacao relampago |
|---|---|---|
| 1 | `nil map` panic | `var m map[string]int; m["k"] = 1` → panic. Use `make()` ou `map[string]int{}` |
| 2 | `range` copia o valor | `for _, v := range items { go func() { fmt.Println(v) }() }` → todas goroutines veem a copia do ULTIMO valor. Passe por parametro ou crie variavel local no corpo do loop |
| 3 | `nil interface != nil pointer` | `var w io.Writer; var b *bytes.Buffer; w = b; w != nil` → true! O tipo nao e nil |
| 4 | `defer` avalia argumentos no momento do `defer` | `defer fmt.Println(i)` com `i=0` depois `i++` ainda imprime 0. Use closure: `defer func() { fmt.Println(i) }()` |
| 5 | `WaitGroup` zero value e util | `var wg sync.WaitGroup` funciona. `var wg *sync.WaitGroup` e nil e PANICA |
| 6 | `close(nil chan)` → panic | So fechar canais inicializados com `make` |
| 7 | `close(chan)` por multiplas goroutines → panic | Exatamente 1 goroutine deve fechar o canal |
| 8 | `sync.Mutex` nao pode ser copiado | `func (s Store) Get() { s.mu.Lock() }` → value receiver copia o mutex. SEMPRE use pointer receiver com Mutex |
| 9 | Goroutine leak | Toda goroutine iniciada deve ter um caminho para terminar (`ctx.Done()`, channel close, etc.) |
| 10 | `for range` sobre channel bloqueia ate o channel fechar | Sem `close(ch)`, `for v := range ch` bloqueia para sempre |
| 11 | Append pode ou nao realocar | `a := b[:2]; a = append(a, 3)` pode sobrescrever `b[2]` se houver capacidade. Ou pode criar novo array |
| 12 | `json.Unmarshal` em struct com campos privados | Campos com letra minuscula sao privados e ignorados pelo `encoding/json`. Use tags: `json:"name"` com letra maiuscula |

---

## 7. Checklist de Revisao

Apos cada modulo, preencher:

```
Data: __/__/____
Modulo concluido: ___
Erros antigos que NAO se repetiram:
  - [ ] 
  - [ ] 
Erros novos:
  - [ ] 
  - [ ] 
Conceito novo que ficou claro:
  - 
Conceito que ainda gera duvida:
  - 
Proximo modulo:
  - 
```

---

## 8. Arquivos Mapeados

```
learnings/
├── README.md                  ← índice geral do repositório
│
├── go/
│   ├── README.md              ← atalhos para o roteiro
│   ├── roteiro.md             ← Este arquivo (plano de estudos Go)
│   │
│   ├── 01-currency-conversor/ ← Exercicio 2.1 (CLI + sentinel errors)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── utils/typos.go
│   ├── 02-bank-account/       ← Exercicio 2.2 (pointer receiver + erro tipado)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── utils/definitions.go
│   ├── 03-worker-pool/         ← Exercicio 2.3 (goroutines + WaitGroup)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── pool/pool.go
│   ├── 04-paralell-query/      ← Exercicio 2.4 (fan-in simples)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── sources/source.go
│   ├── 05-fan-in/              ← Exercicio 2.5 (select multi-canal)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── producers/{producerA,producerB}.go
│   ├── 06-rate-limiter/        ← Exercicio 2.6 (token bucket + sync.Mutex)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── bucket/bucket.go
│   ├── 07-linked-list/         ← Exercicio 2.7 (ponteiros + nil seguro)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── list/list.go
│   ├── 08-shape-interface/     ← Exercicio 2.8 (shape interface + type switch)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── shapes/shapes.go
│   ├── 09-stack-queue/         ← Exercicio 2.9 (stack + queue thread-safe)
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── structures/{stack,queue}.go
│   ├── 10-io-reader-writer/    ← Exercicio 2.10 (countingReader + upperWriter)
│   │   ├── main.go
│   │   ├── go.mod
│   │   ├── countingReader.go
│   │   └── upperWritter.go
│   ├── 11-bst/                 ← Exercicio C.3 (BST) ✅ CONCLUIDO
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── bst.go
│   ├── 12-cache-ttl/           ← Exercicio A.2 (Cache TTL) ✅ CONCLUIDO
│       ├── main.go
│       ├── go.mod
│       └── cache/cache.go
│   └── 13-products-api/        ← Exercicio D.1 (CRUD HTTP API) ✅ CONCLUIDO
│       ├── main.go
│       ├── go.mod
│       ├── product/products.go
│       └── product/store.go
│   └── 14-nil-interface-gotcha/ ← Exercicio B.3 (nil interface gotcha) ✅ CONCLUIDO
│       ├── README.md
│       ├── perguntas.md
│       ├── respostas.md
│       └── main.go
│   └── 15-middleware-chain/      ← Exercicio D.2 (Middleware Chain) ✅ CONCLUIDO
│       ├── README.md
│       ├── perguntas.md
│       ├── respostas.md
│       ├── go.mod
│       ├── main.go
│       └── product/{store,products}.go
│   └── 16-nil-interface-revisao/  ← Exercicio B.3 (nil interface gotcha) ❌ ABANDONADO
│       ├── README.md
│       └── perguntas.md
│   └── 16-select-sem-default/     ← Reparo 3.1 (select sem default)
│       ├── README.md
│       └── perguntas.md
│   └── 17-error-is-as/            ← Reparo 3.2 (errors.Is vs errors.As)
│       ├── README.md
│       └── perguntas.md
│   └── 18-slice-leak/             ← Reparo 3.3 (slice backing array)
│       ├── README.md
│       └── perguntas.md
│   └── 19-nil-interface-revisao/  ← Reparo 3.4 (reativado)
│       ├── README.md
│       └── perguntas.md
│
├── typescript/                 ← Futuros exercícios TS
└── dsa/                        ← Revisoes pendentes de DSA
    └── README.md
```

> **Para retomar a sessao:** "Professor, estou no Modulo ___ do roteiro.md. Vamos comecar o exercicio ___."
