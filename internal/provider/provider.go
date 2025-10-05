//go:generate wire .

package provider

import (
	"github.com/google/wire"
	"seal-x-polarhare-backend/internal/config"
	"seal-x-polarhare-backend/internal/repository"
	"seal-x-polarhare-backend/internal/service"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewProvider()
	if err != nil {
		panic(err)
	}
}

func Get() *Provider {
	return provider
}

// Provider 提供controller依赖的对象
type Provider struct {
	Config      *config.Config
	UserService service.UserService
}

var ServiceSet = wire.NewSet(
	service.UserServiceSet,
)

var RepositorySet = wire.NewSet(
	config.NewConfig,
	repository.NewUserRepository,
)

var AllProvider = wire.NewSet(
	ServiceSet,
	RepositorySet,
)
