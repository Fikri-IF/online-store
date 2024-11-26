package repositoryimplementation

import (
	"context"
	"database/sql"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/model"
	"online-store-golang/repository"
)

type CategoryRepositoryImpl struct {
	Db *sql.DB
}

func NewCategoryRepository(Db *sql.DB) repository.CategoryRepository {
	return &CategoryRepositoryImpl{
		Db: Db,
	}
}

func (cr *CategoryRepositoryImpl) Create(ctx context.Context, category *entity.Category) errs.Error {
	sqlQuery := "insert into category (name) values (?)"
	result, err := cr.Db.ExecContext(ctx, sqlQuery, category.Name)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return errs.NewInternalServerError("error from getting last insert Id")
	}

	category.CategoryId = int(id)
	return nil
}

func (cr *CategoryRepositoryImpl) FindById(ctx context.Context, id int) (*entity.Category, errs.Error) {
	sqlQuery := "select category_id, name from category where category_id = ?"
	category := entity.Category{}
	err := cr.Db.QueryRowContext(ctx, sqlQuery, id).Scan(
		&category.CategoryId,
		&category.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("category not found")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &category, nil
}

func (cr *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]model.CategoryResponse, errs.Error) {
	sqlQuery := "select * from category"
	rows, err := cr.Db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	var categories []model.CategoryResponse
	for rows.Next() {
		category := model.CategoryResponse{}
		err := rows.Scan(
			&category.CategoryId,
			&category.CategoryName,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (cr *CategoryRepositoryImpl) Delete(ctx context.Context, id int) errs.Error {
	sqlQuery := "delete from category where category_id = ?"
	_, err := cr.Db.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}
