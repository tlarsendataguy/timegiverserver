package lang

import "strings"

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

func (l Lang) String() string {
	return []string{`DE`, `EN`, `ES`, `FR`, `HI`, `JA`, `PT`, `RU`, `ZH`}[l]
}

func ParseLang(value string) (Lang, error) {
	if len(value) < 2 {
		return defaultLang, nil
	}
	code := strings.ToLower(value[0:2])
	switch code {
	case `de`:
		return DE, nil
	case `en`:
		return EN, nil
	case `es`:
		return ES, nil
	case `fr`:
		return FR, nil
	case `hi`:
		return HI, nil
	case `ja`:
		return JA, nil
	case `pt`:
		return PT, nil
	case `ru`:
		return RU, nil
	case `zh`:
		return ZH, nil
	default:
		return defaultLang, nil
	}
}
