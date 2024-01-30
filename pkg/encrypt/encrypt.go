package encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
)

// https://www.alexedwards.net/blog/working-with-cookies-in-go

func Encrypt(val string, secretKey []byte) string {
	// Calculate a HMAC signature of the cookie name and value, using SHA256 and
	// a secret key (which we will create in a moment).
	//value := base64.URLEncoding.EncodeToString([]byte(val))
	//log.Println("val", val)
	//log.Println("value", value)

	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(val))
	signature := mac.Sum(nil)

	// Call our Write() helper to base64-encode the new cookie value and write
	// the cookie.
	return string(signature)
}
