package main

// Language is a type that is used to localize strings.
type Language int

// Here are the different available languages.
const (
	languageEnglish Language = iota
	languageFrench
)

// To satisfy the fmt.Stringer interface.
func (l Language) String() string {
	switch l {
	case languageEnglish:
		return "English"
	case languageFrench:
		return "Français"
	}
	return "???"
}

// /!\ TODO - Move everything out in an external config file.

var monthsByLanguage = map[Language][]string{
	languageEnglish: {
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	},
	languageFrench: {
		"Janvier",
		"Février",
		"Mars",
		"Avril",
		"Mai",
		"Juin",
		"Juillet",
		"Août",
		"Septembre",
		"Octobre",
		"Novembre",
		"Décembre",
	},
}

var dateFormatsByLanguage = map[Language]string{
	languageEnglish: "2006-01-02",
	languageFrench:  "02/01/2006",
}

var columnsByLanguage = map[Language]TableColumns{
	languageEnglish: {
		Category:     "Category",
		Date:         "Date",
		Description:  "Description",
		Amount:       "Amount",
		CreationDate: "Created at",
	},
	languageFrench: {
		Category:     "Catégorie",
		Date:         "Date",
		Description:  "Description",
		Amount:       "Montant",
		CreationDate: "Créé le",
	},
}

var homePagesByLanguage = map[Language]HomePageTemplate{
	languageEnglish: {
		TablePageTemplate: TablePageTemplate{
			PageTemplate: PageTemplate{
				Title:   applicationName,
				Version: applicationVersion,
			},
			Search:       "Search...",
			Category:     "Category",
			Date:         "Date",
			Description:  "Description",
			Amount:       "Amount",
			CreationDate: "Created at",
			TableColumns: columnsByLanguage[languageEnglish],
		},
		Total:    "Total:",
		Download: "Download",
		EditMode: "Edit mode",
	},
	languageFrench: {
		TablePageTemplate: TablePageTemplate{
			PageTemplate: PageTemplate{
				Title:   applicationName,
				Version: applicationVersion,
			},
			Search:       "Rechercher...",
			Category:     "Catégorie",
			Date:         "Date",
			Description:  "Description",
			Amount:       "Montant",
			CreationDate: "Créé le",
			TableColumns: columnsByLanguage[languageFrench],
		},
		Total:    "Total :",
		Download: "Télécharger",
		EditMode: "Mode édition",
	},
}

var editPagesByLanguage = map[Language]EditPageTemplate{
	languageEnglish: {
		TablePageTemplate: TablePageTemplate{
			PageTemplate: PageTemplate{
				Title:   applicationName,
				Version: applicationVersion,
			},
			Search:       "Search...",
			Category:     "Category",
			Date:         "Date",
			Description:  "Description",
			Amount:       "Amount",
			CreationDate: "Created at",
		},
		NewCategoryTitle:           "New category",
		NewCategoryName:            "Name:",
		NewCategoryIcon:            "Icon:",
		NewCategoryNamePlaceholder: "Category name...",
		NewCategoryButton:          "OK",
		TooltipAdd:                 "Add",
		TooltipEdit:                "Edit",
		TooltipDelete:              "Delete",
		TooltipAddCategory:         "Add Category",
	},
	languageFrench: {
		TablePageTemplate: TablePageTemplate{
			PageTemplate: PageTemplate{
				Title:   applicationName,
				Version: applicationVersion,
			},
			Search:       "Rechercher...",
			Category:     "Catégorie",
			Date:         "Date",
			Description:  "Description",
			Amount:       "Montant",
			CreationDate: "Créé le",
		},
		NewCategoryTitle:           "Nouvelle catégorie",
		NewCategoryName:            "Nom :",
		NewCategoryIcon:            "Icone :",
		NewCategoryNamePlaceholder: "Nom de la catégorie...",
		NewCategoryButton:          "OK",
		TooltipAdd:                 "Ajouter",
		TooltipEdit:                "Editer",
		TooltipDelete:              "Supprimer",
		TooltipAddCategory:         "Ajouter une catégorie",
	},
}

var settingsPagesByLanguage = map[Language]SettingsPageTemplate{
	languageEnglish: {
		PageTemplate: PageTemplate{
			Title:   applicationName,
			Version: applicationVersion,
		},
		Settings:              "Settings",
		SettingLanguage:       "Language:",
		AvailableLanguages:    availableLanguages,
		SettingServerPort:     "Port:",
		SettingServerDatabase: "Database:",
	},
	languageFrench: {
		PageTemplate: PageTemplate{
			Title:   applicationName,
			Version: applicationVersion,
		},
		Settings:              "Paramètres",
		SettingLanguage:       "Langue :",
		AvailableLanguages:    availableLanguages,
		SettingServerPort:     "Port :",
		SettingServerDatabase: "Base de données :",
	},
}
