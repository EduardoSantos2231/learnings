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
		fmt.Fprintln(osStderr, "  done <challenge>                    Marcar desafio como concluído")
		fmt.Fprintln(osStderr, "  review <challenge> <1d|3d|7d|30d> --pass|--fail  Registrar resultado de revisão")
		fmt.Fprintln(osStderr, "  add-error <challenge> <categoria> \"<descrição>\"  Registrar erro")
		fmt.Fprintln(osStderr, "  check-recurrence <categoria>        Verificar se categoria de erro já apareceu")
		fmt.Fprintln(osStderr, "  start <track>                       Ativar uma track de estudo")
		fmt.Fprintln(osStderr, "  render-roadmap                      Gerar roadmap.md a partir do JSON")
		return
	}

	switch args[0] {
	case "status":
		statusCmd()

	case "done":
		if len(args) < 2 {
			fmt.Fprintln(osStderr, "uso: tracking done <challenge>")
			return
		}
		doneCmd(args[1])

	case "review":
		if len(args) < 4 {
			fmt.Fprintln(osStderr, "uso: tracking review <challenge> <1d|3d|7d|30d> --pass|--fail")
			return
		}
		var result string
		switch args[3] {
		case "--pass":
			result = "passed"
		case "--fail":
			result = "failed"
		default:
			fmt.Fprintln(osStderr, "resultado deve ser --pass ou --fail")
			return
		}
		reviewCmd(args[1], args[2], result)

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
