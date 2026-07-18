# L14 — Perguntas

**Q1 —** O que acontece se você redirecionar stdout e stderr para o mesmo arquivo com comandos separados (`> file 2> file`)? Por quê?

**Q2 —** Como `O_APPEND` difere de abrir o arquivo e fazer `Seek(0, io.SeekEnd)`? Qual é mais seguro para logs?

**Q3 —** Se um comando tem múltiplos redirecionamentos do mesmo tipo (ex: `cmd > a.txt > b.txt`), qual arquivo prevalece? Por quê?
