package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/odmishien/mahjong-parser-go/rules"
)

var ErrInvalidKind = errors.New("invalid kind of hai")

func Parse(i string) (string, error) {
	// TODO: 順番がバラバラのものが入ってきてもパースできるようにする
	ss1 := strings.Split(i, "m")
	if len(ss1) < 2 {
		return "", ErrInvalidKind
	}
	m := ss1[0]
	ss2 := strings.Split(ss1[1], "s")
	if len(ss2) < 2 {
		return "", ErrInvalidKind
	}
	s := ss2[0]
	ss3 := strings.Split(ss2[1], "p")
	if len(ss3) < 2 {
		return "", ErrInvalidKind
	}
	p := ss3[0]
	ss4 := strings.Split(ss3[1], "w")
	if len(ss4) < 2 {
		return "", ErrInvalidKind
	}
	w := ss4[0]
	ss5 := strings.Split(ss4[1], "d")
	if len(ss5) < 2 {
		return "", ErrInvalidKind
	}
	d := ss5[0]

	var result string

	// manzu
	mr := rules.Man()
	ms, err := getUnicodes(m, mr)
	if err != nil {
		return "", err
	}
	for _, muc := range ms {
		result += fmt.Sprintf("%c ", muc)
	}

	// souzu
	sr := rules.Sou()
	ss, err := getUnicodes(s, sr)
	if err != nil {
		return "", err
	}
	for _, suc := range ss {
		result += fmt.Sprintf("%c ", suc)
	}

	// pinzu
	pr := rules.Pin()
	ps, err := getUnicodes(p, pr)
	if err != nil {
		return "", err
	}
	for _, puc := range ps {
		result += fmt.Sprintf("%c ", puc)
	}

	// wind
	wr := rules.Wind()
	ws, err := getUnicodes(w, wr)
	if err != nil {
		return "", err
	}
	for _, wuc := range ws {
		result += fmt.Sprintf("%c ", wuc)
	}

	// dragon
	dr := rules.Dragon()
	ds, err := getUnicodes(d, dr)
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
