# Mixed Practice 1 — TCP vs UDP

> Interleaving: para cada cenário, escolha entre TCP ou UDP e justifique.

## Cenário 1: Streaming de Vídeo

Você está construindo uma plataforma de streaming de vídeo ao vivo.
Pacotes perdidos não devem ser retransmitidos (o vídeo continua).
Latência é mais importante que integridade.

**Pergunta:** TCP ou UDP? Por quê? Implemente um protótipo mínimo.

## Cenário 2: Transferência de Arquivo

Você precisa enviar um arquivo de 1 GB pela rede. Cada byte importa.
O receptor deve ter exatamente o mesmo arquivo que o emissor.

**Pergunta:** TCP ou UDP? Por quê? O que você precisaria implementar se usasse UDP?

## Cenário 3: Jogo Multiplayer

Jogo de tiro online: posição dos jogadores atualizada 60x por segundo.
Se um pacote de posição se perder, o próximo corrige (interpolação).
Mas comandos como "disparou" não podem ser perdidos.

**Pergunta:** Um protocolo ou dois? Qual para posição e qual para comandos?
