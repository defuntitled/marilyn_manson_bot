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
	GetDebtsByDabtorsAndCollector(ctx context.Context, collector int64) ([]model.Debt, error)
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

func (this *debtRepository) GetDebtsByDabtorsAndCollector(ctx context.Context, collector int64) ([]model.Debt, error) {
	result, err := this.cluster.Pool.Query(ctx, kGetDebtsByCollectorId, collector)
	defer result.Close()
	if err != nil {
		return nil, err
	}
	debts, err := pgx.CollectRows(result, pgx.RowToStructByName[model.Debt])
	if err != nil {
		return nil, err
	}
	return debts, nil
}

func (this *debtRepository) AddDebt(ctx context.Context, debt *model.Debt) error {
	trx, err := this.cluster.Pool.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadWrite})
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

func (this *debtRepository) UpdateDebt(ctx context.Context, debt *model.Debt) error {
	trx, err := this.cluster.Pool.BeginTx(ctx, pgx.TxOptions{AccessMode: pgx.ReadWrite})
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
		return NewOptimisticLockError("Cant update debt due to optimistic lock", pgx.ErrNoRows)
	}
	trx.Commit(ctx)
	return nil
}
