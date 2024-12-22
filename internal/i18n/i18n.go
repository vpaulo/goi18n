package i18n

type Translation struct {
	Language            string `json:"language"`
	LanguageSwitch      string `json:"language.switch"`
	LanguageSwitchLabel string `json:"language.switch.label"`
	PageTitle           string `json:"page.title"`
	PageGreeting        string `json:"page.greeting"`
}
