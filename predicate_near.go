package main

import (
	"fmt"
	"io"
	"strconv"
)

type FindNear struct {
	n           float64
	firstNumber float64
	showDecimal bool
}

func NewFindNear(rawN string) (f *FindNear, err error) {
	n, err := strconv.ParseFloat(rawN, 64)
	if err != nil {
		return nil, err
	}
	return &FindNear{
		n:           n,
		firstNumber: 3.0,
		showDecimal: false,
	}, nil
}

func (f *FindNear) Init(firstNumber float64, showDecimal bool) {
	f.firstNumber = firstNumber
	f.showDecimal = showDecimal
}

func (f *FindNear) Do(w io.Writer) error {
	l := f.firstNumber
	for l < f.n {
		l *= goldenRatio
	}
	s := l / goldenRatio

	smaller, larger := s, l
	if f.showDecimal {
		fmt.Fprintln(w, smaller, larger)
	} else {
		fmt.Fprintln(w, round(smaller), round(larger))
	}
	return nil
}
