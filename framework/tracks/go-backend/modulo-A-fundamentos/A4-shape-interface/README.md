# A4 — Shape Interface

> Implementacao | 45 min | Go stdlib

## Objetivo

Modele formas geometricas usando uma interface pequena e calculando area.

## Faca

1. Defina uma interface com o comportamento necessario.
2. Implemente circulo e retangulo.
3. Some areas recebendo apenas a interface.
4. Trate forma desconhecida sem panic.

## Restricoes

- Use interface implicita.
- Nao use reflexao.

## Pronto quando

- Novas formas podem ser adicionadas sem alterar o somador.
- Testes verificam as areas.
- O type switch, se usado, trata o caso desconhecido.

## Responda

- Por que a interface deve ser definida por quem a consome?
- Quando um type switch seria desnecessario?

> Confianca: [1-5]
