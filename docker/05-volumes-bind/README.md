# volumes-bind

Entender a diferença entre bind mount e volume nomeado, e como dados persistem (ou não) em containers.

## Tarefas

### 1 — Dados morrem com o container

Crie um arquivo `dados.txt` com o conteúdo "Olá mundo" dentro de um container. Depois remova o container e prove que o arquivo sumiu.

### 2 — Bind mount: dados vivos do lado de fora

- Crie um diretório `dados/` na sua máquina (fora do Docker) com um arquivo `mensagem.txt`
- Monte esse diretório como bind mount num container (`--mount type=bind,src=$(pwd)/dados,target=/app/dados` ou `-v $(pwd)/dados:/app/dados`)
- Dentro do container, leia o arquivo e depois escreva um novo arquivo
- Saia do container: o arquivo escrito está na sua máquina?

### 3 — Volume nomeado: dados gerenciados pelo Docker

- Crie um volume nomeado com `docker volume create meus_dados`
- Monte o volume num container e escreva um arquivo lá dentro
- Remova o container e crie um **segundo** container montando o mesmo volume
- O arquivo ainda está lá?

### 4 — Limpeza

- Liste os volumes com `docker volume ls`
- Remova o volume criado com `docker volume rm meus_dados`

## Extra

- O que acontece se você montar um bind mount num diretório que **já tem conteúdo** dentro da imagem? (ex: montar algo em `/usr/share/nginx/html` numa imagem nginx)
