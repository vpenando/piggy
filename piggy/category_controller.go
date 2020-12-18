package piggy

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// CategoryController exposes CRUD functions
// for categories.
type CategoryController struct {
	db    *gorm.DB
	Quiet bool
}

// NewCategoryController returns a new CategoryController
// for a given database.
// It updates the Category schema if needed and returns an error if
// something went wrong.
//
// Example:
//  controller, err := NewCategoryController(database)
func NewCategoryController(db *gorm.DB) (*CategoryController, error) {
	controller := &CategoryController{
		db: db,
	}
	err := controller.db.AutoMigrate(&Category{})
	return controller, err
}

// Create inserts a new category into the database and then
// returns it.
// It also returns an error if something went wrong, nil otherwise.
//
// Example:
//  category, err := controller.Create(Category{
//      Name: "My category",
//      Icon: "my_icon.png",
//  })
func (oc *CategoryController) Create(category Category) (Category, error) {
	err := oc.db.Create(&category).Error
	if err != nil {
		log.Println("Error: ", err)
		return category, err
	}
	if !oc.Quiet {
		log.Println("Created category with ID", category.ID)
	}
	return category, nil
}

// ReadAll returns ALL categories from the database.
// It also returns an error if something went wrong.
//
// Example:
//  categories, err := controller.ReadAll()
func (oc *CategoryController) ReadAll() ([]Category, error) {
	var categories []Category
	if err := oc.db.Find(&categories).Error; err != nil {
		if !oc.Quiet {
			log.Println("Error: ", err)
		}
		return nil, err
	}
	return categories, nil
}

// Delete removes all traces of a given category from the database.
// Operations with its ID are updated so that their CategoryID
// is set to 0.
// If the given category doesn't exist, nothing happens.
//
// Examples:
//  err := controller.Delete(42)  // Ok
//  err := controller.Delete(-1)  // Ok too: this is NOT considered as an error
func (oc *CategoryController) Delete(id int) error {
	// If the given category exists in our database:
	//   1. Actually delete category.
	//   2: Update operations => if CategoryID == id, we set it to 0.
	return errors.New("not implemented yet")
}
