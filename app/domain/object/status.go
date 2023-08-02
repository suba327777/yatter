package object

import "time"

type Status struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"-" db:"account_id"`
	Content   string    `json:"content ,omitempty"`
	CreateAt  time.Time `json:"create_at,omitempty" db:"create_at"`
	Account   Account   `json:"account,omitempty"`
}
