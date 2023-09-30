package interfaces

import (
	"context"
	"gorm/core"
)

type UserStore interface {
	Create(context.Context, *core.User) error
	Update(context.Context, *core.User) error
	Delete(context.Context, *core.User) error
	Find(context.Context, int64) (*core.User, error)
	List(context.Context) ([]*core.User, error)
	FindLogin(context.Context, string) (*core.User, error)
	FindActive(context.Context) ([]*core.User, error)
	Count(context.Context) (int64, error)
}
