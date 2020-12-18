package piggy

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func cleanTestOperations() {
	// Note: Where condition is required by Gorm.
	testDatabase.Where("1 = 1").Unscoped().Delete(&Operation{})
}

func newTestOperationController(t *testing.T) *OperationController {
	controller, err := NewOperationController(testDatabase)
	assert.NoError(t, err)
	controller.Quiet = true // We don't want our controller to log anything.
	return controller
}

func TestNewOperationController(t *testing.T) {
	controller := newTestOperationController(t)
	assert.Equal(t, testDatabase, controller.db)
}

func TestCreateManyOperations(t *testing.T) {
	defer cleanTestOperations()
	controller := newTestOperationController(t)
	testCases := Operations{
		newOperation(1., "new operation 1", time.Now()),
		newOperation(2., "new operation 2", time.Now()),
	}
	newOperations, err := controller.CreateMany(testCases)
	operations, err := controller.ReadAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, newOperations)
	assert.Equal(t, len(testCases), len(newOperations))

	for i := 0; i < len(newOperations); i++ {
		assert.NotZero(t, newOperations[i].ID)
		assert.Equal(t, operations[i].ID, newOperations[i].ID)
		assert.Equal(t, testCases[i].Amount, newOperations[i].Amount)
		assert.Equal(t, testCases[i].CategoryID, newOperations[i].CategoryID)
		assert.Equal(t, testCases[i].Description, newOperations[i].Description)
		assert.Equal(t, testCases[i].Date.UTC(), newOperations[i].Date.UTC())
		assert.Equal(t, testCases[i].CreationDate.UTC(), newOperations[i].CreationDate.UTC())
	}
}

func TestReadAllOperations(t *testing.T) {
	defer cleanTestOperations()
	controller := newTestOperationController(t)

	// Empty database
	t.Run("EmptyDatabase", func(t *testing.T) {
		operations, err := controller.ReadAll()
		assert.NoError(t, err)
		assert.Empty(t, operations)
	})

	// Database with 1 record
	t.Run("DatabaseNotEmpty", func(t *testing.T) {
		operation := newOperation(3.14, "someOperation", time.Now())
		assert.Zero(t, operation.ID)
		err := testDatabase.Create(&operation).Error
		assert.NoError(t, err)
		operations, err := controller.ReadAll()
		assert.NoError(t, err)
		assert.NotEmpty(t, operations)
		assert.NotZero(t, operation.ID)
		assert.Equal(t, operation.ID, operations[0].ID)
		assert.Equal(t, operation.Amount, operations[0].Amount)
		assert.Equal(t, operation.CategoryID, operations[0].CategoryID)
		assert.Equal(t, operation.Description, operations[0].Description)
		assert.Equal(t, operation.Date.UTC(), operations[0].Date.UTC())
		assert.Equal(t, operation.CreationDate.UTC(), operations[0].CreationDate.UTC())
	})
}

func TestUpdateManyOperations(t *testing.T) {
	defer cleanTestOperations()
	controller := newTestOperationController(t)
	operations, err := controller.CreateMany(Operations{
		newOperation(10., "new operation 1", time.Now()),
		newOperation(-20., "new operation 2", time.Now()),
		newOperation(30., "new operation 3", time.Now()), // This one should NOT be updated.
	})
	assert.NoError(t, err)
	assert.Equal(t, 3, len(operations))
	// At this point:
	//   - There are exactly three operations in our database.

	operationsToUpdate := make(Operations, 0, 2)
	for i := 0; i < 2; i++ {
		operation := operations[i]
		operation.Amount *= 10
		operation.Description = strings.ReplaceAll(operation.Description, "new", "updated")
		operationsToUpdate = append(operationsToUpdate, operation)
	}

	// At this point:
	//   - There are exactly three operations in our database;
	//   - We have updated only two of them.

	// Updated operations
	t.Run("UpdatedOperations", func(t *testing.T) {
		updatedOperations, err := controller.UpdateMany(Operations(operationsToUpdate))
		assert.NoError(t, err)
		assert.Equal(t, 2, len(updatedOperations))
		for i := 0; i < 2; i++ {
			// Non-updated fields
			assert.Equal(t, operations[i].ID, updatedOperations[i].ID)
			assert.Equal(t, operations[i].CategoryID, updatedOperations[i].CategoryID)
			assert.Equal(t, operations[i].Date.UTC(), updatedOperations[i].Date.UTC())
			assert.Equal(t, operations[i].CreationDate.UTC(), updatedOperations[i].CreationDate.UTC())
			// Amount
			assert.Equal(t, operationsToUpdate[i].Amount, updatedOperations[i].Amount)
			assert.Equal(t, operations[i].Amount*10, updatedOperations[i].Amount)
			// Description
			assert.Equal(t, operationsToUpdate[i].Description, updatedOperations[i].Description)
			expectedDescription := strings.ReplaceAll(operations[i].Description, "new", "updated")
			assert.Equal(t, expectedDescription, updatedOperations[i].Description)
		}
	})

	// Non-updated operation
	t.Run("NonUpdatedOperations", func(t *testing.T) {
		// Operation with index 2 was not updated.
		allOperations, err := controller.ReadAll()
		assert.NoError(t, err)
		assert.Equal(t, operations[2].ID, allOperations[2].ID)
		assert.Equal(t, operations[2].Amount, allOperations[2].Amount)
		assert.Equal(t, operations[2].Description, allOperations[2].Description)
		assert.Equal(t, operations[2].CategoryID, allOperations[2].CategoryID)
		assert.Equal(t, operations[2].Date.UTC(), allOperations[2].Date.UTC())
		assert.Equal(t, operations[2].CreationDate.UTC(), allOperations[2].CreationDate.UTC())
	})
}

func TestDeleteManyOperations(t *testing.T) {
	defer cleanTestOperations()
	controller := newTestOperationController(t)
	var count int64
	testDatabase.Table("operations").Count(&count)
	assert.Zero(t, count)

	t.Run("EmptyDatabase", func(t *testing.T) {
		err := controller.DeleteMany([]int{1, 2, 3})
		assert.NoError(t, err)
		testDatabase.Table("operations").Count(&count)
		assert.Zero(t, count)
	})

	t.Run("DatabaseNotEmpty", func(t *testing.T) {
		_, err := controller.CreateMany(Operations{
			newOperation(10., "new operation 1", time.Now()),
			newOperation(-20., "new operation 2", time.Now()),
			newOperation(30., "new operation 3", time.Now()), // This one should NOT be updated.
		})
		assert.NoError(t, err)
		testDatabase.Table("operations").Count(&count)
		assert.Equal(t, int64(3), count)
		err = controller.DeleteMany([]int{1, 2, 3})
		assert.NoError(t, err)
		testDatabase.Table("operations").Count(&count)
		assert.Zero(t, count)
	})
}
