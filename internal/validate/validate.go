package validate

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Handle(stdin bool, args []string) error {
	if stdin {
		return validateFromStdin()
	}
	if len(args) == 0 {
		return errors.New("no card numbers provided")
	}
	for _, s := range args {
		if len(s) < 13 {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			return fmt.Errorf("card too short")
		}
		if IsValidLuhn(s) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			return fmt.Errorf("invalid card")
		}
	}
	return nil
}

func validateFromStdin() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 13 {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			return fmt.Errorf("card too short")
		}
		if IsValidLuhn(line) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			return fmt.Errorf("invalid card")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
