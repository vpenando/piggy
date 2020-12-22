package main

import (
	"errors"
	"time"
)

// HomePageTemplate contains every needed placeholder for the
// 'edit' template. (See views/home.html)
// It embeds PageTemplate.
type HomePageTemplate struct {
	TablePageTemplate

	// Header
	EditMode string
	Download string
	Total    string

	EditOperationTitle string
}

// newHomePageTemplate returns an HomePageTemplate for given year, month and language.
//
// Example:
//  template := newHomePageTemplate(2020, time.December, languageEnglish)
func newHomePageTemplate(year int, month time.Month, language Language) (HomePageTemplate, error) {
	template, ok := homePagesByLanguage[language]
	if !ok {
		return template, errors.New("invalid language")
	}
	template.Year = year
	template.Month = monthsByLanguage[language][month-1]
	return template, nil
}
