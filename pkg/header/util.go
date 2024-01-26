package header

import (
	"net/http"
	"strings"
)

func Get(r *http.Request, key, cutText string) string {
	return strings.TrimSpace(strings.Replace(r.Header.Get(key), cutText, "", 1))
}
