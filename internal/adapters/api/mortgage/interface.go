package mortgage

import (
	"back/internal/domain/mortgage"
	"context"
)

type mortgageService interface {
	CalculateMonthlyPayment(ctx context.Context, data mortgage.CalculateMortgagePaymentInputDTO) (*mortgage.CalculateMortgagePaymentOutputDTO, error)
}
