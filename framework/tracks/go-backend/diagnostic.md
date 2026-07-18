# Diagnóstico — Go Backend

> Responda antes de começar a track. Isso posiciona você no ponto certo do roadmap.
> Sem consulta. Se não souber, deixe em branco (não chute).

## Q1 — Erros

O que `errors.Is` faz que `==` não faz? Dê um exemplo em que `==` falharia.

> Confiança: [1-5]

## Q2 — Interfaces

```go
var w io.Writer
var buf *bytes.Buffer
w = buf
fmt.Println(w == nil) // true ou false?
```

O que imprime e por quê?

> Confiança: [1-5]

## Q3 — Concorrência

Você tem 100 URLs para buscar. Escreva o esqueleto (assinaturas) de uma solução
concorrente com limite de 5 requisições simultâneas e timeout de 2 segundos por URL.

> Confiança: [1-5]

## Q4 — Slices

```go
s := make([]int, 0, 5)
s = append(s, 1, 2, 3)
t := s[1:3]
t = append(t, 4, 5, 6, 7)
fmt.Println(s) // ???
```

O que imprime? Explique o mecanismo por trás.

> Confiança: [1-5]

## Q5 — HTTP

Escreva um handler HTTP em Go (stdlib) que:
- Recebe POST com JSON `{"name": "..."}`
- Se `name` estiver vazio, retorna 400
- Se OK, retorna 201 com `{"id": 1, "name": "..."}`
- Usa um mutex para proteger o contador de IDs

Escreva só as assinaturas e a estrutura do handler (não precisa do main).

> Confiança: [1-5]

---

## Resultado

| Questão | Acertou? | Módulo relacionado | Ação |
|---------|----------|--------------------| ---- |
| Q1      |          | A — Fundamentos    |      |
| Q2      |          | E — Armadilhas     |      |
| Q3      |          | B — Concorrência   |      |
| Q4      |          | C — Estruturas     |      |
| Q5      |          | D — HTTP & APIs    |      |

**Posicionamento:** [definido pelo professor após correção]
