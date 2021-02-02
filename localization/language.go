package localization

// Language is a type that is used to localize strings.
type Language int

// Here are the different available languages.
const (
	LanguageEnglish Language = iota + 1
	LanguageFrench
)

// To satisfy the fmt.Stringer interface.
func (l Language) String() string {
	switch l {
	case LanguageEnglish:
		return "English"
	case LanguageFrench:
		return "Français"
	}
	return "???"
}
