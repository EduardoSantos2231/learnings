package main

import (
	"01-currency-conversor/utils"
	"errors"
	"flag"
	"fmt"
	"os"
)

var errEmptyCurrency = errors.New("currency must contain a valid value")
var errCurrencyLessThanZero = errors.New("currency must be greater than zero ")

func main() {
	conversor := utils.NewConverter()
	flags := defineFlags()
	flagsError := validateFlags(flags)
	if flagsError != nil {
		if errors.Is(flagsError, errEmptyCurrency) {
			fmt.Fprintf(os.Stderr, "CURRENCY EMPTY\n")
		} else if errors.Is(flagsError, errCurrencyLessThanZero) {
			fmt.Fprintf(os.Stderr, "CURRENCY IS NEGATIVE\n")
		}
		os.Exit(2)
	}

	output, outputErr := conversor.Convert(flags.From, flags.To, flags.Amount)
	if outputErr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", outputErr)
		os.Exit(1)
	}
	fmt.Printf("Converting from [%s] to [%s] the amount [%.2f]\n", flags.From, flags.To, flags.Amount)
	fmt.Printf("\n--> %.2f\n", output)
}

func defineFlags() *utils.Flags {
	var flags utils.Flags
	flag.Float64Var(&flags.Amount, "amount", 0.0, "specify the currency Amount")
	flag.StringVar(&flags.From, "from", "", "specify the starting currency")
	flag.StringVar(&flags.To, "to", "", "specify to which currency should we convert to")
	flag.Parse()
	return &flags
}

func validateFlags(f *utils.Flags) error {
	if f.Amount <= 0 {
		return fmt.Errorf("%w", errCurrencyLessThanZero)
	}
	if f.From == "" || f.To == "" {
		return fmt.Errorf("%w", errEmptyCurrency)
	}

	return nil
}
