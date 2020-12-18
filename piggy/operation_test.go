package piggy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewOperation(t *testing.T) {
	testCases := []struct {
		amount float32
		desc   string
		date   time.Time
	}{
		{42, "some desc", time.Now()},
		{42, "some desc", time.Date(2020, 11, 14, 11, 30, 32, 0, time.UTC)},
		{0, "", time.Time{}},
	}
	for _, tc := range testCases {
		operation := newOperation(tc.amount, tc.desc, tc.date)
		assert.Zero(t, operation.ID)
		assert.Zero(t, operation.CategoryID)
		assert.Equal(t, tc.amount, operation.Amount)
		assert.Equal(t, tc.desc, operation.Description)
		assert.Equal(t, tc.date, operation.Date)
	}
}
