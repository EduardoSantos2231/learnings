# Framework de Estudos

> Zero atrito decisório. Você escolhe uma track, o framework cuida do resto.

## Como funciona

```
Track → Diagnóstico → Desafios → Rotação → Capstone → Revisão Espaçada
```

1. **Escolha uma track** (única decisão que você toma)
2. **Diagnóstico inicial** — 5-10 perguntas para posicionar você no ponto certo do roadmap
3. **Inicie a sessão** — `tracking session` escolhe desafio ou revisão
4. **Revisões progressivas** — blocos práticos reaparecem em 1d, 7d e 30d

## Formato de uma sessão típica

```
Professor: "Track: go-backend. Próximo: Módulo C — BST Delete.
           Revisão pendente: rate-limiter (7 dias). Revisa antes ou vai direto?"

[Você só decide o ritmo, nunca o conteúdo]
```

## Tracks disponíveis

| Track | Status | Descrição |
|-------|--------|-----------|
| go-backend | Principal | Go do zero a APIs production-ready |
| docker-devops | Complementar | Operar e empacotar os serviços Go |
| redes-pratica | Nova | TCP/UDP, HTTP do zero, DNS, TLS |
| linux-systems | Nova | Processos, sinais, FDs, shell mínimo |
| dsa | Em revisão | Estruturas de dados e algoritmos |
| typescript-fullstack | Futura | TypeScript full-stack |

## Estrutura de uma track

```
tracks/<nome>/
├── roadmap.md          # Sequência linear com flag de posição atual
├── diagnostic.md       # Quiz de entrada
├── modulo-X-nome/      # Exercícios monotemáticos
├── mixed-practice/     # Sessões intercaladas (a cada 4 desafios)
├── capstones/          # Projetos integradores (a cada 5 desafios)
└── correcoes.md        # Erros acumulados da track
```

## Formatos de desafio

Os desafios rotacionam entre 6 formatos para variabilidade de prática:

| Formato | O que você faz |
|---------|---------------|
| Implementação | Codar do zero a partir de uma especificação |
| Debug | Corrigir bugs em código quebrado |
| Otimização | Melhorar código funcional mas ineficiente |
| Explicação | Explicar conceito como se ensinasse alguém |
| Design | Projetar API/arquitetura antes de implementar |
| Capstone | Projeto multi-conceito integrador |

## Método de ensino

Baseado em evidência científica. Detalhes em [`metodo.md`](./metodo.md).

- Prática de evocação (perguntas pós-exercício)
- Repetição espaçada (revisitas programadas)
- Prática intercalada (mixed-practice entre módulos)
- Síntese cumulativa (capstones)
- Dificuldades desejáveis (sem resposta pronta)
- Fading de suporte (menos scaffolding ao longo dos módulos)

## Para começar

1. Leia `metodo.md` e `AGENTS.md` para entender o papel do professor
2. Escolha uma track: `framework/tracks/<track>/roadmap.md`
3. O professor assume o controle da sessão
