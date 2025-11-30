package app

import (
	"errors"
	"flag"
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
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	useStdin := fs.Bool("stdin", false, "read from stdin")
	if err := fs.Parse(args); err != nil {
		return err
	}
	rest := fs.Args()
	return validate.Handle(*useStdin, rest)
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
