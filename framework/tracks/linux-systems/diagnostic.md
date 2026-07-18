# Diagnóstico — Linux Systems

> Sem consulta. Se não souber, deixe em branco.

## Q1 — Processos

O que acontece quando você executa `ls` no terminal? Descreva o fluxo:
shell → fork → exec → wait.

> Confiança: [1-5]

## Q2 — Sinais

O que acontece quando você aperta Ctrl+C no terminal? Qual sinal é enviado?
Como um programa pode interceptá-lo?

> Confiança: [1-5]

## Q3 — File Descriptors

O que são file descriptors 0, 1, e 2? O que este comando faz?
```bash
ls > output.txt 2>&1
```

> Confiança: [1-5]

## Q4 — /proc

O que é o diretório `/proc`? O que você encontraria em `/proc/self/fd/`?

> Confiança: [1-5]

## Q5 — Hardlink vs Symlink

Qual a diferença entre hardlink e symlink? Se eu deleto o arquivo original,
qual dos dois ainda funciona?

> Confiança: [1-5]

---

## Resultado

| Questão | Acertou? | Módulo relacionado |
|---------|----------|--------------------|
| Q1      |          | 1 — Processos      |
| Q2      |          | 1 — Sinais         |
| Q3      |          | 2 — FDs            |
| Q4      |          | 3 — Filesystem     |
| Q5      |          | 3 — Filesystem     |

**Posicionamento:** [definido pelo professor]
