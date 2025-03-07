package null_value

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"math/big"
	"time"
)

func String(value string) (nullValue pgtype.Text) {
	if value != "" {
		nullValue = pgtype.Text{
			String: value,
			Valid:  true,
		}
	}
	return
}

func Bool(value *bool) (nullValue pgtype.Bool) {
	if value != nil {
		nullValue = pgtype.Bool{
			Bool:  *value,
			Valid: true,
		}
	}
	return
}

func UUID(value uuid.UUID) (id pgtype.UUID) {
	if value != uuid.Nil {
		id = pgtype.UUID{
			Bytes: value,
			Valid: true,
		}
	}
	return
}

func Timestamp(value time.Time) (nullValue pgtype.Timestamptz) {
	if !value.IsZero() {
		nullValue = pgtype.Timestamptz{
			Time:  value,
			Valid: true,
		}
	}
	return
}

func Number(value *int) (nullValue pgtype.Numeric) {
	if value != nil {
		nullValue = pgtype.Numeric{
			Int:   big.NewInt(int64(*value)),
			Valid: true,
		}
	}
	return
}

func Int64(value *int64) (nullValue pgtype.Numeric) {
	if value != nil {
		nullValue = pgtype.Numeric{
			Int:   big.NewInt(*value),
			Valid: true,
		}
	}
	return
}
