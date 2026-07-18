# Módulo 2 — HTTP do Zero

> Scaffolding: alto | Go `net` + `bufio`

**Objetivo:** Implementar um servidor HTTP/1.1 do zero — parse de requisição,
construção de resposta, keep-alive, chunked encoding. Nada de `net/http`.

## Desafios

| # | Desafio | O que você constrói |
|---|---------|-------------------|
| R5 | http-parser | Parser de requisição HTTP/1.1: method, path, headers, body |
| R6 | http-server | Servidor HTTP mínimo que serve arquivos estáticos e rotas GET |
| R7 | keep-alive | Conexões persistentes: reusar conexão TCP para múltiplas requests |
| R8 | chunked | Transfer-Encoding: chunked — enviar resposta sem Content-Length |

## Conceitos ao final do módulo

- Formato de request HTTP: `METHOD /path HTTP/1.1\r\nHeader: value\r\n\r\nbody`
- Headers: `Host`, `Content-Length`, `Connection`, `Transfer-Encoding`
- Status line: `HTTP/1.1 200 OK\r\n`
- `Connection: keep-alive` — reusar conexão (HTTP/1.1 padrão)
- `Connection: close` — fechar após resposta (HTTP/1.0 padrão)
- Chunked encoding: `tamanho_em_hex\r\ndados\r\n...0\r\n\r\n`
- `bufio.Reader` — leitura bufferizada de conexão
