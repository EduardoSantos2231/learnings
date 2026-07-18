# Capstone 1 — Calculadora de Expressões CLI

> Template: Capstone | Síntese do Módulo A | Sem scaffolding

## Contexto

Você vai construir uma calculadora que avalia expressões matemáticas via linha de comando.
Ela recebe uma string como `"3 + 4 * 2 / ( 1 - 5 )"` e retorna o resultado.

Este projeto exercita todos os conceitos do Módulo A em um contexto integrado.

## Requisitos

### Funcionalidades obrigatórias

1. **Parser de expressões**: converter string infixa para notação polonesa reversa (RPN)
   usando o algoritmo Shunting-yard. Use uma stack (linked list do A3).
2. **Evaluator**: avaliar a expressão RPN e retornar o resultado float64.
3. **Tratamento de erros**: use erros sentinela para situações previsíveis
   (`ErrDivisionByZero`, `ErrSyntax`, `ErrUnmatchedParenthesis`) e `errors.Is`
   para verificá-los.
4. **CLI**: leia a expressão de `os.Args[1]` ou stdin. Exiba o resultado ou o erro.
5. **Operadores**: suporte `+`, `-`, `*`, `/` e parênteses. Use uma interface
   `Operator` com método `Apply(a, b float64) (float64, error)` e type switch
   para selecionar o operador correto.

### Requisitos não-funcionais

- Precisão float64 (sem big decimal)
- Mensagens de erro claras, incluindo posição do erro na expressão

## Tarefas

### Fase 1: Design (sem código)

Desenhe a arquitetura:
- Quais pacotes? (`parser`, `evaluator`, `stack`, `operator`?)
- Quais interfaces? Onde?
- Fluxo de dados: string → tokens → RPN → resultado
- Como cada conceito do Módulo A se manifesta?

### Fase 2: Implementação

Implemente seguindo seu design. Documente desvios.

### Fase 3: Testes

Escreva casos de teste para:
- Expressões válidas: `"3+4*2"`, `"(1+2)*3"`, `"10/3"`
- Erros: divisão por zero, parênteses desbalanceados, operador inválido
- Edge cases: número negativo no início? espaços em branco?

### Fase 4: Retrospectiva

- O design sobreviveu à implementação?
- Onde os conceitos do Módulo A ajudaram?
- Onde você sentiu lacunas?

## Conceitos envolvidos

- Erros sentinela e `errors.Is` — A1 (currency-conversor)
- Erros tipados e `errors.As` — A2 (bank-account)
- Lista ligada como stack — A3 (linked-list)
- Interface implícita e type switch — A4 (shape-interface)
