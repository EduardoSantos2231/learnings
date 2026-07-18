# L4 — zombies

> Template: Debug | Scaffolding: alto

Crie e observe processos zumbis, depois corrija-os.

## Tarefas

1. Crie um programa que spawne um filho, o filho saia imediatamente, o pai durma 60s sem chamar `Wait()`
2. Enquanto o pai dorme: observe o zumbi com `ps aux | grep defunct`
3. **Corrija:** adicione `Wait()` após o filho sair, observe o zumbi desaparecer
4. **Explique:** e se o pai sair antes do filho? Quem vira o novo pai?

## Verificação

- `ps aux | grep defunct` mostra o zumbi antes do fix
- Após adicionar `Wait()`, o zumbi não aparece mais
