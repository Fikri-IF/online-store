package repositoryimplementation

import (
	"context"
	"database/sql"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/repository"
)

type userRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(Db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{
		Db: Db,
	}
}

func (userRepo *userRepositoryImpl) Create(ctx context.Context, userpPayload *entity.User) errs.Error {
	sql := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := userRepo.Db.ExecContext(ctx, sql, userpPayload.Username, userpPayload.Password)

	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}
