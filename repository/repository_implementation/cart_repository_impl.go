package repositoryimplementation

import (
	"context"
	"database/sql"
	"online-store-golang/entity"
	"online-store-golang/errs"
	"online-store-golang/model"
	"online-store-golang/repository"
)

type CartRepositoryImpl struct {
	Db *sql.DB
}

func NewCartRepository(Db *sql.DB) repository.CartRepository {
	return &CartRepositoryImpl{
		Db: Db,
	}
}

func (cartRepo *CartRepositoryImpl) AddItem(ctx context.Context, cartItemPayload *entity.CartItem) errs.Error {
	sqlQuery := `insert into cart (user_id, product_id, quantity)
		values (?, ?, ?)
	`
	_, err := cartRepo.Db.ExecContext(
		ctx,
		sqlQuery,
		cartItemPayload.UserId,
		cartItemPayload.ProductId,
		cartItemPayload.Quantity,
	)

	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}

func (cartRepo *CartRepositoryImpl) GetUserCart(ctx context.Context, userId int) ([]model.UserCartResponse, errs.Error) {
	sqlQuery := `select c.product_id , p.name , c.quantity, p.price
		from cart c 
		left join user u on c.user_id = u.user_id 
		left join product p on c.product_id = p.product_id 
		where u.user_id = ?`

	rows, err := cartRepo.Db.QueryContext(ctx, sqlQuery, userId)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var items []model.UserCartResponse
	for rows.Next() {
		item := model.UserCartResponse{}
		err := rows.Scan(&item.ProductId, &item.ProductName, &item.Quantity, &item.Price)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		items = append(items, item)
	}

	return items, nil
}
