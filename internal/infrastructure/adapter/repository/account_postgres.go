package repository

import "github.com/jackc/pgx/v4"

type IAccountRepoImpl struct {
	Pg *pgx.Conn
}

func NewAccountRepo(pg *pgx.Conn) *IAccountRepoImpl {
	return &IAccountRepoImpl{
		Pg: pg,
	}
}

func (accountRepo *IAccountRepoImpl) InsertAccountOperations() error {
	return nil
}

func (accountRepo *IAccountRepoImpl) InsertAccountTransfers() error {
	return nil
}

func (accountRepo *IAccountRepoImpl) InsertAccountBalances() error {
	return nil
}

func (accountRepo *IAccountRepoImpl) UpdateAccountBalances() error {
	return nil
}
