package repo

import (
	"context"

	"github.com/huyhoang1990/beego-investigation/entity"
)

type UserRepo interface {
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
	InsertOne(ctx context.Context, newUser *entity.User) error
	FindById(ctx context.Context, id string) (*entity.User, error)
}
