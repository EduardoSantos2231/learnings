# Roadmap — Redes na Prática

> Track de redes implementando protocolos do zero em Go com `net` da stdlib.
> Siga a ordem. O professor anuncia o próximo automaticamente.

## Posição atual

| Campo | Valor |
|-------|-------|
| Módulo atual | 1 — Sockets TCP/UDP |
| Último concluído | R1-echo-server |
| Próximo desafio | R2-chat-tcp |
| Próximo formato | Implementação |

---

## Módulo 1 — Sockets TCP/UDP

> Scaffolding: alto

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| R1 | echo-server | Implementação | `net.Listen`, `net.Dial`, TCP handshake | ✅ |
| R2 | chat-tcp | Implementação | Goroutines por conexão, broadcast | ⬜ |
| R3 | udp-echo | Implementação | `net.ListenUDP`, datagramas, perda de pacotes | ⬜ |
| R4 | timeout-retry | Otimização | Deadline, retry, buffer sizing | ⬜ |

## Mixed Practice 1

| # | Desafio | Template | Status |
|---|---------|----------|--------|
| MP1 | tcp-vs-udp | Mixed Practice | ⬜ |

---

## Módulo 2 — HTTP do Zero

> Scaffolding: alto

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| R5 | http-parser | Implementação | Parse de requisição HTTP/1.1 bruta | ⬜ |
| R6 | http-server | Implementação | Servidor HTTP mínimo, resposta a GET | ⬜ |
| R7 | keep-alive | Design | Conexões persistentes, `Connection: keep-alive` | ⬜ |
| R8 | chunked | Implementação | `Transfer-Encoding: chunked` | ⬜ |

## Mixed Practice 2

| # | Desafio | Template | Status |
|---|---------|----------|--------|
| MP2 | cliente-http | Mixed Practice | ⬜ |

## Capstone 1 — Proxy HTTP Reverso

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| C1 | proxy-http | Capstone | R1-R8 | ⬜ |

---

## Módulo 3 — DNS Resolver

> Scaffolding: médio

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| R9 | dns-wire | Implementação | Construir query DNS no wire format | ⬜ |
| R10 | dns-resolver | Implementação | Resolução iterativa (root → TLD → auth) | ⬜ |
| R11 | dns-cache | Design | Cache de respostas com TTL do record | ⬜ |

## Mixed Practice 3

| # | Desafio | Template | Status |
|---|---------|----------|--------|
| MP3 | debugging-dns | Mixed Practice | ⬜ |

---

## Módulo 4 — TLS na Prática

> Scaffolding: médio

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| R12 | certs | Explicação | Gerar, ler, validar certificados X.509 | ⬜ |
| R13 | tls-server | Implementação | Servidor TLS com `crypto/tls` | ⬜ |
| R14 | tls-mitm | Debug | Proxy MITM educativo que intercepta TLS | ⬜ |

## Capstone 2 — Netcat com TLS + DNS

| # | Desafio | Template | Conceitos | Status |
|---|---------|----------|-----------|--------|
| C2 | netcat-tls | Capstone | R1-R14 | ⬜ |
