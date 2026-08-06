# A2 — Bank Account

> Implementacao | 45 min | Go stdlib

## Objetivo

Implemente uma conta bancaria com saldo protegido por regras simples.

## Faca

1. Crie `Deposit` e `Withdraw` com pointer receiver.
2. Rejeite valor zero ou negativo.
3. Rejeite saque maior que o saldo.
4. Retorne um erro tipado para falha de saque.

## Pronto quando

- O saldo nunca fica inconsistente.
- O chamador identifica o erro com `errors.As`.
- Testes cobrem deposito, saque e falhas.

## Responda

- Por que os metodos alteram um pointer receiver?
- O que deve ser parte da API publica?

> Confianca: [1-5]
