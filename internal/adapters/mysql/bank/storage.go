package bank

import (
	bankDomain "back/internal/domain/bank"
	"back/pkg/mysqlClient"
	"context"
	"fmt"
)

const (
	tableName = `bank`
)

type Storage struct {
	client *mysqlClient.MySQLClient
}

func NewBankStorage(mysql *mysqlClient.MySQLClient) *Storage {
	return &Storage{
		client: mysql,
	}
}

func (s *Storage) Create(ctx context.Context, dto bankDomain.CreateBankInputDTO) (id int64, err error) {
	query := fmt.Sprintf("INSERT INTO %s (name,interestRate,maxLoan,minDownPayment,loanTerm) VALUES(?,?,?,?,?)", tableName)

	res, err := s.client.Db.ExecContext(
		ctx,
		query,
		&dto.Name,
		&dto.InterestRate,
		&dto.MaxLoan,
		&dto.MinDownPayment,
		&dto.LoanTerm,
	)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) Update(ctx context.Context, dto bankDomain.UpdateBankInputDTO) (err error) {
	query := fmt.Sprintf(`UPDATE %s SET name = ?, interestRate = ?, maxLoan = ?, minDownPayment = ?, loanTerm = ? WHERE id = ?`,
		tableName,
	)

	_, err = s.client.Db.ExecContext(
		ctx,
		query,
		&dto.Name,
		&dto.InterestRate,
		&dto.MaxLoan,
		&dto.MinDownPayment,
		&dto.LoanTerm,
		&dto.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetAll(ctx context.Context) (*[]bankDomain.BankModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	var banks []bankDomain.BankModel

	rows, err := s.client.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var bank bankDomain.BankModel
		err = rows.Scan(
			&bank.Id,
			&bank.Name,
			&bank.InterestRate,
			&bank.MaxLoan,
			&bank.MinDownPayment,
			&bank.LoanTerm,
		)
		if err != nil {
			return nil, err
		}

		banks = append(banks, bank)
	}

	if banks == nil {
		banks = []bankDomain.BankModel{}
	}
	return &banks, nil
}

func (s *Storage) GetById(ctx context.Context, id int64) (*bankDomain.BankModel, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	var bank bankDomain.BankModel

	row := s.client.Db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&bank.Id,
		&bank.Name,
		&bank.InterestRate,
		&bank.MaxLoan,
		&bank.MinDownPayment,
		&bank.LoanTerm,
	)
	if err != nil {
		return nil, err
	}

	return &bank, nil
}

func (s *Storage) Remove(ctx context.Context, id int64) (err error) {
	query := fmt.Sprintf(`DELETE FROM %s  WHERE id = ?`,
		tableName,
	)

	_, err = s.client.Db.ExecContext(
		ctx,
		query,
		&id,
	)
	if err != nil {
		return err
	}

	return nil
}
