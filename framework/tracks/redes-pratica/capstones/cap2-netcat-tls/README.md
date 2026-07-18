# Capstone 2 — Netcat com TLS + DNS

> Síntese de todos os módulos | Sem scaffolding

## Contexto

Construa uma ferramenta estilo netcat que resolve domínios via DNS,
conecta via TCP ou UDP, opcionalmente usa TLS, e faz streaming bidirecional.

## Requisitos

1. CLI: `./netcat [--tls] [--udp] host port`
2. Resolve `host` via DNS (seu resolver do Módulo 3 com cache)
3. Conecta via TCP (padrão) ou UDP (flag --udp)
4. Se --tls: faz handshake TLS e verifica certificado
5. Streaming bidirecional: stdin → socket, socket → stdout
6. Graceful shutdown com Ctrl+C (SIGINT)

## Conceitos envolvidos

- TCP/UDP sockets — R1, R3
- DNS resolver + cache — R9, R10, R11
- TLS client — R12, R13
- Sinais — conceito de Linux Systems
- Streaming I/O bidirecional
