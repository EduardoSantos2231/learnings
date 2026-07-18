# Capstone 3 — Cache com Persistência

> Template: Capstone | Síntese do Módulo C | Sem scaffolding

## Contexto

Você vai construir uma biblioteca de cache em memória com TTL, thread-safe,
que persiste os dados em disco e os recupera no próximo start.
Combina estruturas de dados, I/O e armadilhas da linguagem.

## Requisitos

### Funcionalidades obrigatórias

1. **Cache thread-safe** (C1): `Get`, `Set`, `Delete` com `sync.RWMutex`.
   Separação público/privado: struct exportada, campos internos não exportados.
2. **TTL com lazy eviction** (D1): itens expiram no acesso. Cleanup goroutine
   opcional que varre expirados periodicamente.
3. **Persistência em disco** (C2): implemente `io.WriterTo` e `io.ReaderFrom`
   na struct do cache. Use `io.Copy` internamente para streaming eficiente.
4. **Índice por ordem de inserção** (C3): implemente uma BST que mantém as
   chaves ordenadas por timestamp de inserção, permitindo iterar na ordem.
5. **Error wrapping** (C4): todos os erros de I/O devem ser wrappeados com
   contexto (`%w`) e usar `errors.Is`/`errors.As` corretamente.
6. **Zero leak** (C5): operações de remoção não devem manter referências
   ao backing array. Implemente `Delete` e `Clear` sem vazamento.

### Requisitos não-funcionais

- Salvar/carregar usa formato binário próprio (gob, json ou binário customizado)
- Thread-safe em todas as operações

## Tarefas

### Fase 1: Design

Desenhe a struct `Cache[K comparable, V any]`:
- Campos internos (mutex, mapa, BST, cleanup channel)
- Métodos públicos (`Get`, `Set`, `Delete`, `Clear`, `Save`, `Load`)
- Como a BST e o mapa se mantêm sincronizados?

### Fase 2: Implementação

### Fase 3: Testes

- Concorrência: 100 goroutines fazendo Get/Set simultâneos
- TTL: item expira e é removido no Get
- Persistência: salva → novo cache → carrega → dados intactos
- Memory leak: Delete seguido de GC → sem referências residuais

### Fase 4: Retrospectiva

## Conceitos envolvidos

- Estrutura thread-safe, pub/priv — C1 (stack-queue)
- io.Reader/Writer, io.Copy — C2 (io-reader-writer)
- BST, recursão — C3 (bst)
- errors.Is, errors.As, %w — C4 (error-is-as)
- Slice backing array, memory leak — C5 (slice-leak)
