package repositoryimplementation

import (
	"context"
	"database/sql"
	"online-store-golang/entity"
	"online-store-golang/errs"
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
