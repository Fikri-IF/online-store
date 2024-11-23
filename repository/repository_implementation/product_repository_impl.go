package repositoryimplementation

import (
	"context"
	"database/sql"
	"online-store-golang/errs"
	"online-store-golang/model"
	"online-store-golang/repository"
)

type productRepositoryImpl struct {
	Db *sql.DB
}

func NewProductRepository(Db *sql.DB) repository.ProductRepository {
	return &productRepositoryImpl{
		Db: Db,
	}
}

func (productRepo *productRepositoryImpl) FindAll(ctx context.Context) ([]model.GetProductResponse, errs.Error) {
	sqlQuery := `select  p.product_id , p.name , p.price , c.name , p.stock 
		from product p 
		inner join category c 
		on c.category_id = p.category`
	rows, err := productRepo.Db.QueryContext(ctx, sqlQuery)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	var products []model.GetProductResponse
	for rows.Next() {
		product := model.GetProductResponse{}
		err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.Price,
			&product.Category,
			&product.Stock,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		products = append(products, product)
	}
	return products, nil
}
func (productRepo *productRepositoryImpl) FindByCategory(ctx context.Context, id int) ([]model.GetProductResponse, errs.Error) {
	sqlQuery := `select  p.product_id , p.name , p.price , c.name , p.stock 
		from product p 
		inner join category c 
		on c.category_id = p.category
		where p.category = ?`
	rows, err := productRepo.Db.QueryContext(ctx, sqlQuery, id)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	var products []model.GetProductResponse
	for rows.Next() {
		product := model.GetProductResponse{}
		err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.Price,
			&product.Category,
			&product.Stock,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		products = append(products, product)
	}
	return products, nil
}
