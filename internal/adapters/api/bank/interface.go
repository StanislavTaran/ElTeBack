package bank

import (
	"back/internal/domain/bank"
	"context"
)

type bankService interface {
	GetAll(ctx context.Context) (*[]bank.BankModel, error)
	Create(ctx context.Context, dto bank.CreateBankInputDTO) (int64, error)
	Update(ctx context.Context, dto bank.UpdateBankInputDTO) error
	Remove(ctx context.Context, id int64) error
}
