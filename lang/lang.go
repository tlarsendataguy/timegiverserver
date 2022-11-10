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

func ParseLang(value string) Lang {
	if len(value) < 2 {
		return defaultLang
	}

	options := strings.Split(value, `,`)
	for _, option := range options {
		if len(option) < 2 {
			continue
		}
		code := strings.ToLower(option[0:2])
		switch code {
		case `de`:
			return DE
		case `en`:
			return EN
		case `es`:
			return ES
		case `fr`:
			return FR
		case `hi`:
			return HI
		case `ja`:
			return JA
		case `pt`:
			return PT
		case `ru`:
			return RU
		case `zh`:
			return ZH
		}
	}
	return defaultLang
}
