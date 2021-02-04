package localization

// TableColumns contains the translation of the
// columns names for a given language.
type TableColumns struct {
	Category     string
	Date         string
	Description  string
	Amount       string
	CreationDate string
}

// TablePageTemplate embeds PageTemplate and adds
// some common placeholders needed by home and edit pages.
type TablePageTemplate struct {
	TableColumns

	// Header
	Year  int
	Month string

	// Search bar
	Search string

	// Table
	Category     string
	Date         string
	Description  string
	Amount       string
	CreationDate string
}
