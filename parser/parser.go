package parser

import (
	"errors"
	"strconv"
)

var ErrInvalidKind = errors.New("invalid kind of hai")
var ErrInvalidLiteral = errors.New("invalid literal of hai")

func GetUnicodes(s string, r map[int]int) ([]int, error) {
	ucs := make([]int, 0, len(s))
	for _, ss := range s {
		n, err := strconv.Atoi(string(ss))
		if err != nil {
			return nil, errors.New("invalid input")
		}
		uc, ok := r[n]
		if !ok {
			return nil, errors.New("out of range")
		}
		ucs = append(ucs, uc)
	}
	return ucs, nil
}
