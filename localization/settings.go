package localization

import "errors"

// SettingsPageTemplate contains every needed placeholder for the
// 'edit' template. (See views/edit.html)
// It embeds PageTemplate.
type SettingsPageTemplate struct {
	PageTemplate

	// Title
	Settings string

	// Table
	SettingLanguage       string
	CurrentLanguage       string
	AvailableLanguages    []Language
	SettingServerPort     string
	CurrentServerPort     string
	SettingServerDatabase string
	CurrentServerDatabase string
}

// NewSettingsPageTemplate returns an HomePageTemplate for a given language,
// server port and database.
//
// Example:
//  template := NewSettingsPageTemplate(LanguageEnglish, "8080", "piggy.db")
func NewSettingsPageTemplate(language Language, serverPort, serverDatabase string) (SettingsPageTemplate, error) {
	template, ok := settingsPagesByLanguage[language]
	if !ok {
		return template, errors.New("invalid language")
	}
	template.CurrentLanguage = language.String()
	template.CurrentServerPort = serverPort
	template.CurrentServerDatabase = serverDatabase
	return template, nil
}

var availableLanguages = []Language{
	LanguageEnglish,
	LanguageFrench,
}
