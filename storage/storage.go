package storage

import (
	"bankomat/api/models"
	"context"
)

type StorageI interface {
	Account() AccountRepoI
}

type AccountRepoI interface {
	Create(ctx context.Context, req *models.CreateAccount) error
	Deposit(ctx context.Context, req *models.Deposit) error
	Withdraw(ctx context.Context, req *models.Withdraw) error
	GetBalance(ctx context.Context, req string) float64
}
