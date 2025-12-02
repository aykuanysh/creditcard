package info

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aykuanysh/creditcard/internal/validate"
)

type Brand struct {
	Name   string
	Prefix string
}

type Issuer struct {
	Name   string
	Prefix string
}

func CardInfo(brandsFile, issuersFile string, useStdin bool, args []string) error {
	if brandsFile == "" || issuersFile == "" {
		return errors.New("brands and issuers files must be provided")
	}

	brands, err := LoadBrands(brandsFile)
	if err != nil {
		return err
	}

	issuers, err := LoadIssuers(issuersFile)
	if err != nil {
		return err
	}

	numbers := args
	if useStdin {
		stdinNumbers, err := readStdin()
		if err != nil {
			return err
		}
		numbers = append(numbers, stdinNumbers...)
	}

	if len(numbers) == 0 {
		return errors.New("no card numbers provided")
	}

	for _, num := range numbers {
		processNumber(num, brands, issuers)
	}

	return nil
}

func readStdin() ([]string, error) {
	var numbers []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			parts := strings.Fields(line)
			numbers = append(numbers, parts...)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

func LoadBrands(path string) ([]Brand, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var brands []Brand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		brands = append(brands, Brand{
			Name:   strings.TrimSpace(parts[0]),
			Prefix: strings.TrimSpace(parts[1]),
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return brands, nil
}

func LoadIssuers(path string) ([]Issuer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var issuers []Issuer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		issuers = append(issuers, Issuer{
			Name:   strings.TrimSpace(parts[0]),
			Prefix: strings.TrimSpace(parts[1]),
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return issuers, nil
}

func processNumber(number string, brands []Brand, issuers []Issuer) {
	valid := validate.IsValidLuhn(number)
	brand := findBrand(number, brands)
	issuer := findIssuer(number, issuers)

	fmt.Println(number)
	fmt.Printf("Correct: %s\n", boolToYesNo(valid))
	fmt.Printf("Card Brand: %s\n", brand)
	fmt.Printf("Card Issuer: %s\n\n", issuer)
}

func findBrand(number string, brands []Brand) string {
	for _, b := range brands {
		if strings.HasPrefix(number, b.Prefix) {
			return b.Name
		}
	}
	return "-"
}

func findIssuer(number string, issuers []Issuer) string {
	for _, i := range issuers {
		if strings.HasPrefix(number, i.Prefix) {
			return i.Name
		}
	}
	return "-"
}

func boolToYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
