package service

import (
	"fmt"

	"github.com/TemirB/rest-api-wallet/internal/models"
	"github.com/TemirB/rest-api-wallet/internal/repository"
	"github.com/google/uuid"
)

type Service struct {
	Repo *repository.Repository
}

func (s *Service) UpdateWalletBalance(req models.WalletRequest) error {
	wallet, err := s.Repo.GetWallet(models.Wallet{Id: req.Id})
	if err != nil {
		return fmt.Errorf("wallet not found: %w", err)
	}

	switch req.OperationType {
	case "DEPOSIT":
		wallet.Balance += req.Amount
	case "WITHDRAW":
		if wallet.Balance < req.Amount {
			return fmt.Errorf("insufficient balance")
		}
		wallet.Balance -= req.Amount
	default:
		return fmt.Errorf("unknown operation type")
	}

	err = s.Repo.UpdateWallet(*wallet)
	if err != nil {
		return fmt.Errorf("error updating wallet: %w", err)
	}

	return nil
}

func (s *Service) GetWalletBalance(walletID uuid.UUID) (float64, error) {
	wallet, err := s.Repo.GetWallet(models.Wallet{Id: walletID})
	if err != nil {
		return 0, fmt.Errorf("wallet not found: %w", err)
	}

	return wallet.Balance, nil
}
