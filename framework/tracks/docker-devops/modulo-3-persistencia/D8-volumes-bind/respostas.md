# Respostas

1) Num bind mount apontamos para um diretório e tudo que for modificado é modificado e refletido também dentro do container em tempo real. No entanto, é preciso prover o path absoluto toda vez que queremos que um container aponte para o "bind" que criamos. O volume nomeado provê um gerenciamento automático do docker, que cria volumes em um local específico, e ao nomear podemos fazer com que diferentes containers "consumam" o mesmo container sem nos precoupar com o caminho;

Bind mount é útil em desenvolvimento (código fonte editado no host reflete no container), volume nomeado é melhor em produção (portabilidade, backup com docker volume)


2) O erro é explicitado, o container sequer chega a rodar

3) Em ambos os casos as modificações acontecem em tempo real, então sim, os dois containers apontando para o mesmo volume iriam enxergar modificações feitas; Se for um bind mount o mesmo acontece

4) Serve para deletar todos os volumes que não estão sendo utilizados por NENHUM container. Existem variações como `docker image prune` -> remover imagens que não são usadas por nenhum container...etc
