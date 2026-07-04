# Respostas

1) A tag i significa interactive e a t significa TTY (teletypewriter, ou terminal). Se eu rodar apenas com a tag `-i` eu terei a interatividade mas não vou receber o retorno de que estou no terminal, segue a saída:

```

edu@edu:~$ sudo docker run -i ubuntu bash
 

ls       
bin
boot
dev
etc
home
lib
lib64
media
mnt
opt
proc
root
run
sbin
srv
sys
tmp
usr
var
exit

```
```
```


Já rodando com a tag `-t` eu irei para o terminal mas não terei interatividade, comandos não serão interpretados pelo shell do nosso ubuntu

2) O comando `docker run -it` instancia um novo container do zero. O comando `docker exec` serve para executar um container já criado e ativo. O comando `docker start` serve para retomar a execução de um container a partir do ponto de saída anterior.


3) Não consegui criar nenhum arquivo dentro do ubuntu, talvez tenha sido ignorância minha:

```

root@8683006551a9:/# touch /temp/eu-estou-aqui.txt
touch: cannot touch '/temp/eu-estou-aqui.txt': No such file or directory
root@8683006551a9:/# ls 
bin   dev  home  lib64  mnt  proc  run   srv  tmp  var
boot  etc  lib   media  opt  root  sbin  sys  usr
root@8683006551a9:/# touch ./temp/eu-estou-aqui.txt
touch: cannot touch './temp/eu-estou-aqui.txt': No such file or directory
root@8683006551a9:/# exit
```
```
```

4) Não sei responder, na minha cabeça seria a mesma coisa, afinal o `docker run <image_name>` iniciará o container e o `docker exec <container_id>` iniciará outra "visualização" do mesmo container, compartilhando os mesmos recursos que o container dispoe...
