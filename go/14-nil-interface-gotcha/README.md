# nil interface gotcha

## Contexto

Em Go, uma interface é representada internamente como um par **(type, value)**.
Ambos precisam ser `nil` para que a interface seja `nil`. Se você atribuir um
ponteiro `nil` a uma interface, a interface **não** será `nil` — porque o `type`
do par não é nil.

Esse é um dos gotchas mais clássicos de Go, e já pegou todo mundo pelo menos
uma vez em entrevista ou em produção.

---

## Tarefa 1 — Demonstrar o gotcha

Crie `main.go` que:

1. Declare `var w io.Writer` (nil interface)
2. Print se `w == nil` → deve printar `true`
3. Declare `var buf *bytes.Buffer` (nil pointer)
4. Atribua `w = buf`
5. Print `w == nil` → deve printar `false` (o gotcha!)
6. Tente chamar `w.Write([]byte("hello"))` e veja o panic (`nil pointer dereference`)
7. Capture o panic com `recover()` e explique o que aconteceu

Saída esperada (aproximada):
```
w == nil: true
buf == nil: true
w == nil after assignment: false
about to call w.Write...
PANIC: runtime error: invalid memory address or nil pointer dereference
```

## Tarefa 2 — Correção

Crie uma função segura `safeWrite(w io.Writer, data []byte)` que:
- Use **reflexão** (`reflect.ValueOf(w).IsNil()`) para checar se o valor interno da interface é nil
- Só chame `w.Write(data)` se não for nil
- Retorne `(n int, err error)` igual ao `Write` normal

---

## Arquivos esperados

```
14-nil-interface-gotcha/
├── README.md       ← este arquivo
├── perguntas.md    ← questões conceituais
├── respostas.md    ← SUAS respostas (crie você)
└── main.go         ← SEU código
```

Depois de implementar, me avise: "respondi, dá uma olhada".
