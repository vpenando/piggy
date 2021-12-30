package piggy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	categoryName = "name"
	categoryIcon = "icon"
)

func TestNewCategory(t *testing.T) {
	category := NewCategory(categoryName, categoryIcon)
	assert.Zero(t, category.ID)
	assert.Equal(t, categoryName, category.Name)
	assert.Equal(t, categoryIcon, category.IconURI)
}
