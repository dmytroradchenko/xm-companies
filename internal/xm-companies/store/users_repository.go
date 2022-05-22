//go:generate mockery --name UsersRepository --output mocks --case underscore
package store

import (
	"context"
	"xm-companies/internal/xm-companies/model"
)

type UsersRepository interface {
	Create(ctx context.Context, user *model.User) (string, error)
	Find(ctx context.Context, userName string) (*model.User, error)
}
