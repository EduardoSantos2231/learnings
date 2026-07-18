# Respostas

1) Pelo que entendi, o que muda é que agora w armazena o valor de buf, que é um pointer que aponta para nil, ou seja, ainda que o buf seja nil, w não é nil

2) Type assertion é mais eficiente, mas garante menos flexibilidade quanto ao tipo, é preciso ter ao menos uma ideia do tipo de dado que se espera e pode ser necessário cobrir uma variedade de casos quando trabalhando com Type Switch e assertion. O Reflect embora menos eficiente se apresenta mais flexível, permitindo até que você explore campos de um struct que sequer conhece.

3) eu não sei, mas imagino que algumas vezes por descuido de inicialização é possível que um erro seja considerado um falso positivo já que para ser considerado nil é necessário um tipo e valor nulos 
