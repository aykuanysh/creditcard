package issue

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

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

func CardIssue(brandsFile, issuersFile, brandName, issuerName string) error {
	if brandsFile == "" || issuersFile == "" {
		return errors.New("brands and issuers files must be provided")
	}
	if brandName == "" {
		return errors.New("--brand required")
	}
	if issuerName == "" {
		return errors.New("--issuer required")
	}

	brands, err := LoadBrands(brandsFile)
	if err != nil {
		return err
	}

	issuers, err := LoadIssuers(issuersFile)
	if err != nil {
		return err
	}

	brandPrefix := findBrandPrefix(brandName, brands)
	if brandPrefix == "" {
		return fmt.Errorf("unknown brand: %s", brandName)
	}

	issuerPrefix := findIssuerPrefix(issuerName, issuers)
	if issuerPrefix == "" {
		return fmt.Errorf("unknown issuer: %s", issuerName)
	}

	if !strings.HasPrefix(issuerPrefix, brandPrefix) {
		return fmt.Errorf("issuer %q does not belong to brand %q", issuerName, brandName)
	}

	card, err := generateCard(issuerPrefix)
	if err != nil {
		return err
	}

	fmt.Println(card)
	return nil
}


func LoadBrands(path string) ([]Brand, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var list []Brand
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		list = append(list, Brand{
			Name:   strings.TrimSpace(parts[0]),
			Prefix: strings.TrimSpace(parts[1]),
		})
	}
	return list, sc.Err()
}

func LoadIssuers(path string) ([]Issuer, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var list []Issuer
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		list = append(list, Issuer{
			Name:   strings.TrimSpace(parts[0]),
			Prefix: strings.TrimSpace(parts[1]),
		})
	}
	return list, sc.Err()
}


func findBrandPrefix(name string, brands []Brand) string {
	for _, b := range brands {
		if strings.EqualFold(b.Name, name) {
			return b.Prefix
		}
	}
	return ""
}

func findIssuerPrefix(name string, issuers []Issuer) string {
	for _, i := range issuers {
		if strings.EqualFold(i.Name, name) {
			return i.Prefix
		}
	}
	return ""
}


func generateCard(prefix string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	length := 16
	digits := []byte(prefix)

	for len(digits) < length-1 {
		digits = append(digits, byte('0'+rand.Intn(10)))
	}

	check := luhnChecksum(digits)
	digits = append(digits, byte('0'+check))

	card := string(digits)

	if !validate.IsValidLuhn(card) {
		return "", errors.New("internal error: generated invalid Luhn card")
	}

	return card, nil
}

func luhnChecksum(num []byte) int {
	sum := 0
	double := true

	for i := len(num) - 1; i >= 0; i-- {
		d := int(num[i] - '0')
		if double {
			d = d * 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}

	return (10 - (sum % 10)) % 10
}
