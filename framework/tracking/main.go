package main

import (
	"fmt"
	"os"
)

var osStderr = os.Stderr

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(osStderr, "tracking — CLI de gerenciamento de estudos")
		fmt.Fprintln(osStderr, "")
		fmt.Fprintln(osStderr, "Comandos:")
		fmt.Fprintln(osStderr, "  status                              Posição atual, próximo desafio, revisões")
		fmt.Fprintln(osStderr, "  session                             Escolher a sessão de hoje")
		fmt.Fprintln(osStderr, "  finish <id> --pass|--fail            Registrar a sessão escolhida")
		fmt.Fprintln(osStderr, "  rebaseline                          Arquivar revisões antigas da track ativa")
		fmt.Fprintln(osStderr, "  done <challenge>                    Marcar desafio como concluído")
		fmt.Fprintln(osStderr, "  add-error <challenge> <categoria> \"<descrição>\"  Registrar erro")
		fmt.Fprintln(osStderr, "  check-recurrence <categoria>        Verificar se categoria de erro já apareceu")
		fmt.Fprintln(osStderr, "  start <track>                       Ativar uma track de estudo")
		fmt.Fprintln(osStderr, "  render-roadmap                      Gerar roadmap.md a partir do JSON")
		return
	}

	switch args[0] {
	case "status":
		statusCmd()

	case "session":
		sessionCmd()

	case "finish":
		if len(args) < 3 {
			fmt.Fprintln(osStderr, "uso: tracking finish <id> --pass|--fail")
			return
		}
		result := ""
		switch args[2] {
		case "--pass":
			result = "passed"
		case "--fail":
			result = "failed"
		default:
			fmt.Fprintln(osStderr, "resultado deve ser --pass ou --fail")
			return
		}
		finishCmd(args[1], result)

	case "rebaseline":
		rebaselineCmd()

	case "done":
		if len(args) < 2 {
			fmt.Fprintln(osStderr, "uso: tracking done <challenge>")
			return
		}
		finishCmd(args[1], "passed")

	case "add-error":
		if len(args) < 4 {
			fmt.Fprintln(osStderr, "uso: tracking add-error <challenge> <categoria> \"<descrição>\"")
			return
		}
		addErrorCmd(args[1], args[2], args[3])

	case "check-recurrence":
		if len(args) < 2 {
			fmt.Fprintln(osStderr, "uso: tracking check-recurrence <categoria>")
			return
		}
		checkRecurrenceCmd(args[1])

	case "start":
		if len(args) < 2 {
			fmt.Fprintln(osStderr, "uso: tracking start <track>")
			return
		}
		startCmd(args[1])

	case "render-roadmap":
		renderRoadmapCmd()

	case "book":
		bookCmd(args[1:])

	case "--init-all":
		initAllCmd()

	default:
		fmt.Fprintf(osStderr, "comando desconhecido: %s\n", args[0])
	}
}
