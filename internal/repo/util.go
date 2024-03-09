package repo

import (
	"github.com/jackc/pgx/v5/pgtype"
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
