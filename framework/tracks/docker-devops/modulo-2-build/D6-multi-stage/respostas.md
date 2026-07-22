# Perguntas - D6 

## Q1

Não, as layers do stage 1 não são incluídas, o que fica são as layers do último stage a ser concluído, as outras são descartadas

confiança = 3


## Q2

Pelo que pesquisei, a imagem scratch não tem nenhum programa instalado, sequer tem bash; É altamente configurável e deve ser vantajosa em cenários muito específicos quando necessita-se de controle total...não é meu caso

confiança = 2

## Q3
Sim, é possível; E existem vários cenários em que é desejável utilizar múltiplos stages:

1) Binário único executável - não faz sentido levar para a produção as dependências do projeto quando se trata de um binário compilado - então dividi-se em duas etapas: build, prod

2) Cenário de testes integrados - queremos rodar os testes, caso não passem, a imagem não deve ser construída...

confiança = 2

## Tasks - extra

Para apontar para um arquivo docker
```
docker -f nome_arquivo_docker -t nome_imagem .
```

O nome da minha imagem gerada de modo single stage foi `single-test` e esse foi o tamanho:

```
edu@edu:~/Codes/learnings/framework/tracks/docker-devops/modulo-2-build/D6-multi-stage$ sudo docker images
                                                                                                                              i Info →   U  In Use
IMAGE                ID             DISK USAGE   CONTENT SIZE   EXTRA
mount-bind:latest    a8c6ae1c5082        224MB         70.6MB    U   
single-test:latest   c55bdda33f99        482MB         97.3MB        
ubuntu:latest        b7f48194d4d8        158MB         45.3MB        

```


Agora usando o multiStage:

```
edu@edu:~/Codes/learnings/framework/tracks/docker-devops/modulo-2-build/D6-multi-stage$ sudo docker images
                                                                                                                              i Info →   U  In Use
IMAGE                ID             DISK USAGE   CONTENT SIZE   EXTRA
mount-bind:latest    a8c6ae1c5082        224MB         70.6MB    U   
multi-test:latest    bb79ac5d06d7       26.1MB         8.29MB        
single-test:latest   c55bdda33f99        482MB         97.3MB    U   
ubuntu:latest        b7f48194d4d8        158MB         45.3MB        
edu@edu:~/Codes/learnings/framework/tracks/docker-devops/modulo-2-build/D6-multi-stage$ 
```



