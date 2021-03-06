package localization

// MonthsByLanguage returns the names of the months
// for a given language.
func MonthsByLanguage(language Language) []string {
	return monthsByLanguage[language]
}

// DateFormatsByLanguage returns the standard date format
// for a given language.
func DateFormatsByLanguage(language Language) string {
	return dateFormatsByLanguage[language]
}

// ColumnsByLanguage returns the columns header
// for a given language.
func ColumnsByLanguage(language Language) TableColumns {
	return columnsByLanguage[language]
}

// /!\ TODO - Move everything out in an external config file.

var monthsByLanguage = map[Language][]string{
	LanguageEnglish: {
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
	LanguageFrench: {
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
	LanguageEnglish: "2006-01-02",
	LanguageFrench:  "02/01/2006",
}

var columnsByLanguage = map[Language]TableColumns{
	LanguageEnglish: {
		Category:     "Category",
		Date:         "Date",
		Description:  "Description",
		Amount:       "Amount",
		CreationDate: "Created at",
	},
	LanguageFrench: {
		Category:     "Catégorie",
		Date:         "Date",
		Description:  "Description",
		Amount:       "Montant",
		CreationDate: "Créé le",
	},
}

var homePagesByLanguage = map[Language]HomePageTemplate{
	LanguageEnglish: {
		TablePageTemplate: TablePageTemplate{
			Search:       "Search...",
			Category:     "Category",
			Date:         "Date",
			Description:  "Description",
			Amount:       "Amount",
			CreationDate: "Created at",
			TableColumns: ColumnsByLanguage(LanguageEnglish),
		},
		Total:                "Total:",
		Download:             "Download",
		EditMode:             "Edit mode",
		NewOperationTitle:    "New operation",
		EditOperationTitle:   "Edit operation",
		EditOperationTooltip: "Edit operation",
		AmountType:           "Expense",
	},
	LanguageFrench: {
		TablePageTemplate: TablePageTemplate{
			Search:       "Rechercher...",
			Category:     "Catégorie",
			Date:         "Date",
			Description:  "Description",
			Amount:       "Montant",
			CreationDate: "Créé le",
			TableColumns: ColumnsByLanguage(LanguageFrench),
		},
		Total:                "Total :",
		Download:             "Télécharger",
		EditMode:             "Mode édition",
		NewOperationTitle:    "Nouvelle opération",
		EditOperationTitle:   "Modifier une opération",
		EditOperationTooltip: "Modifier une opération",
		AmountType:           "Dépense",
	},
}

var editPagesByLanguage = map[Language]EditPageTemplate{
	LanguageEnglish: {
		TablePageTemplate: TablePageTemplate{
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
	LanguageFrench: {
		TablePageTemplate: TablePageTemplate{
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
	LanguageEnglish: {
		Settings:           "Settings",
		SettingLanguage:    "Language:",
		AvailableLanguages: availableLanguages,
		SettingServerPort:  "Port:",
	},
	LanguageFrench: {
		Settings:           "Paramètres",
		SettingLanguage:    "Langue :",
		AvailableLanguages: availableLanguages,
		SettingServerPort:  "Port :",
	},
}
