# Mixed Practice 1 — Escolha de Ferramentas

> Interleaving: você não sabe de antemão qual padrão usar. Decida.

## Cenário 1: Sistema de Desfazer/Refazer

Você precisa implementar undo/redo para um editor de texto.
Cada ação do usuário é registrada e pode ser desfeita ou refeita.

**Pergunta:** Qual estrutura usar — Stack, Queue, ou LinkedList?
Implemente a operação de undo.

## Cenário 2: Log de Erros com Níveis

Você precisa de um sistema de log com níveis (DEBUG, INFO, WARN, ERROR).
Erros de parsing de configuração devem ser detectáveis com `errors.Is`.
Erros de rede devem ser inspecionáveis com `errors.As` para extrair o status code.

**Pergunta:** Como estruturar os erros? Sentinel errors, tipo customizado, ou ambos?
Implemente.

## Cenário 3: Plugin System

Você está construindo um sistema que suporta plugins: cada plugin processa
um tipo de arquivo (PDF, imagem, texto). No futuro, novos formatos serão adicionados.

**Pergunta:** Interface com método único ou type switch? Qual escala melhor?
Implemente a abordagem escolhida e justifique.
