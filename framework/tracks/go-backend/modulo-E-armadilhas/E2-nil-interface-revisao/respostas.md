# Respostas

1) Imaginando nesse cenário 
```go

type signer interface{
  sign() int
}

var tool signer 

```

A variável `tool` deverá futuramente conter um pointer que aponte para alguém que implemente o método signer

```go
type fakeSigner struct {
  name string
}

func (fs *fakeSigner) sign() int{
  return 1
}

var fakeSigner *fakeSigner

tool = fakeSigner
```


O erro acontece porque tool acredita que recebeu um valor que possui uma instância de um `signer`, no entanto, embora o seu tipo não seja nulo, seu value é;

Ao tentar acessar, por exemplo, o método sign, será impossível, pois o ponteiro é `nil` 

No primeiro momento:

a) `var w io.Writer`
- type nil
- value nil
b) w = buff
- type pointer para bytes.buffer 
- value nil 

c) w.Write
- type pointer para bytes.buffer 
- value nil

2) Fiz uma implementação que escrevi com auxilio do gemini nós primeiros garantimos que todas as implementações devem ser de tipos que podem ser do tipo `nil` e evitar panicar (map, chan, func, pointer); Ele iria panicar caso chamasse o método `isNil` em um valor que não pode assumir o valor de nil (strings, floats, ints, bool, structs)

3) O primeiro caso que consigo pensar é quando temos um erro que pode ser um falso positivo e ativar os errors handlers:


```go
func isValidId(id string) error{
  var myErr *EspecificError = nil
      //tratamento
      if id==""{
          myErr = &EspecificError{...}
      }
  return myErr
}
```

Note que o tipo aqui não seria `nil`, como consequência ele dispararia um erro. O type é um pointer para um struct e o value é nil; Para uma interface ser considerada `nil` precisamos dos dois valores como `nil`;

O segundo caso seria possível imaginar da seguinte forma:

```go
var buffer *bytes.Buffer

func writeSomething(b io.Writer) //escreve em um buffer 

writeSomething(buffer) //panic
```



Buffer possui o tipo que atende à assinatura da função mas não há uma instância de um buffer

