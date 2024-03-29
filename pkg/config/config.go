package config

import (
	"log"
	"strconv"

	"gopkg.in/ini.v1"

	"github.com/vpenando/piggy/pkg/localization"
)

const (
	// app infos
	applicationName    = "PiggyBox"
	applicationVersion = "v0.4.0"

	// config infos
	configFile             = "config.ini"
	serverConfigName       = "server"
	localizationConfigName = "localization"
)

// Default config.
var (
	// Server side config
	ServerPort     = "8081"
	ServerDatabase = "piggy.db"

	// Localization config
	CurrentLanguage = localization.LanguageEnglish
)

func ReadConfig() {
	config, err := ini.Load(configFile)
	if err != nil {
		log.Printf("Failed to open config file '%s': %s.\n", configFile, err)
		log.Print("Using default values.\n")
		return
	}
	if serverConfig := tryReadSection(config, serverConfigName); serverConfig != nil {
		readServerConfig(serverConfig)
	} else {
		logNilSection(serverConfigName)
	}
	if localizationConfig := tryReadSection(config, localizationConfigName); localizationConfig != nil {
		readLocalizationConfig(localizationConfig)
	} else {
		logNilSection(localizationConfigName)
	}
}

func tryReadSection(config *ini.File, sectionName string) *ini.Section {
	if config == nil {
		panic("config == nil")
	}
	section := config.Section(sectionName)
	return section
}

func readServerConfig(section *ini.Section) {
	if databaseKey := section.Key("database"); databaseKey.String() != "" {
		log.Printf("Read database '%s' from config file.", databaseKey.String())
		ServerDatabase = databaseKey.String()
	} else {
		logKeyNotFound("database")
	}
	if portKey := section.Key("port"); portKey.String() != "" {
		if _, err := strconv.Atoi(portKey.String()); err != nil {
			logNotSupportedValue("port", portKey.String())
			return
		}
		log.Printf("Read port '%s' from config file.", portKey.String())
		ServerPort = portKey.String()
	} else {
		logKeyNotFound("port")
	}
}

var languages = map[string]localization.Language{
	"en": localization.LanguageEnglish,
	"fr": localization.LanguageFrench,
}

func readLocalizationConfig(section *ini.Section) {
	if languageKey := section.Key("language"); languageKey.String() != "" {
		if languageValue, ok := languages[languageKey.String()]; ok {
			log.Printf("Read language '%s' from config file.", languageKey.String())
			CurrentLanguage = languageValue
		} else {
			logNotSupportedValue("language", languageKey.String())
		}
	} else {
		logKeyNotFound("language")
	}
}

func logNilSection(sectionName string) {
	log.Printf("Failed to open '%s' section. Using default values.\n", sectionName)
}

func logKeyNotFound(key string) {
	log.Printf("Failed to open key '%s'. Using default value.\n", key)
}

func logNotSupportedValue(key, value string) {
	log.Printf("Value '%s' not supported for key '%s'. Using default value.\n", value, key)
}
