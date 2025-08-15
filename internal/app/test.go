package app

import (
	"auth-service/internal/pkg/mock"

	"github.com/golang/mock/gomock"
)

type Setup struct {
	Service
	store      *mock.MockIStore
	jwtManager *mock.MockIJWT
}

func MustSetup(ctrl *gomock.Controller) *Setup {
	store := mock.NewMockIStore(ctrl)
	jwtManager := mock.NewMockIJWT(ctrl)
	return &Setup{
		Service: Service{
			store:      store,
			jwtManager: jwtManager,
		},
		store:      store,
		jwtManager: jwtManager,
	}
}
