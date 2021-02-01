package piggy

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

// OperationController exposes CRUD functions or operations.
type OperationController struct {
	db    *gorm.DB
	Quiet bool
}

// NewOperationController returns a new OperationController
// for a given database.
// It updates the Operation schema if needed and returns an error if
// something went wrong.
//
// Example:
//  controller, err := NewOperationController(database)
func NewOperationController(db *gorm.DB) (*OperationController, error) {
	controller := &OperationController{
		db: db,
	}
	err := controller.db.AutoMigrate(&Operation{})
	return controller, err
}

// CreateMany inserts new operations into the database and then
// returns them.
// It also returns an error if something went wrong, nil otherwise.
// Nothing happens if no operation was provided.
//
// Example:
//  operations, err := controller.CreateMany(Operations{
//      operation1,
//      operation2,
//      ...,
//  })
func (oc *OperationController) CreateMany(operations Operations) (Operations, error) {
	if len(operations) == 0 {
		return operations, nil
	}
	created := make(Operations, 0, len(operations))
	err := oc.db.Transaction(func(tx *gorm.DB) error {
		for _, op := range operations {
			if e := tx.Create(&op).Error; e != nil {
				return e
			}
			created = append(created, op)
		}
		return nil
	})
	return created, err
}

// ReadAll returns ALL operations from the database.
// It also returns an error if something went wrong, nil otherwise.
// Nothing happens if no operation was provided.
//
// Example:
//  operations, err := controller.ReadAll()
func (oc *OperationController) ReadAll() (operations Operations, err error) {
	err = oc.db.Find(&operations).Error
	return
}

// ReadAllBetween returns ALL operations that are between two given dates.
// It also returns an error if something went wrong, nil otherwise.
//
// Example:
//  startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
//  endDate := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.Local)
//  operations, err := controller.ReadAllBetween(startDate, endDate)
func (oc *OperationController) ReadAllBetween(startDate, endDate time.Time) (operations Operations, err error) {
	err = oc.db.Where("date >= ? AND date < ?", startDate.String(), endDate.String()).Find(&operations).Error
	return
}

// UpdateMany updates given operations. Nothing happens if
// no operation was provided.
//
// Example:
//  updated, err := controller.UpdateMany(Operations{
//      operation1,
//      operation2,
//      ...,
//  })
func (oc *OperationController) UpdateMany(operations Operations) (Operations, error) {
	if len(operations) == 0 {
		return operations, nil
	}
	unexistingOperations := operations.Where(func(op Operation) bool { return op.ID == 0 })
	if len(unexistingOperations) > 0 {
		// ID 0 means that these operations do NOT exist in database.
		return Operations{}, errors.New("cannot update operation with ID 0")
	}
	updated := make(Operations, 0, len(operations))
	err := oc.db.Transaction(func(tx *gorm.DB) error {
		for _, op := range operations {
			newOperation := map[string]interface{}{
				"amount":      op.Amount,
				"category_id": op.CategoryID,
				"description": op.Description,
				"date":        op.Date,
			}
			err := oc.db.Model(&op).Where("id = ?", op.ID).Updates(newOperation).Error
			if err != nil {
				return err
			}
			updated = append(updated, op)
			if !oc.Quiet {
				log.Println("Updated operation with ID", op.ID)
			}
		}
		return nil
	})
	return updated, err
}

// DeleteMany deletes operations with given IDs. Nothing happens
// if no ID was provided.
//
// Example:
//  err := controller.DeleteMany([]int{1, 2, 3})
func (oc *OperationController) DeleteMany(ids []int) error {
	if len(ids) == 0 {
		return nil
	}
	err := oc.db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if e := oc.db.Delete(&Operation{}, id).Error; e != nil {
				return e
			}
			if !oc.Quiet {
				log.Println("Deleted operation with ID", id)
			}
		}
		return nil
	})
	return err
}
