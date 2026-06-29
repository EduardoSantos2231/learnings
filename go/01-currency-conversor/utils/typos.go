package utils

import (
	"fmt"
)

type Converter struct {
	rates map[string]float64
}

type Flags struct {
	Amount float64
	From   string
	To     string
}

func (c *Converter) Convert(from string, to string, amount float64) (float64, error) {
	_, okFrom := c.rates[from]
	_, okTo := c.rates[to]
	if !okFrom {
		return 0, fmt.Errorf("[ERROR]: You must provide a valid currency")

	}
	if !okTo {
		return 0, fmt.Errorf("[ERROR]: You must provide a valid currency %s, is not a valid currency", to)
	}

	result := amount * (c.rates[to] / c.rates[from])
	return result, nil
}

// GetSupportedCurrencies returns an slice of supported currencies declared in the converter
func (c *Converter) GetSupportedCurrencies() []string {
	result := make([]string, 0, 3)
	for key := range c.rates {
		result = append(result, key)
	}
	return result
}

// we choose to use USD as our reference
func NewConverter() *Converter {
	return &Converter{
		rates: map[string]float64{
			"BRL": 5.17,
			"USD": 1.0,
			"EUR": 0.87,
		},
	}
}
