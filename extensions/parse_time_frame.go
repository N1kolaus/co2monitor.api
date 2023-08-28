package extensions

import (
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

func ValidateTimeDuration(input string) time.Duration {
	defaultDuration := 6 * time.Hour

	// Prüfe, ob der String mindestens zwei Zeichen hat
	if len(input) < 2 {
		log.Infof("Ungültige Eingabe bei TimeDuration: %s", input)
		return defaultDuration
	}

	// Extrahiere die letzte Zeichen als Einheit (m, h, d)
	unit := input[len(input)-1]

	// Prüfe, ob die Einheit gültig ist
	if unit != 'm' && unit != 'h' && unit != 'd' {
		log.Infof("Ungültige Einheit bei TimeDuration: %s", input)
		return defaultDuration
	}

	// Extrahiere den numerischen Wert aus dem String
	valueStr := input[:len(input)-1]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Infof("Ungültiger numerischer Wert bei TimeDuration: %s", input)
		return defaultDuration
	}

	// Konvertiere den numerischen Wert in eine TimeDuration basierend auf der Einheit
	switch unit {
	case 'm':
		return time.Duration(value) * time.Minute
	case 'h':
		return time.Duration(value) * time.Hour
	case 'd':
		return time.Duration(value) * time.Hour * 24
	default:
		return defaultDuration
	}
}
