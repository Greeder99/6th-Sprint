package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Conversation(stringOrMorse string) string {
	var formattext bool
	for _, i := range strings.ToLower(stringOrMorse) {
		// Провожу проверку на то, есть ли в переменной какие то другие символы кроме - и . Проверка по русскому алфавиту потому что в пакете morse мы работает именно с русским алфавитом
		if (i >= 1072 && i <= 1103) || i == 1105 || i == 34 || (i >= 39 && i <= 41) || i == 43 || i == 44 || (i >= 47 && i <= 58) || i == 61 || i == 63 || i == 64 {
			formattext = true
		} else {
			break
		}
	}
	if formattext {
		return morse.ToMorse(stringOrMorse)
	} else {
		return morse.ToText(stringOrMorse)
	}
}
