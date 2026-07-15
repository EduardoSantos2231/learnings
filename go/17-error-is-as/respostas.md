# Respostas

1) Por que o erro foi envolvido dentro de outra formatação que não é a previamente vista quando o erro foi criado pela primeira vez. O verbo `%w` serve justamente para envolver o erro à nova formatação; `errors.Is` percorre a cadeia de erros "descasacando" e comparando o retorno, `==` compara diretamente os valores dos dois ponteiros e não compara os valores;

2) `errros.As` precisa escrever o erro correspondente no alvo, modificando o seu valor. Caso não seja um pointer não será possível modificar o valor de maneira a persistir a atribuição - semelhante ao caso de declaração de métodos com value receiver e pointer receiver

3) O uso de `errors.Is`se dá quando é necessário apenas saber se houve um valor igual ao valor de um determinado erro para definir qual ação a ser tomada, por exemplo: 
-  erro de timeout -> seria interessante um retry daqui a x período de tempo...

O uso de errors.As ou `erros.AsType` se dá quando queremos saber, por exemplo: o que levou esse erro a acontecer? Um cenário provável seria uma tentativa de criar um usuário: qual campo está inválido? Precisaríamos saber a partir do `erros.AsType`
