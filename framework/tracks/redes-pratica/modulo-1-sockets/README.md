# Módulo 1 — Sockets TCP/UDP

> Scaffolding: alto | Go `net` puro

**Objetivo:** Entender comunicação de rede no nível mais baixo: abrir sockets,
conectar, enviar/receber bytes, diferenças entre TCP e UDP.

## Desafios

| # | Desafio | O que você constrói |
|---|---------|-------------------|
| R1 | echo-server | Servidor TCP que ecoa o que recebe. Cliente que envia mensagem e lê resposta |
| R2 | chat-tcp | Chat multithreaded: cada conexão é um cliente, servidor faz broadcast |
| R3 | udp-echo | Servidor/cliente UDP. Entender datagramas, perda, ordenação |
| R4 | timeout-retry | Timeout, retry com backoff, escolha de buffer size |

## Conceitos ao final do módulo

- `net.Listen("tcp", ":port")` — aceitar conexões
- `net.Dial("tcp", "addr:port")` — conectar
- `net.Conn` — interface para ler/escrever em conexão
- TCP: stream, ordenado, confiável, handshake de 3 vias
- UDP: datagrama, não ordenado, não confiável, sem conexão
- `SetDeadline`, `SetReadDeadline` — timeout em operações de rede
- Buffer sizing: tamanho do buffer de leitura afeta throughput
- Backpressure: o que acontece quando o leitor é mais lento que o escritor
