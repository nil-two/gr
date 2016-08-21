package main

import (
	"io"
)

type Predicate interface {
	Init(firstNumber float64, showDecimal bool)
	Do(w io.Writer) error
}
