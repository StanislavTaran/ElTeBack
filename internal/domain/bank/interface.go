package bank

import (
	"context"
)

type bankStorage interface {
	Create(ctx context.Context, dto CreateBankInputDTO) (id int64, err error)
	Update(ctx context.Context, dto UpdateBankInputDTO) (err error)
	Remove(ctx context.Context, id int64) (err error)
	GetAll(ctx context.Context) (*[]BankModel, error)
	GetById(ctx context.Context, id int64) (*BankModel, error)
}
