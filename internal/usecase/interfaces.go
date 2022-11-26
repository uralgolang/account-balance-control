package usecase

import (
	"context"
	"github.com/uralgolang/account-balance-control/internal/entity/DTO/request"
)

type IAccountUseCase interface {
	DoDepositAccount(ctx context.Context, req request.DepositAccountRequest) error
	//DoCreditAccount(ctx context.Context, req request.CreditAccountRequest) error
	//DoTransferAccount(ctx context.Context, req request.TransferAccountRequest) error
	//GetBalanceAccount(ctx context.Context, req request.BalanceAccountRequest) (res response.BalanseAccountResponse, err error)
	//GetOperationHistoryAccount(ctx context.Context, req request.OperationHitoryAccountRequest) (res response.OperationHitoryAccountResponse, err error)
}

type IAccountRepo interface {
	InsertAccountOperations() error
	InsertAccountTransfers() error
	InsertAccountBalances() error
	UpdateAccountBalances() error
}
