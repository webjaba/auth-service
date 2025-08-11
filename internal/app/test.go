package app

import (
	"auth-service/internal/pkg/mock"

	"github.com/golang/mock/gomock"
)

type Setup struct {
	Service
	store      *mock.MockIStore
	jwtManager *mock.MockIJWTManager
}

func MustSetup(ctrl *gomock.Controller) *Setup {
	store := mock.NewMockIStore(ctrl)
	jwtManager := mock.NewMockIJWTManager(ctrl)
	return &Setup{
		Service: Service{
			store:      store,
			jwtManager: jwtManager,
		},
		store:      store,
		jwtManager: jwtManager,
	}
}
