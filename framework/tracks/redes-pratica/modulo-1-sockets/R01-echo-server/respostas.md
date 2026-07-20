# R01 - Respostas

### Q01

O three-way handshake acontece quando estabelecemos uma conexão, ela acontece em três etapas:

- Sync (quem deseja estabelecer a conexão informa que deseja sincronizar)
- Sync + Ack (quem recebe o pedido informa que aceita e informa que quer também se conectar) 
- Ack (quem solicitou a conexão agora confirma a sincronização)

Durante a troca de mensagens eles definem também um número de sequência aleatório para manter registro dos pacotes que irão enviar;

O three-way handshake acontece antes de o nosso código aceitar novas mensagens, e quem gerencia isso é o sistema operacional;

### Q02 

Pelo que pesquisei o Sistema Operacional cuida disso mantendo numa fila e buffers; Dados enviados antes de a conexão ser aceita ficam "hanging" até que a conexão seja aceita pelo programa e o envio começa.

### Q03

Não existe tratamento de erros, e nesse cenário de 1000 conexões os dados podem ser perdidos num control C; Não existe tolerância pra erros e gerencimento de recurso;

