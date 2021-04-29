package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/odmishien/mahjong-parser-go/rules"
)

var ErrInvalidKind = errors.New("invalid kind of hai")
var ErrInvalidLiteral = errors.New("invalid literal of hai")

func Parse(i string) (string, error) {
	m, err := splitByKind(i)
	if err != nil {
		return "", err
	}

	var result string

	// manzu
	mr := rules.Man()
	ms, err := getUnicodes(m["m"], mr)
	if err != nil {
		return "", err
	}
	for _, muc := range ms {
		result += fmt.Sprintf("%c ", muc)
	}

	// souzu
	sr := rules.Sou()
	ss, err := getUnicodes(m["s"], sr)
	if err != nil {
		return "", err
	}
	for _, suc := range ss {
		result += fmt.Sprintf("%c ", suc)
	}

	// pinzu
	pr := rules.Pin()
	ps, err := getUnicodes(m["p"], pr)
	if err != nil {
		return "", err
	}
	for _, puc := range ps {
		result += fmt.Sprintf("%c ", puc)
	}

	// wind
	wr := rules.Wind()
	ws, err := getUnicodes(m["w"], wr)
	if err != nil {
		return "", err
	}
	for _, wuc := range ws {
		result += fmt.Sprintf("%c ", wuc)
	}

	// dragon
	dr := rules.Dragon()
	ds, err := getUnicodes(m["d"], dr)
	if err != nil {
		return "", err
	}
	for _, duc := range ds {
		result += fmt.Sprintf("%c ", duc)
	}

	return result, nil
}

func getUnicodes(s string, r map[int]int) ([]int, error) {
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

func splitByKind(s string) (map[string]string, error) {
	kind := []string{"m", "s", "p", "w", "d"}
	res := make(map[string]string)

	for _, k := range kind {
		r := regexp.MustCompile(fmt.Sprintf(`([1-9]+)(%s{1})`, k))
		g := r.FindStringSubmatch(s)
		if len(g) < 1 {
			return nil, ErrInvalidLiteral
		}
		res[k] = g[1]
	}
	return res, nil
}
