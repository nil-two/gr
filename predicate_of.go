package main

import (
	"fmt"
	"io"
	"strconv"
)

type FindOf struct {
	n           float64
	firstNumber float64
	showDecimal bool
}

func NewFindOf(rawN string) (f *FindOf, err error) {
	n, err := strconv.ParseFloat(rawN, 64)
	if err != nil {
		return nil, err
	}
	return &FindOf{
		n:           n,
		firstNumber: 3.0,
		showDecimal: false,
	}, nil
}

func (f *FindOf) Init(firstNumber float64, showDecimal bool) {
	f.firstNumber = firstNumber
	f.showDecimal = showDecimal
}

func (f *FindOf) Do(w io.Writer) error {
	smaller, larger := f.n/goldenRatio, f.n*goldenRatio
	if f.showDecimal {
		fmt.Fprintln(w, smaller, larger)
	} else {
		fmt.Fprintln(w, round(smaller), round(larger))
	}
	return nil
}
