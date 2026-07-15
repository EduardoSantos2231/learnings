## Perguntas

1. Qual a diferença entre bind mount e volume nomeado? Em que cenário você usaria cada um?

2. O bind mount depende da estrutura de diretórios da máquina host. O que acontece se o diretório fonte do bind mount não existir? (Teste com `docker run`)

3. Se dois containers montam o **mesmo volume nomeado** e um escreve um arquivo, o outro enxerga? E se for um bind mount apontando pro mesmo diretório?

4. `docker volume prune` — o que ele faz? Quando é seguro rodar?
