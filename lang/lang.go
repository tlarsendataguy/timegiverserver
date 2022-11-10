package lang

type Lang int

const (
	DE Lang = iota
	EN
	ES
	FR
	HI
	JA
	PT
	RU
	ZH
)

const defaultLang = EN

func ParseLang(value string) (Lang, error) {
	return defaultLang, nil
}
