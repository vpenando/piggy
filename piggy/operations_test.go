package piggy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOperationsWhere(t *testing.T) {
	now := time.Now()
	operations := Operations{
		Operation{},
		Operation{ID: 1},
		Operation{ID: 2, Amount: 10},
		Operation{ID: 3, Amount: 20, CategoryID: 0},
		Operation{ID: 4, Amount: 30, CategoryID: 1, Description: "d"},
		Operation{ID: 5, Amount: 40, CategoryID: 2, Description: "desc", Date: now},
		Operation{ID: 6, Amount: 59.5, CategoryID: 3, Description: "description", Date: now, CreationDate: now},
	}
	testCases := []struct {
		name      string
		expected  Operations
		predicate func(Operation) bool
	}{
		{"true", operations, func(op Operation) bool { return true }},
		{"false", Operations{}, func(op Operation) bool { return false }},
		{"ID==0", operations[:1], func(op Operation) bool { return op.ID == 0 }},
		{"ID!=0", operations[1:], func(op Operation) bool { return op.ID != 0 }},
		{"Amount==0", operations[:2], func(op Operation) bool { return op.Amount == 0 }},
		{"Amount!=0", operations[2:], func(op Operation) bool { return op.Amount != 0 }},
		{"len(Description)==0", operations[:4], func(op Operation) bool { return len(op.Description) == 0 }},
		{"len(Description)!=0", operations[4:], func(op Operation) bool { return len(op.Description) != 0 }},
		{"Date==now", operations[5:], func(op Operation) bool { return op.Date == now }},
		{"Date!=now", operations[:5], func(op Operation) bool { return op.Date != now }},
		{"CreationDate==now", operations[6:], func(op Operation) bool { return op.CreationDate == now }},
		{"CreationDate!=now", operations[:6], func(op Operation) bool { return op.CreationDate != now }},
	}
	for _, tc := range testCases {
		ops := operations.Where(tc.predicate)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, ops)
			for _, op := range ops {
				assert.True(t, tc.predicate(op))
			}
		})
	}
}
