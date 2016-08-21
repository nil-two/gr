package main

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var rangeSyntax = regexp.MustCompile(`(\d*\.?\d*)?-(\d*\.?\d*)?`)

func parseRange(rawRange string) (first float64, last float64, err error) {
	if !rangeSyntax.MatchString(rawRange) {
		return 0, 0, fmt.Errorf("")
	}

	m := rangeSyntax.FindStringSubmatch(rawRange)
	rawFirst, rawLast := m[1], m[2]

	first, err = strconv.ParseFloat(rawFirst, 64)
	if rawFirst != "" && err != nil {
		return 0, 0, err
	}
	last, err = strconv.ParseFloat(rawLast, 64)
	if rawLast != "" && err != nil {
		return 0, 0, err
	}

	return first, last, nil
}

type FindRange struct {
	isEndless   bool
	first       float64
	last        float64
	firstNumber float64
	showDecimal bool
}

func NewFindRange(rawRange string) (f *FindRange, err error) {
	first, last, err := parseRange(rawRange)
	if err != nil {
		return nil, err
	}
	return &FindRange{
		isEndless:   last == 0.0,
		first:       first,
		last:        last,
		firstNumber: 3.0,
		showDecimal: false,
	}, nil
}

func (f *FindRange) Init(firstNumber float64, showDecimal bool) {
	f.firstNumber = firstNumber
	f.showDecimal = showDecimal
}

func (f *FindRange) Do(w io.Writer) error {
	n := f.firstNumber
	for f.isEndless || n <= f.last {
		if n >= f.first {
			if f.showDecimal {
				fmt.Fprintln(w, n)
			} else {
				fmt.Fprintln(w, round(n))
			}
		}
		n *= goldenRatio
	}
	return nil
}
