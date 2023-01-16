package utils

import (
	"log"
	"regexp"
	"time"
)

func IsValidAlphabet(s string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z ]*$`)
	return regex.MatchString(s)
}

func IsValidNumeric(s string) bool {
	regex, _ := regexp.Compile(`([0-9])`)
	return regex.MatchString(s)
}

func IsValidAlphaNumeric(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z0-9]*$)`)
	return regex.MatchString(s)
}

func IsValidAlphaNumericHyphen(s string) bool {
	regex, _ := regexp.Compile(`[a-zA-Z0-9-]+`)
	return regex.MatchString(s)
}

func FormattedTime(ts string) string {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Println(err)
		return ""
	}

	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}
