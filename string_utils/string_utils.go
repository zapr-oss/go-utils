package string_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func GetStructAsString(obj interface{}) interface{} {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		log.Println("Error in converting object to string.", err)
		return nil
	}
	if string(objBytes) == "null" {
		return nil
	}

	return string(objBytes)
}

func RemoveNonAlphaNumeric(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Println(fmt.Sprintf("Error applying Non Alphanumeric regex to string, Err: %v", str), err)
		return str
	}
	return reg.ReplaceAllString(str, "")
}

func RemoveWords(str string, wordsToRemove []string) string {
	for _, word := range wordsToRemove {
		str = strings.Replace(str, word, "", -1)
	}

	return strings.TrimSpace(str)
}

func CaseInsensitiveStringEquals(str1, str2 string) bool {
	return strings.ToLower(str1) == strings.ToLower(str2)
}
