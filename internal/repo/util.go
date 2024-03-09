package repo

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func ConvertToNullStr(value string) (nullValue pgtype.Text) {
	if value != "" {
		nullValue = pgtype.Text{
			String: value,
			Valid:  true,
		}
	}
	return
}

func ConvertToNullBool(value *bool) (nullValue pgtype.Bool) {
	if value != nil {
		nullValue = pgtype.Bool{
			Bool:  *value,
			Valid: true,
		}
	}
	return
}

func ConvertToUUID(value uuid.UUID) (id pgtype.UUID) {
	if value != uuid.Nil {
		id = pgtype.UUID{
			Bytes: value,
			Valid: true,
		}
	}
	return
}

func ConvertToTimestamptz(value time.Time) (nullValue pgtype.Timestamptz) {
	if !value.IsZero() {
		nullValue = pgtype.Timestamptz{
			Time:  value,
			Valid: true,
		}
	}
	return
}
