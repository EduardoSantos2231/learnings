# Perguntas — 04 dockerignore-layers

1. O que o `.dockerignore` faz? Como ele se compara a um `.gitignore`?

2. Se você esquecer o `.dockerignore`, qual o impacto no build e na imagem final?

3. Por que o Docker reexecuta uma layer e **todas as seguintes** se algo naquela layer muda?

4. Qual a ordem ideal das instruções no `Dockerfile` para maximizar o reaproveitamento de cache?

5. `COPY . .` vs `COPY index.html .` — qual a diferença prática para o cache de build?

6. O que `docker history` mostra? Como interpretar o tamanho de cada layer?

7. Por que é boa prática incluir `rm -rf /var/lib/apt/lists/*` no mesmo `RUN` do `apt-get install`?
