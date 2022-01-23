package localization

import "errors"

// SettingsPageTemplate contains every needed placeholder for the
// 'edit' template. (See static/views/edit.html)
// It embeds PageTemplate.
type SettingsPageTemplate struct {
	// Title
	Settings string

	// Table
	SettingLanguage    string
	CurrentLanguage    string
	AvailableLanguages []Language
	SettingServerPort  string
	CurrentServerPort  string
}

// NewSettingsPageTemplate returns an HomePageTemplate for a given language
// and server port.
//
// Example:
//  template := NewSettingsPageTemplate(LanguageEnglish, "8080", "piggy.db")
func NewSettingsPageTemplate(language Language, serverPort string) (SettingsPageTemplate, error) {
	template, ok := settingsPagesByLanguage[language]
	if !ok {
		return template, errors.New("invalid language")
	}
	template.CurrentLanguage = language.String()
	template.CurrentServerPort = serverPort
	return template, nil
}

var availableLanguages = []Language{
	LanguageEnglish,
	LanguageFrench,
}
