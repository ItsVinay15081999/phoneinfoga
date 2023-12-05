package test

import "github.com/sundowndev/phoneinfoga/v2/lib/number"

func NewFakeUSNumber() *number.Number {
	n, _ := number.NewNumber("+91.6393000802")
	return n
}
