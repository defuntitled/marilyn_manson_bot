package model

import (
	"time"

	"github.com/google/uuid"
)

type DebtStatus int

const (
	currency = "RUB"
	active   = iota
	paid
)

func (ds DebtStatus) String() string {
	switch ds {
	case active:
		return "active"
	case paid:
		return "paid"
	default:
		return "venom"
	}
}

type Debt struct {
	DebtId      string     `db:"debt_id" json:"debt_id"`
	Amount      int        `db:"amount" json:"amount"`
	Currency    string     `db:"currency" json:"currency"`
	DebtorId    string     `db:"debtor_id" json:"debtor_id"`
	CollectorId int64      `db:"collector_id" json:"collector_id"`
	Status      DebtStatus `db:"status" json:"status"`
	Version     int        `db:"version" json:"version"`
	CreatedTs   time.Time  `db:"created_ts" json:"created_ts"`
	UpdatedTs   time.Time  `db:"updated_ts" json:"updated_ts"`
}

func NewDebt(debtor_id string, collector_id int64, amount int) *Debt {
	return &Debt{
		DebtId:      uuid.New().String(),
		Amount:      amount,
		Currency:    currency,
		DebtorId:    debtor_id,
		CollectorId: collector_id,
		Status:      active,
		Version:     1,
		CreatedTs:   time.Now(),
		UpdatedTs:   time.Now(),
	}
}

func (d *Debt) SetPaid() {
	d.Status = paid
}

func (d *Debt) AddAmount(am int) {
	d.Amount += am
}
