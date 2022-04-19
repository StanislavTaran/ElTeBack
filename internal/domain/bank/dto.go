package bank

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateBankInputDTO struct {
	Name           string `json:"name"`
	InterestRate   int    `json:"interestRate"`
	MaxLoan        int    `json:"maxLoan"`
	MinDownPayment int    `json:"minDownPayment"`
	LoanTerm       int    `json:"loanTerm"`
}

func (c CreateBankInputDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(2, 20)),
		validation.Field(&c.InterestRate, validation.Required, validation.Min(1), validation.Max(99)),
		validation.Field(&c.MaxLoan, validation.Required, validation.Min(1), validation.Max(100000000)),
		validation.Field(&c.MinDownPayment, validation.Required, validation.Min(1), validation.Max(100000000)),
		validation.Field(&c.LoanTerm, validation.Required, validation.Min(1), validation.Max(240)),
	)
}

type UpdateBankInputDTO struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	InterestRate   int    `json:"interestRate"`
	MaxLoan        int    `json:"maxLoan"`
	MinDownPayment int    `json:"minDownPayment"`
	LoanTerm       int    `json:"loanTerm"`
}

func (c UpdateBankInputDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Id, validation.Required, validation.Min(1)),
		validation.Field(&c.Name, validation.Required, validation.Length(2, 20)),
		validation.Field(&c.InterestRate, validation.Required, validation.Min(1), validation.Max(99)),
		validation.Field(&c.MaxLoan, validation.Required, validation.Min(1), validation.Max(100000000)),
		validation.Field(&c.MinDownPayment, validation.Required, validation.Min(1), validation.Max(100000000)),
		validation.Field(&c.LoanTerm, validation.Required, validation.Min(1), validation.Max(240)),
	)
}
