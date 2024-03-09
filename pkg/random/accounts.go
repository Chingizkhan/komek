package random

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"math/rand"
)

var (
	currency = []string{"USD", "KZT", "EUR", "KGS"}
)

func Owner() pgtype.UUID {
	return pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
}

func Money() int64 {
	return Int(0, 1000)
}

func Currency() string {
	n := len(currency)
	return currency[rand.Intn(n)]
}
