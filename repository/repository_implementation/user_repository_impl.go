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
	sqlQuery := "INSERT INTO user (username, password) VALUES (?, ?)"
	_, err := userRepo.Db.ExecContext(ctx, sqlQuery, userpPayload.Username, userpPayload.Password)

	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}

func (userRepo *userRepositoryImpl) FetchByUsername(ctx context.Context, username string) (*entity.User, errs.Error) {
	sqlQuery := "SELECT user_id, username, password FROM user WHERE username = ?"
	rows, err := userRepo.Db.QueryContext(ctx, sqlQuery, username)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong on fetching username")
		}
		return &user, nil
	} else {
		return nil, errs.NewNotFoundError("user not found")
	}
}

func (userRepo *userRepositoryImpl) FetchById(ctx context.Context, id int) (*entity.User, errs.Error) {
	sqlQuery := "SELECT user_id, username, password FROM user WHERE user_id = ?"
	rows, err := userRepo.Db.QueryContext(ctx, sqlQuery, id)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong on fetching id")
		}
		return &user, nil
	} else {
		return nil, errs.NewNotFoundError("user not found")
	}
}
