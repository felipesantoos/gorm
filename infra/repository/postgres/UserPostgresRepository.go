package postgres

import (
	"context"
	"gorm/core"
	"gorm/infra/repository/database"
	"gorm/infra/repository/postgres/queryMapper"
)

type UserPostgresRepository struct {
	db database.DB
}

func (repository *UserPostgresRepository) Find(ctx context.Context, id int64) (*core.User, error) {
	out := &core.User{ID: id}
	err := repository.db.View(func(queryer database.Queryer, binder database.Binder) error {
		params := queryMapper.ToParams(out)
		query, args, err := binder.BindNamed(queryMapper.UserQueryKey, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args)
		return queryMapper.ScanRow(row, out)
	})
	return out, err
}

func NewUserPostgresRepository() *UserPostgresRepository {
	return &UserPostgresRepository{}
}
