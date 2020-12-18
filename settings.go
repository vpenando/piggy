package main

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

func newSettingsPageTemplate(language Language) (SettingsPageTemplate, error) {
	template, ok := settingsPagesByLanguage[language]
	if !ok {
		return template, errors.New("invalid language")
	}
	template.CurrentLanguage = currentLanguage.String()
	template.CurrentServerPort = serverPort
	template.CurrentServerDatabase = serverDatabase
	return template, nil
}

var availableLanguages = []Language{
	languageEnglish,
	languageFrench,
}
