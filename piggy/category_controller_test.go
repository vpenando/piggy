package piggy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const testDatabaseDSN = "file:memdb1?mode=memory&cache=shared"

var testDatabase *gorm.DB

func init() {
	var err error
	testDatabase, err = gorm.Open(sqlite.Open(testDatabaseDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func cleanTestCategories() {
	// Note: Where condition is required by Gorm.
	testDatabase.Where("1 = 1").Unscoped().Delete(&Category{})
}

func newTestCategoryController(t *testing.T) CategoryController {
	controller, err := NewCategoryController(testDatabase)
	assert.NoError(t, err)
	//controller.Quiet = true // We don't want our controller to log anything.
	return controller
}

func TestNewCategoryController(t *testing.T) {
	controller, err := NewCategoryController(testDatabase)
	assert.NoError(t, err)
	assert.NotNil(t, controller)
}

func TestCreateCategory(t *testing.T) {
	defer cleanTestCategories()
	controller := newTestCategoryController(t)
	category := NewCategory("catName", "catIcon")
	result, err := controller.Create(category)
	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
	assert.Equal(t, category.Name, result.Name)
	assert.Equal(t, category.IconURI, result.IconURI)
}

func TestReadAllCategories(t *testing.T) {
	defer cleanTestCategories()
	controller := newTestCategoryController(t)

	// Empty database
	categories, err := controller.ReadAll()
	assert.NoError(t, err)
	assert.Empty(t, categories)

	// Database with 1 record
	category := NewCategory("catName", "catIcon")
	assert.Zero(t, category.ID)
	err = testDatabase.Create(&category).Error
	assert.NoError(t, err)
	categories, err = controller.ReadAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, categories)
	assert.NotZero(t, category.ID)
	assert.Equal(t, category, categories[0])
}
