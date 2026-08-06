# learnings

Framework de estudos com zero atrito decisório baseado em evidência científica.
Go Backend e a espinha dorsal; Docker entra como competencia operacional sobre os mesmos projetos.

## Como funciona

Escolha uma track → o professor cuida do resto (próximo desafio, revisões, capstones).

Leia [`framework/README.md`](./framework/README.md) para entender o sistema.

## Tracks

| Track | Progresso | Descrição |
|-------|-----------|-----------|
| [Go Backend](./framework/tracks/go-backend/) | 19 ex, 4 MP, 5 capstones | Go do zero a APIs production-ready |
| [Docker & DevOps](./framework/tracks/docker-devops/) | 6 ex ✅ + 5 ex ⬜, 2 MP, 3 capstones | Contêineres, builds, redes, Compose |
| [Redes na Prática](./framework/tracks/redes-pratica/) | 14 ex, 3 MP, 2 capstones | TCP/UDP, HTTP do zero, DNS, TLS |
| [Linux Systems](./framework/tracks/linux-systems/) | 14 ex, 3 MP, 2 capstones | Processos, sinais, FDs, shell |

**Total:** 58 exercícios, 12 mixed-practices, 12 capstones — 263 arquivos.

## DSA em revisão

- **Linked List Reverse** — ponteiros (prev, current, next)
- **BST Delete** — 3 casos (folha, 1 filho, 2 filhos, sucessor inorder)

## Estrutura

```
framework/              # Meta-framework de estudos
├── AGENTS.md           # Instruções do professor (enxuto, ~80 linhas)
├── metodo.md           # Método de ensino (base científica)
├── templates/          # 6 formatos de desafio
│   ├── implementacao/    debug/    otimizacao/
│   ├── explicacao/       design/    capstone/
└── tracks/             # Trilhas de estudo
    ├── go-backend/       docker-devops/
    ├── redes-pratica/    linux-systems/
spaced-repetition/      # Agenda de revisão espaçada (1d, 3d, 7d, 30d)
```
