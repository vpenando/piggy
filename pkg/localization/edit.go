package localization

import (
	"errors"
	"time"
)

// EditPageTemplate contains every needed placeholder for the
// 'edit' template. (See static/views/edit.html)
// It embeds PageTemplate.
type EditPageTemplate struct {
	TablePageTemplate

	// Date
	MonthIndex int

	// New category
	NewCategoryTitle           string
	NewCategoryName            string
	NewCategoryIcon            string
	NewCategoryNamePlaceholder string
	NewCategoryButton          string

	// Tooltips
	TooltipAdd         string
	TooltipEdit        string
	TooltipAddCategory string
	TooltipDelete      string
}

// NewEditPageTemplate returns an EditPageTemplate for given year, month and language.
//
// Example:
//  template := NewEditPageTemplate(2020, time.December, languageEnglish)
func NewEditPageTemplate(year int, month time.Month, language Language) (EditPageTemplate, error) {
	template, ok := editPagesByLanguage[language]
	if !ok {
		return template, errors.New("invalid language")
	}
	template.Year = year
	template.Month = monthsByLanguage[language][month-1]
	template.MonthIndex = int(month) - 1
	return template, nil
}
