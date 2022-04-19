package mortgage

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CalculateMortgagePaymentInputDTO struct {
	BankId      int64   `json:"bankId"`
	InitialLoan float64 `json:"initialLoan"`
	DownPayment float64 `json:"downPayment"`
}

type CalculateMortgagePaymentOutputDTO struct {
	MonthlyPayment float64 `json:"monthlyPayment"`
}

func (c CalculateMortgagePaymentInputDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.BankId, validation.Required, validation.Min(1)),
		validation.Field(&c.InitialLoan, validation.Required, validation.Min(1.0)),
		validation.Field(&c.DownPayment, validation.Required, validation.Min(1.0)),
	)
}
