# Módulo C — Estruturas de Dados & I/O

> Scaffolding: médio

**Objetivo:** Construir estruturas thread-safe, dominar io.Reader/Writer, entender armadilhas de slices.

## Desafios

| # | Desafio | O que você construiu |
|---|---------|---------------------|
| C1 | stack-queue | Stack e Queue thread-safe com split público/privado |
| C2 | io-reader-writer | CountingReader e UpperWriter com delegação e io.Copy |
| C3 | bst | Árvore binária de busca: Insert, Search, Delete (3 casos), Traverse |
| C4 | error-is-as | Correção de uso incorreto de errors.Is e errors.As com %w wrapping |
| C5 | slice-leak | Correção de memory leak em Pop vs Dequeue, backing array |

## Conceitos dominados ao final do módulo

- Estruturas thread-safe — mutex protegendo estado interno
- Separação público/privado — struct exportada, campos não exportados
- `io.Reader` / `io.Writer` — implementar e delegar
- `io.Copy` — streaming eficiente entre reader e writer
- Árvore binária — recursão, busca, inserção, remoção (3 casos)
- BST Delete — folha, 1 filho, 2 filhos (sucessor inorder)
- `errors.Is` — unwrapping automático na cadeia %w
- `errors.As` — extração de tipo na cadeia de erros
- Slice backing array — compartilhamento entre slices derivados
- Memory leak — referência mantida ao array subjacente

## Para revisitar

Consulte o `spaced-repetition/schedule.md` para as datas de revisão agendadas.
