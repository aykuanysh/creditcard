package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/aykuanysh/creditcard/internal/validate"
)

func Run() error {
	if len(os.Args) < 2 {
		return errors.New("usage: creditcard <validate|generate|information|issue> [options]")
	}

	cmd := os.Args[1]
	switch cmd {
	case "validate":
		return runValidate(os.Args[2:])
	case "generate":
		return runGenerate(os.Args[2:])
	case "information":
		return runInformation(os.Args[2:])
	case "issue":
		return runIssue(os.Args[2:])
	default:
		return fmt.Errorf("uknown command: %s", cmd)

	}
}

func runValidate(args []string) error {
	if len(args) < 1 {
		return errors.New("please provide a card number to validate")
	}

	valid := validate.IsValidLuhn(args[0])
	if valid {
		fmt.Println("Card is valid")
	} else {
		fmt.Println("Card is invalid")
	}
	return nil

}

func runGenerate(args []string) error {
	return fmt.Errorf("not implemented")
}

func runInformation(args []string) error {
	return fmt.Errorf("not implemented")
}

func runIssue(args []string) error {
	return fmt.Errorf("not implemented")
}
