# Diagnóstico — Docker & DevOps

> Responda antes de continuar a track. Posiciona você no ponto certo.
> Sem consulta. Se não souber, deixe em branco.

## Q1 — Imagem vs Container

Explique a diferença entre imagem e container. Um container pode existir sem
uma imagem? Uma imagem pode existir sem containers?

> Confiança: [1-5]

## Q2 — Camadas (Layers)

O que acontece com o cache de build quando você altera uma linha no meio do
Dockerfile? Por que a ordem das instruções importa?

> Confiança: [1-5]

## Q3 — ENTRYPOINT vs CMD

```dockerfile
ENTRYPOINT ["echo"]
CMD ["hello"]
```

Se eu rodar `docker run <imagem> world`, o que é executado? E se rodar
`docker run <imagem>`?

> Confiança: [1-5]

## Q4 — Volumes

Qual a diferença prática entre bind mount e named volume? Em qual situação
você usaria cada um?

> Confiança: [1-5]

## Q5 — Redes

Dois containers na mesma bridge network podem se comunicar. Como um container
descobre o IP do outro? Existe uma alternativa a usar IPs?

> Confiança: [1-5]

---

## Resultado

| Questão | Acertou? | Módulo relacionado | Ação |
|---------|----------|--------------------| ---- |
| Q1      |          | 1 — Contêineres    |      |
| Q2      |          | 2 — Build          |      |
| Q3      |          | 1 — Contêineres    |      |
| Q4      |          | 3 — Persistência   |      |
| Q5      |          | 3 — Redes          |      |

**Posicionamento:** [definido pelo professor após correção]
