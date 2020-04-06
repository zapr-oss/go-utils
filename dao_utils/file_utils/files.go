package file_utils

import (
	"bitbucket.org/zapr/go-utils/common_utils"
	"encoding/json"
	"log"
	"os"
)

/*
	Loads the config file into an object.
	On error, the program will exit.
	Returns the configuration object.
*/
func LoadConfiguration(configFilePath string) interface{} {
	// Initialize
	var configuration interface{}

	// Read config file
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Error while opening configuration file:", err)
	}
	defer common_utils.CloseStream(configFile)

	// Parse config file
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&configuration)
	if err != nil {
		log.Fatal("Error while parsing the configuration file:", err)
	}

	return configuration
}
