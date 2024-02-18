package random

import (
	"math/rand"
)

var (
	currency = []string{"USD", "KZT", "EUR", "KGS"}
)

func Owner() string {
	return String(6)
}

func Money() int64 {
	return Int(0, 1000)
}

func Currency() string {
	n := len(currency)
	return currency[rand.Intn(n)]
}
