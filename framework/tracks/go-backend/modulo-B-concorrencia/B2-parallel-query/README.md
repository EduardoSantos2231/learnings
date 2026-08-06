# B2 — Parallel Query

> Debug | 45 min | Go stdlib

## Objetivo

Corrija o agregador que consulta fontes em paralelo e retorna antes do timeout.

## Faca

1. Reproduza a falha com o teste fornecido.
2. Encontre goroutines, canais ou timers que nao encerram.
3. Preserve resultados validos e erros individuais.
4. Corrija sem reescrever o programa.

## Pronto quando

- O teste de timeout termina sempre.
- Nenhuma goroutine fica bloqueada.
- O resultado identifica a fonte que falhou.

## Responda

- Qual operacao bloqueava?
- Como provaria que o vazamento desapareceu?

> Confianca: [1-5]
