package bank

import "context"

type BankService struct {
	bankStorage bankStorage
}

func NewBankService(bankStorage bankStorage) *BankService {
	return &BankService{
		bankStorage: bankStorage,
	}
}

func (c *BankService) GetAll(ctx context.Context) (*[]BankModel, error) {
	return c.bankStorage.GetAll(ctx)
}

func (c *BankService) GetById(ctx context.Context, id int64) (*BankModel, error) {
	return c.bankStorage.GetById(ctx, id)
}

func (c *BankService) Create(ctx context.Context, dto CreateBankInputDTO) (int64, error) {
	return c.bankStorage.Create(ctx, dto)
}

func (c *BankService) Update(ctx context.Context, dto UpdateBankInputDTO) error {
	return c.bankStorage.Update(ctx, dto)
}

func (c *BankService) Remove(ctx context.Context, id int64) error {
	return c.bankStorage.Remove(ctx, id)
}
