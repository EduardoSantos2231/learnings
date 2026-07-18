# Módulo 4 — TLS na Prática

> Scaffolding: médio | Go `crypto/tls`, `crypto/x509`

**Objetivo:** Entender TLS na prática: gerar certificados, construir servidor TLS,
entender a cadeia de confiança, e fazer um MITM educativo.

## Desafios

| # | Desafio | O que você constrói |
|---|---------|-------------------|
| R12 | certs | Gerar CA root, assinar certificado de servidor, validar cadeia |
| R13 | tls-server | Servidor TLS com `crypto/tls`, cliente que verifica certificado |
| R14 | tls-mitm | Proxy MITM: intercepta conexão TLS, gera certificado on-the-fly |

## Conceitos ao final do módulo

- X.509: formato de certificado (Subject, Issuer, SAN, NotBefore/After)
- CA (Certificate Authority): emite e assina certificados
- Cadeia de confiança: root CA → intermediate CA → leaf certificate
- `crypto/x509` — `CreateCertificate`, `ParseCertificate`, `Verify`
- `crypto/tls` — `Listen`, `Dial`, `Config`
- TLS handshake: ClientHello → ServerHello → Certificate → ServerHelloDone
- SNI (Server Name Indication): um IP, múltiplos certificados
- MITM: proxy que gera certificado falso assinado por CA própria
- Por que MITM funciona se a vítima não instalar a CA do atacante
