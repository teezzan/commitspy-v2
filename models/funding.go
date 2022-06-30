package models

import (
	"time"

	"github.com/google/uuid"
)

type Funding struct {
	ID            int64     `gorm:"id, primarykey, autoincrement" json:"id"`
	User          User      `json:"user"`
	Project       Project   `json:"project"`
	Amount        int64     `json:"amount"`
	TransactionID uuid.UUID `json:"transaction_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UpdatedAt     time.Time `json:"-"`
	CreatedAt     time.Time `json:"created"`
}
