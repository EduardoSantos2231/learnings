# Mixed Practice 3 — Debugging DNS

> Interleaving: use ferramentas reais + seu resolver para diagnosticar problemas.

## Cenário 1: Domínio Não Resolve

`dig example.com` retorna NXDOMAIN. Mas o site abre no navegador.

**Pergunta:** Liste 3 possíveis causas. Como verificar cada uma?
Use `dig`, `nslookup`, e seu resolver do R10.

## Cenário 2: Resolução Lenta

`curl google.com` demora 5 segundos para começar a transferência.

**Pergunta:** O gargalo é DNS? Como provar? Se for DNS: cache?
Servidor lento? Timeout de IPv6?

## Cenário 3: DNS Spoofing (teórico)

Um atacante na mesma rede responde mais rápido que o servidor DNS real.

**Pergunta:** Como isso funciona? Que defesas existem (DNSSEC, DoH, DoT)?
Seu resolver do R10 seria vulnerável?
