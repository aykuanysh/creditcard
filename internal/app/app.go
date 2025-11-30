package app

import (
	"errors"
	"fmt"
	"os"
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
		return fmt.Errorf("Uknown command: %s", cmd)

	}
}

func runValidate(args []string) error {
	return fmt.Errorf("not implemented")
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
