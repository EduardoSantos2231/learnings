# C2 — Reader e Writer

> Implementacao | 45 min | Go stdlib

## Objetivo

Crie wrappers de `io.Reader` e `io.Writer` que contem bytes e transformam texto.

## Faca

1. Implemente um reader que conta bytes lidos.
2. Implemente um writer que transforma letras em maiusculas.
3. Componha ambos com `io.Copy`.
4. Preserve e retorne erros do componente interno.

## Pronto quando

- Os wrappers funcionam com qualquer reader ou writer.
- O contador corresponde aos bytes lidos.
- Testes cobrem leitura parcial e erro interno.

## Responda

- Por que delegar em vez de duplicar o loop?
- O que significa `io.EOF`?

> Confianca: [1-5]
