package cli

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/example-nucleus-molecule/ci-calc/internal/cicalc"

	"github.com/manifoldco/promptui"
)

func Start() {
	fmt.Println(asciiTitle)
	fmt.Println()
	fmt.Println()
	principal := getPrincipal()
	rate := getRate()
	term := getTerm()
	freq := getFreq()
	fmt.Printf("Final Balance: $%d", cicalc.Periodic(principal, rate, term, freq))
}

func getPrincipal() float64 {
	var principal float64
	prompt := promptui.Prompt{
		Label:    "Enter Principal Amount ($)",
		Validate: validateFloat64(&principal),
	}

	_, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return principal
}

func getRate() float64 {
	var rate float64
	prompt := promptui.Prompt{
		Label:    "Enter Interest Rate (%)",
		Validate: validateFloat64(&rate),
	}

	_, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return rate
}

func getTerm() float64 {
	var term float64
	prompt := promptui.Prompt{
		Label:    "Enter Deposit Term (Years)",
		Validate: validateFloat64(&term),
	}

	_, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return term
}

func getFreq() cicalc.Frequency {
	prompt := promptui.Select{
		Label: "Select Interest Frequency",
		Items: []string{"Monthly", "Quarterly", "Annual", "Paid at Maturity"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	switch result {
	case "Monthly":
		return cicalc.Monthly
	case "Quarterly":
		return cicalc.Quarterly
	case "Annual":
		return cicalc.Annual
	case "Paid at Maturity":
		return cicalc.AtMaturity
	}

	log.Fatal("Something went horribly wrong")
	return 0
}

func validateFloat64(parsedFloat *float64) promptui.ValidateFunc {
	var err error
	return func(input string) error {
		*parsedFloat, err = strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Number must contain only digits, `-`, and `.`")
		}
		return nil
	}
}
