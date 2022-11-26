package usecase

import (
	"context"
	"github.com/uralgolang/account-balance-control/internal/entity/DTO/request"
	"github.com/uralgolang/account-balance-control/internal/infrastructure/adapter/repository"
)

type IAccountUseCaseImpl struct {
	Repo *repository.IAccountRepoImpl
}

func NewAccountUseCase(repo *repository.IAccountRepoImpl) *IAccountUseCaseImpl {
	return &IAccountUseCaseImpl{
		Repo: repo,
	}
}

func (accountUseCase *IAccountUseCaseImpl) DoDepositAccount(ctx context.Context, req request.DepositAccountRequest) error {
	return nil
}

func (accountUseCase *IAccountUseCaseImpl) DoCreditAccount() error {
	return nil
}

func (accountUseCase *IAccountUseCaseImpl) DoTransferAccount() error {
	return nil
}

func (accountUseCase *IAccountUseCaseImpl) GetBalanceAccount() error {
	return nil
}

func (accountUseCase *IAccountUseCaseImpl) GetOperationHistoryAccount() error {
	return nil
}
