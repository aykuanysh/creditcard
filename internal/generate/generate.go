package generate

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/aykuanysh/creditcard/internal/validate"
)

func Handle(pick bool, args []string) error {
	if len(args) == 0 {
		return errors.New("no pattern provided")
	}
	pattern := args[0]

	starCount := strings.Count(pattern, "*")
	if starCount > 4 {
		return errors.New("too many asterisks")
	}

	if starCount > 0 {
		if !strings.HasSuffix(pattern, strings.Repeat("*", starCount)) {
			return errors.New("asterisks must be at the end")
		}
		return generateWithStars(pattern, starCount, pick)
	}

	if len(pattern) < 13 {
		return errors.New("card number too short (min 13 digits)")
	}

	if validate.IsValidLuhn(pattern) {
		fmt.Println(pattern)
		return nil
	} else {
		return errors.New("invalid card number")
	}
}

func generateWithStars(pattern string, starCount int, pick bool) error {
	prefix := pattern[:len(pattern)-starCount]
	limit := 1
	for i := 0; i < starCount; i++ {
		limit *= 10
	}

	results := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		suf := fmt.Sprintf("%0*d", starCount, i)
		candidate := prefix + suf
		if validate.IsValidLuhn(candidate) {
			results = append(results, candidate)
		}
	}

	if len(results) == 0 {
		return errors.New("no valid card numbers found")
	}

	sort.Strings(results)

	if pick {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(results[rand.Intn(len(results))])
	} else {
		for _, r := range results {
			fmt.Println(r)
		}
	}

	return nil
}
