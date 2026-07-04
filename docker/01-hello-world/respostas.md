# Respostas hello world docker

## 2

O hello world apresenta todo o ciclo de vida de execução do docker:
1) contactar o daemon do docker
2) daemon faz um pull do docker hub de acordo com a imagem que selecionamos (hello-world:latest)
3) o daemon cria o container e roda
4) O output é enviado para o terminal 

## 3
#### Vendo todos os containers:
Para listar container que já rodaram eu uso: `docker ps -a`. Esse comando lista os stauts dos processos, a flag `-a` indica que todos devem ser listados

#### Removendo container antigo:
Para remover um container que já não me serve mais: `docker rm <nome_container or id>`

#### Checando containers ativos:
Para verificar os containers ativos, basta buscar por: `docker ps `, se após o header não houver mais informações, então não existem containers rodando

## Perguntas extras

> O comando docker ps mostra os process status, e a tag -a indica que todos os process status manejados pelo docker devem ser exibidos

> Toda vez que utilizamos um `docker run` estamos criando uma nova instância da aplicação. `docker start` resume a aplicação anterior, reaproveitando o container

> Não é possível entrar num container que já finalizou, se ele não está ativo não existe meio de comunicação ativo
