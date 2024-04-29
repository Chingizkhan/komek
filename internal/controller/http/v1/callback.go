package v1

import (
	"golang.org/x/oauth2"
	"net/http"
)

//func (h *Handler) callback(w http.ResponseWriter, r *http.Request) {
//	tokens, introspectResponse, err := h.sso.ProcessCallback(r)
//	if err != nil {
//		return
//	}
//
//	log.Printf("tokens: %#v", tokens)
//	log.Printf("introspectReesponse: %#v", introspectResponse)
//}

func processCookies(w http.ResponseWriter, token *oauth2.Token) {
	cookieAccess := &http.Cookie{
		Name:     "Access-Token",
		Value:    token.AccessToken,
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	}
	cookieRefresh := &http.Cookie{
		Name:     "Refresh-Token",
		Value:    token.RefreshToken,
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookieAccess)
	http.SetCookie(w, cookieRefresh)
}
