package repository

import (
	"context"
	"encoding/json"
	"marilyn_manson_bot/internal/model"

	"marilyn_manson_bot/pkg/logger"
	"marilyn_manson_bot/pkg/postgres"

	pgx "github.com/jackc/pgx/v5"
)

type DebtRepository interface {
	GetDebtsByCollector(ctx context.Context, collector int64) ([]model.Debt, error)
	GetDebtByCollectorAndDebtor(ctx context.Context, collector int64, debtor string) (*model.Debt, error)
	AddDebt(ctx context.Context, debt *model.Debt) error
	UpdateDebt(ctx context.Context, debt *model.Debt) error
}

type debtRepository struct {
	cluster *postgres.Postgres
	logger  logger.Logger
}

func NewDebtRepo(cluster *postgres.Postgres, logger logger.Logger) *debtRepository {
	return &debtRepository{
		cluster: cluster,
		logger:  logger,
	}
}

func (rep *debtRepository) GetDebtsByCollector(ctx context.Context, collector int64) ([]model.Debt, error) {
	result, err := rep.cluster.Pool.Query(ctx, kGetDebtsByCollectorId, collector)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	debts, err := pgx.CollectRows(result, pgx.RowToStructByName[model.Debt])
	if err != nil {
		return nil, err
	}
	return debts, nil
}

func (rep *debtRepository) AddDebt(ctx context.Context, debt *model.Debt) error {
	trx, err := rep.cluster.Pool.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadWrite})
	if err != nil {
		return err
	}
	debt_record, err := json.Marshal(debt)
	if err != nil {
		return err
	}
	_, err = trx.Exec(ctx, kInsertDebt, debt_record)
	if err != nil {
		trx.Rollback(ctx)
		return err
	}
	trx.Commit(ctx)
	return nil
}

func (rep *debtRepository) UpdateDebt(ctx context.Context, debt *model.Debt) error {
	trx, err := rep.cluster.Pool.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadWrite})
	if err != nil {
		return err
	}
	debt_record, err := json.Marshal(debt)
	if err != nil {
		return err
	}
	res, err := trx.Exec(ctx, kUpdateDebt, debt_record)
	if err != nil {
		trx.Rollback(ctx)
		return err
	}
	if res.RowsAffected() != 1 {
		return NewOptimisticLockError("Cant update debt due to optimistic lock")
	}
	trx.Commit(ctx)
	return nil
}

func (rep *debtRepository) GetDebtByCollectorAndDebtor(ctx context.Context, collector int64, debtor string) (*model.Debt, error) {
	debts, err := rep.GetDebtsByCollector(ctx, collector)
	if err != nil {
		return nil, err
	}
	for _, debt := range debts {
		if debt.DebtId == debtor {
			return &debt, nil
		}
	}
	return nil, nil
}
