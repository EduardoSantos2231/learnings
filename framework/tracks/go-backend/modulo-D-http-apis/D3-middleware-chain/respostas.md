# Repostas 

1) O Handler explicita o método e a assinatura que o Handler deve conter e quando usamos http.HandleFunc estamos fazendo um typecasting, pois o HandlerFunc implementa esse método com a assinatura solicitada pela interface;

2) O que estiver abaixo e executado quando o próximo handler desempilhar, ou não...por exemplo supondo que temos a seguinte ordem:  auth ->  jsonHeaderSetter

Se o auth quiser logar algo, ele deverá experar o retorno do json, pois ele chamou o método implementado pelo handler "next" declarado na assinatura 

3) Ele deve ser chamado por último, envolvendo todos os middlewares, caso o erro estoure ele cai no recoverer e exibe uma resposta

4) Por que não queremos nem que o user entre na aplicação caso ele não tenha o token de autorização, assim economizamos recursos e diminuímos a superfície que o user pode acessar. 

5) Por que o options é um método http que não realiza ações e é enviado pelo navegador para saber quais métodos são permitidos e headers antes de o navegador fazer requisições mais complexas (put, por exemplo)

6) O header vary explicita o seguinte: se esse campo mudar, precisamos gerar uma nova resposta e o cache que tínhamos não servirá como resposta. O vary: * informa que nada deverá ser cacheado. Supondo então um header Vary: origin, teríamos então definido que quando a origem mudar, os dados cacheados não servirão. Se nós definimos no CORS que todos podem acessar nosso endpoint, então ele é público, se ele é público e não varia, a resposta é a mesma, logo, deveríamos cachear as resposes sim.
