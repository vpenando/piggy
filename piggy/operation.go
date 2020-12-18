package piggy

import (
	"time"
)

// Operation is the main entity of the program.
// It contains an amount, a title, a description and a date.
// It can also belong to a category.
type Operation struct {
	ID           int       `json:"id"`
	Amount       float32   `json:"amount"`
	CategoryID   int       `json:"category"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	CreationDate time.Time `json:"creation_date"`
}

// newOperation returns a new Operation. Its creation date is the current day
// and it does NOT belong to any category.
func newOperation(amount float32, description string, date time.Time) Operation {
	operation := Operation{
		Amount:       amount,
		Description:  description,
		Date:         date,
		CreationDate: time.Now(),
	}
	return operation
}
