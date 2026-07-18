# L3 — daemon

> Template: Design | Scaffolding: alto

Projete e implemente um processo daemon.

## Tarefas

1. **Design:** quais passos um processo precisa para virar daemon? (fork, setsid, chdir, fechar fds)
2. **Implemente `daemonize()`:** o processo se desacopla do terminal e continua rodando em background
3. Escreva arquivo PID em `/tmp/mydaemon.pid`
4. Trate sinais: SIGTERM limpa o PID file e sai

## Verificação

Explique como verificar que é um daemon: `ps aux | grep mydaemon` deve mostrar PPID=1 e sem TTY.
