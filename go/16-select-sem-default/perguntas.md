## Perguntas

1. O que acontece se você colocar um `default:` vazio no `select` do `fanInCancelable`? O comportamento muda? Por que?

2. No Worker Pool original (`03-worker-pool`), o `GenWorker` tinha `select` com `ctx.Done()` e `ResultChan <- result` como cases, sem `default`. No `ProducerA` do Fan-In (`05-fan-in`), o send estava dentro do `default`. Qual dos dois padrões está correto para responder a cancelamento? Explique.

3. No `sendOrCancel`, se o buffer do canal estiver cheio, a função ainda responde a cancelamento? E se não tivesse `select` nenhum e fosse só `ch <- val`?
