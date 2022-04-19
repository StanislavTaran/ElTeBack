package bank

type BankModel struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	InterestRate   int    `json:"interestRate"`
	MaxLoan        int    `json:"maxLoan"`
	MinDownPayment int    `json:"minDownPayment"`
	LoanTerm       int    `json:"loanTerm"`
}
