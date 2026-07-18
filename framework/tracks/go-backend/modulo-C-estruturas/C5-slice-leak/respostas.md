# Respostas

1) No caso do que fiz não é possível ver o valor removido porque torno ele nulo antes de fazer o reeslicing. O que o dequeue faz é mudar a janela do slice, quando olhamos a partir da capacidade do array que está por detrás ainda é possível expandir a janela e ver todos os elementos, mesmo aqueles que o slice originalmente tenta "esconder" movendo o ponteiro;

2) Os dois vazam, mas o pop normalmente irá sobrescrever com um append, enquanto o dequeue fica lá até o array inteiro ser descartado.  

3) o valor continua lá na array que está por debaixo dos panos, o que acontece é que o valor será adicionado ao final e o espaço "" irá desaparecer quando a array for recopiada para ser redimensionada 
