package stringutil

import (
	"encoding/json"
	"log"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Converts a go struct to string.
func GetStructAsString(obj interface{}) string {
	if obj == nil {
		return ""
	}

	objBytes, err := json.Marshal(obj)

	if err != nil {
		log.Println("Error in converting object to string.", err)
		return ""
	}
	if string(objBytes) == "null" {
		return ""
	}

	return string(objBytes)
}

// TypeCast string as interface.
func GetStringAsInterface(str string) interface{} {
	if str == "" {
		return nil
	}

	return str
}

// TypeCast struct as interface.
func GetStructAsInterface(obj interface{}) interface{} {

	if obj == nil {
		return nil
	}

	objBytes, err := json.Marshal(obj)

	if err != nil {
		log.Println("error in marshalling object")
		return nil
	}

	if string(objBytes) == "null" || string(objBytes) == "[]" || string(objBytes) == "{}" {
		return nil
	}

	return string(objBytes)
}

// Removes all non-alphanumeric characters.
func RemoveNonAlphaNumeric(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Println(fmt.Sprintf("error applying Non Alphanumeric regex to string, Err: %v", str), err)
		return str
	}
	return reg.ReplaceAllString(str, "")
}

// Removes given words from the provided text.
func RemoveWords(str string, wordsToRemove []string) string {
	for _, word := range wordsToRemove {
		str = strings.Replace(str, word, "", -1)
	}

	return strings.TrimSpace(str)
}

func CaseInsensitiveStringEquals(str1, str2 string) bool {
	return strings.ToLower(str1) == strings.ToLower(str2)
}

func CaseInsensitiveContains(str1, str2 string) bool {
	lowerStr1 := strings.ToLower(str1)
	lowerStr2 := strings.ToLower(str2)
	return strings.Contains(lowerStr1, lowerStr2)
}

// Get pointer to string.
func GetStringPtr(str string) *string {
	return &str
}

// Converts string to int pointer.
func GetStringAsIntPtr(str string) *int {
	yr, err := strconv.Atoi(str)

	if err != nil || yr <= 0 {
		return nil
	}

	return &yr
}