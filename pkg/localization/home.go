package localization

import (
	"errors"
	"time"
)

// HomePageTemplate contains every needed placeholder for the
// 'edit' template. (See static/views/home.html)
// It embeds PageTemplate.
type HomePageTemplate struct {
	TablePageTemplate

	// Header
	EditMode string
	Download string
	Total    string

	NewOperationTitle    string
	EditOperationTitle   string
	EditOperationTooltip string
	AmountType           string
}

// NewHomePageTemplate returns an HomePageTemplate for given year, month and language.
//
// Example:
//  template := NewHomePageTemplate(2020, time.December, languageEnglish)
func NewHomePageTemplate(year int, month time.Month, language Language) (HomePageTemplate, error) {
	template, ok := homePagesByLanguage[language]
	if !ok {
		return template, errors.New("invalid language")
	}
	template.Year = year
	template.Month = monthsByLanguage[language][month-1]
	return template, nil
}
