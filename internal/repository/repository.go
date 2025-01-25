package repository

import (
	"database/sql"
	"fmt"

	"github.com/TemirB/rest-api-wallet/internal/models"
)

type Repository struct {
	DB *sql.DB
}

func (s *Repository) GetWallet(wallet models.Wallet) (*models.Wallet, error) {
	var w models.Wallet
	err := s.DB.QueryRow("SELECT * FROM wallets WHERE id = $1", wallet.Id).Scan(&w.Id, &w.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			err = s.CreateWallet(wallet)
			if err != nil {
				return nil, fmt.Errorf("could not create wallet: %w", err)
			}
			return &wallet, nil
		}
		return nil, fmt.Errorf("could not get wallet: %w", err)
	}
	return &w, nil
}

func (s *Repository) UpdateWallet(wallet models.Wallet) error {
	result, err := s.DB.Exec("UPDATE wallets SET balance = $1 WHERE id = $2", wallet.Balance, wallet.Id)
	if err != nil {
		return fmt.Errorf("could not update wallet: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not determine rows affected: %w", err)
	}
	if rowsAffected == 0 {
		err = s.CreateWallet(wallet)
		if err != nil {
			return fmt.Errorf("could not create wallet: %w", err)
		}
	}
	return nil
}

func (s *Repository) CreateWallet(wallet models.Wallet) error {
	_, err := s.DB.Exec("INSERT INTO wallets (id, balance) VALUES ($1, $2)", wallet.Id, wallet.Balance)
	if err != nil {
		return fmt.Errorf("could not create wallet: %w", err)
	}
	return nil
}
