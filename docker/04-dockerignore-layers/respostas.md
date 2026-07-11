# Respostas

1) O docker ignore serve ao mesmo propósito que o gitignore, ignorar arquivos que não devem ser "vistos" pelo docker. Logo, por exemplo, ao utilizar a instrução `COPY . .` o docker não veria os arquivos que inserimos no `.dockerignore` para serem copiados - veria mas seriam ignorados.  

2) Se o docker ignore for esquecido todos os arquivos que constarem dentro do diretório atual do arquivo `Dockerfile` serão colocadas no build context da imagem, o que é uma má prática com certas aplicações, como por exemplo aplicações que possuem algumas centenas de dependências. Elas seriam todas copiadas para dentro da imagem, ao invés de instaladas. 


3) Por que as camadas são construídas uma em cima da outra, por assim dizer. Se a camada 3 foi alterada o docker força a re-execução da instrução 4.


4) Pelo que observei a melhor ordem é:

```

FROM -> sem indicar a imagem base nem tem sequência

RUN -> quais comandos rodar nessa imagem base


WORKDIR -> qual diretório padrão de trabalho nesse container

COPY -> o que copiar do diretório atual pra dentro da imagem

ENTRYPOINT -> defina o ponto de alterada

CMD -> o que rodar nesse ponto de entrada

```

A ordem pode conter algumas pequenas variações, mas o que muda MENOS vem primeiro

5) A primeira instrução copia tudo a partir do diretório atual, enquanto a segunda explicita o que deverá ser copiado, copiando apenas o index.html para dentro da nossa imagem 

6) Ele mostra as camadas que foram construídas. A coluna image mostra o ID local da imagem, o created mostra a data, e o created by mostra o comando que gerou aquela layer

7) Colocar comandos juntos no mesmo run evita que criem layer diferentes, daí ambos comandos são unificados numa única layer 
