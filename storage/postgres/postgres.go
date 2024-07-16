package postgres

import (
	"bankomat/config"
	"bankomat/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db      *pgxpool.Pool
	account storage.AccountRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	),
	)
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pgxpool,
	}, err
}

func (s *Store) Account() storage.AccountRepoI {
	if s.account == nil {
		s.account = NewAccountRepo(s.db)
	}

	return s.account
}
