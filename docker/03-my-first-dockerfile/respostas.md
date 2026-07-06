# Respostas 03

1) `CMD` diz o comando que irá rodar por padrão durante a inicialização do container, o `ENTRYPOINT` definirá qual será a "aplicação" equivalente daquele container, por exemplo, é possível definir que o `curl ` será o `ENTRYPOINT` e quando rodarmos o nosso container podemos definir um website para fazer o curl como segundo argumento: `docker run <nome_container> website.com`

Algo como: 
```bash
ENTRYPOINT ["curl"]
```

> obs: o curl precisaria estar instalado nesse container


2) No caso específico da instrução CMD o último comando sobrescreve o primeiro. Em alguns casos as instruções são "concatenadas", como no caso da instrução `WORKDIR`

3) O `WORKDIR` define o diretório de trabalho do container, facilitando a organização. Se o diretorio não existe, é criado pelo docker. Sem um diretorio de trabalho a aplicação trabalha no padrão da imagem, no caso das imagens linux no diretório raíz `/`

4) Caso algo tenha mudado a partir de uma determinada instrução, o que acontece é que o docker irá refazer todas as instruções que se seguem, como num efeito dominó. Por isso, é importante escolher a ordem com que escrevemos as instruções, aquilo que corre mais riscos de mudar deve ficar após o que é imutável (preciso destudar mais sobre caching do docker, não captei muito bem)

5) A tag `--rm` apaga o container assim que ele é encerrado

6) Mostra como o container foi construído, comandos utilizados para montar a imagem dele e as camadas

7) O comando `ADD ` é meio que um `COPY ` bombado, ele pode baixar arquivos de endereços e descompactar arquivos do tipo tar.gz (seria interessante utilizar depois)
