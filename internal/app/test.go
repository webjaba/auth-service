package app

import (
	"auth-service/internal/pkg/mock"

	"github.com/golang/mock/gomock"
)

type Setup struct {
	Service
	store *mock.MockIStore
	jwt   *mock.MockIJWT
}

func MustSetup(ctrl *gomock.Controller) *Setup {
	store := mock.NewMockIStore(ctrl)
	jwt := mock.NewMockIJWT(ctrl)
	return &Setup{
		Service: Service{
			store: store,
			jwt:   jwt,
		},
		store: store,
		jwt:   jwt,
	}
}
