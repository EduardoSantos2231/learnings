# Respostas

1) No `ENTRYPOINT` eu declaro qual será o programa padrão de entrada para o meu container, qualquer argumento passado ao executar o container é "injetado" dentro do programa selecionado como `ENTRYPOINT`. O `CMD` indica qual comando irá ser executado ao iniciar o container. Portanto, é possível definir diferentes responsabilidades de ENTRYPOINT e CMD para garantir mais flexibilidade, por exemplo, tendo um default no cmd e permitindo que o user insira um argumento

2) Ele irá executar o `ENTRYPOINT` e executar o default do programa para quando não há argumentos. Se rodar a imagem com a tag --help ela exibe o que o comando `cat` faz:

```
Concatenate FILE(s), or standard input, to standard output
With no FILE, or when FILE is -, read standard input.

Usage: cat [OPTION]... [FILE]...

Options:
  -A, --show-all          equivalent to -vET
  -b, --number-nonblank   number nonempty output lines, overrides -n
  -e                      equivalent to -vE
  -E, --show-ends         display $ at end of each line
  -n, --number            number all output lines
  -s, --squeeze-blank     suppress repeated empty output lines
  -t                      equivalent to -vT
  -T, --show-tabs         display TAB characters at ^I
  -v, --show-nonprinting  use ^ and M- notation, except for LF (\n) and TAB (\t)
  -u                      (ignored)
  -h, --help              Print help
  -V, --version           Print version
```


3) O `CMD` sem entrypoint, ao receber argumentos, eles substituem o cmd por inteiro 

```
edu@edu:~/Codes/learnings/docker/03b-entrypoint-vs-cmd$ sudo docker run meu-site /etc/os-release
docker: Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed:
 unable to start container process: error during container init: exec: "/etc/os-release": permission denied

Run 'docker run --help' for more information
```
4) Usando a tag `--entrypoint` é possível sobrescrever o entrypoint definido na imagem


5) Então nesse caso seria importante utilizar a instrução `cmd`; 

```
ENTRYPOINT ["ping"]
CMD ["google.com"]
```
