package models

import (
	"math/big"
)

type Fib struct {
	Val *big.Int
}

func NewFib(val *big.Int) *Fib {
	return &Fib{
		Val: val,
	}
}
