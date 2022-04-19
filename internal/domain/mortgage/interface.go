package mortgage

import (
	"back/internal/domain/bank"
	"context"
)

type bankService interface {
	GetById(ctx context.Context, id int64) (*bank.BankModel, error)
}
