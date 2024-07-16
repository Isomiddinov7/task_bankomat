package postgres

import (
	"bankomat/api/models"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type accountRepo struct {
	db *pgxpool.Pool
}

func NewAccountRepo(db *pgxpool.Pool) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

func (r *accountRepo) Create(ctx context.Context, req *models.CreateAccount) error {
	var (
		id    = uuid.New().String()
		query = `
			INSERT INTO "account"(
				"id",
				"name",
				"phone",
				"balance"
			) VALUES($1, $2, $3, $4)`
	)
	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Phone,
		req.Balance,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepo) Deposit(ctx context.Context, req *models.Deposit) error {
	var (
		queryDeposit = `
			UPDATE "account"
				SET "balance" = $2,
				"updated_at" = NOW()

			WHERE "id" = $1
		`
		query = `
			SELECT
				"balance"
			FROM "account"
			WHERE "id" = $1
		`
		balance sql.NullFloat64
		sum     float64
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(&balance)
	if err != nil {
		return err
	}
	sum = balance.Float64 + req.Amount
	_, err = r.db.Exec(ctx, queryDeposit,
		req.Id,
		sum,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *accountRepo) Withdraw(ctx context.Context, req *models.Withdraw) error {
	var (
		queryWithdraw = `
			UPDATE "account"
				SET "balance" = $2,
				"updated_at" = NOW()

			WHERE "id" = $1
		`
		query = `
			SELECT
				"balance"
			FROM "account"
			WHERE "id" = $1
		`
		balance sql.NullFloat64
		sum     float64
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(&balance)
	if err != nil {
		return err
	}

	if balance.Float64 >= req.Amount {
		sum = balance.Float64 - req.Amount
		_, err = r.db.Exec(ctx, queryWithdraw,
			req.Id,
			sum,
		)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Not enough money")
	}

	return nil
}

func (r *accountRepo) GetBalance(ctx context.Context, req string) float64 {
	var (
		query = `
	SELECT
		"balance"
	FROM "account"
	WHERE "id" = $1
	`
		balance sql.NullFloat64
	)
	err := r.db.QueryRow(ctx, query, req).Scan(&balance)
	if err != nil {
		return 0
	}
	return balance.Float64
}
