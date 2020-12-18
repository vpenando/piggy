package piggy

// Category can be considered as the group to which an operation belongs.
// A category contains a name and has an icon. A category with ID 0
// is considered as not stored in the database.
type Category struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	IconURI string `json:"icon"`
}

// NewCategory returns a new category with  a name
// and an icon URI. It has a zero ID as it is not
// stored in the database.
//
// Example:
//  category := NewCategory("Shopping", "assets/icon_shopping.png")
func NewCategory(name, icon string) Category {
	category := Category{
		Name:    name,
		IconURI: icon,
	}
	return category
}
