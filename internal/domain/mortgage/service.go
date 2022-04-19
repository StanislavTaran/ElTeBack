package mortgage

import (
	"context"
	"errors"
	"fmt"
	"math"
)

type mortgageService struct {
	bankService bankService
}

func NewMortgageService(bankService bankService) *mortgageService {
	return &mortgageService{
		bankService: bankService,
	}
}

func (m *mortgageService) CalculateMonthlyPayment(ctx context.Context, data CalculateMortgagePaymentInputDTO) (*CalculateMortgagePaymentOutputDTO, error) {
	bank, err := m.bankService.GetById(ctx, data.BankId)
	if err != nil {
		return nil, err
	}

	if data.InitialLoan > float64(bank.MaxLoan) {
		return nil, errors.New(fmt.Sprintf("Max Loan is : %.2f", float64(bank.MaxLoan)))
	}

	if data.DownPayment < float64(bank.MinDownPayment) {
		return nil, errors.New(fmt.Sprintf("MIN Down Payment is : %.2f", float64(bank.MinDownPayment)))
	}

	if data.DownPayment > data.InitialLoan {
		return nil, errors.New("MIN Down Payment is greater then Initial Loan")
	}

	loanSum := data.InitialLoan - data.DownPayment
	payment := (loanSum * (float64(bank.InterestRate) / 12) * math.Pow(1+float64(bank.InterestRate)/12, float64(bank.LoanTerm))) / (math.Pow(1+float64(bank.InterestRate)/12, float64(bank.LoanTerm)) - 1)

	return &CalculateMortgagePaymentOutputDTO{MonthlyPayment: payment}, nil
}
