# Respostas

1) O select statement sem um default é um bloco que bloqueia a execução até que um dos casos seja atendido. Se um `default` é inserido sem nenhum caso, a execução acontece e saímos do bloco `select` já que não há nenhuma instrução abaixo dele. Caso seja haja um bloco `for` envolvendo o `select` então a CPU executará o laço for repetidas vezes até que algum dos cases seja atendido.

2) O padrão correto é utilizar o bloco `default` apenas quando existe uma ação que queremos tomar imediatamente no instante em que os casos acima dele não forem atendidos. Uma ação pontual e não bloqueante. O padrão correto, portanto, é considerar o `<- ctx.Done()` como um caso e o envio/leitura de dados por um channel também como um `case`. 
O Fan-In com send no default falha porque, se o buffer encher, o send bloqueia fora do select e a goroutine não responde mais a ctx.Done().

3) Se o buffer ainda estiver cheio o cancelamento ainda responde, pois o `select` aguarda até que algum dos estados dos channels mude ou seja atentido. Daí a ação ficaria bloqueada até que o envio fosse feito. Sem select, um send em canal cheio bloqueia a goroutine para sempre, tornando-a insensível a cancelamento — só um receiver drenando o canal a liberaria 
