package repositoryimplementation

import (
	"context"
	"database/sql"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/model"
	"online-store-golang/repository"
)

type ProductRepositoryImpl struct {
	Db *sql.DB
}

func NewProductRepository(Db *sql.DB) repository.ProductRepository {
	return &ProductRepositoryImpl{
		Db: Db,
	}
}

func (productRepo *ProductRepositoryImpl) Create(ctx context.Context, product *entity.Product) errs.Error {
	sqlQuery := `insert into product (name, price, category, stock)
		values (?, ?, ?, ?)`
	result, err := productRepo.Db.ExecContext(ctx, sqlQuery, product.ProductName, product.Price, product.Category, product.Stock)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return errs.NewInternalServerError("error from getting last insert Id")
	}

	product.ProductId = int(id)

	return nil

}

func (productRepo *ProductRepositoryImpl) FindAll(ctx context.Context) ([]model.GetProductResponse, errs.Error) {
	sqlQuery := `select  p.product_id , p.name , p.price , c.name , p.stock 
		from product p 
		inner join category c 
		on c.category_id = p.category`
	rows, err := productRepo.Db.QueryContext(ctx, sqlQuery)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer rows.Close()

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
func (productRepo *ProductRepositoryImpl) FindByCategory(ctx context.Context, id int) ([]model.GetProductResponse, errs.Error) {
	sqlQuery := `select  p.product_id , p.name , p.price , c.name , p.stock 
		from product p 
		inner join category c 
		on c.category_id = p.category
		where p.category = ?`
	rows, err := productRepo.Db.QueryContext(ctx, sqlQuery, id)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer rows.Close()

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

func (productRepo *ProductRepositoryImpl) FindById(ctx context.Context, id int) (*model.GetProductResponse, errs.Error) {
	sqlQuery := `select  p.product_id , p.name , p.price , c.name , p.stock 
		from product p 
		inner join category c 
		on c.category_id = p.category
		where p.product_id = ?`
	product := model.GetProductResponse{}
	err := productRepo.Db.QueryRowContext(ctx, sqlQuery, id).Scan(
		&product.ProductId,
		&product.ProductName,
		&product.Price,
		&product.Category,
		&product.Stock,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("product not found")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &product, nil
}

func (productRepo *ProductRepositoryImpl) Delete(ctx context.Context, id int) errs.Error {
	sqlQuery := "delete from product where product_id = ?"
	_, err := productRepo.Db.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}
