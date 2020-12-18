package main

// PageTemplate contains the common placeholders used
// in the application templates.
type PageTemplate struct {
	Title   string
	Version string
}

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
	PageTemplate
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
