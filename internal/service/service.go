package service

import (
	"errors"
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertText(text string) (string, error) {
	if text == "" {
		err := errors.New("пустая строка")
		log.Println("ошибка конвертирования текста:", err)
		return "", err
	}

	clean := strings.ReplaceAll(text, " ", "")
	clean = strings.ReplaceAll(clean, "\n", "")
	clean = strings.ReplaceAll(clean, "\r", "")
	clean = strings.ReplaceAll(clean, ".", "")
	clean = strings.ReplaceAll(clean, "-", "")

	if clean == "" {
		return morse.ToText(text), nil
	}

	return morse.ToMorse(text), nil
}
