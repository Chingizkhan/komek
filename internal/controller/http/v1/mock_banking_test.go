package v1

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"komek/internal/domain"
	mock_banking "komek/internal/service/banking/mock"
	"komek/pkg/httpserver"
	"komek/pkg/random"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAccount(t *testing.T) {
	acc := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	banking := mock_banking.NewMockBanking(ctrl)
	// build stubs
	banking.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(acc.ID)).
		Times(1).
		Return(acc, nil)

	// set http server and send request
	handler := NewHandler(&HandlerParams{
		Logger:            nil,
		Cfg:               nil,
		User:              nil,
		Banking:           banking,
		CookieSecret:      nil,
		OauthServerClient: nil,
	})
	httpServer := httpserver.New(
		chi.NewRouter(),
		httpserver.Port("8889"),
	)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/account/get")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	httpServer.
}

func randomAccount() domain.Account {
	return domain.Account{
		ID:        random.Int(1, 1000),
		Owner:     random.Owner(),
		Balance:   random.Money(),
		Currency:  random.Currency(),
		CreatedAt: time.Time{},
	}
}
