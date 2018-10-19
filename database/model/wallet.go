package dbmodel

import (
	"log"
	"time"

	"github.com/ExpenseLedger/expense-ledger-web-service/database"
	"github.com/shopspring/decimal"
)

// Wallet the structure represents a stored wallet on database
type Wallet struct {
	Name      string          `db:"name"`
	Type      string          `db:"type"`
	Balance   decimal.Decimal `db:"balance"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
	DeletedAt *time.Time      `db:"deleted_at"`
}

// Insert inserts a wallet into the database
func (wallet Wallet) Insert() (*Wallet, error) {
	query :=
		`
		INSERT INTO wallet (name, type, balance)
		VALUES (:name, :type, :balance)
		RETURNING *;
		`
	db := database.GetDB()

	stmt, err := db.PrepareNamed(query)
	if err != nil {
		log.Println("Error inserting a wallet", err)
		return nil, err
	}

	var createdWallet Wallet
	if err := stmt.Get(&createdWallet, &wallet); err != nil {
		log.Println("Error inserting a wallet", err)
		return nil, err
	}

	return &createdWallet, nil
}
