package file_utils

import (
	"bitbucket.org/zapr/go-utils/common_utils"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
)

/*
	Loads the config file into an object.
	On error, the program will exit.
	Returns the configuration object.
*/
func LoadConfiguration(configFilePath string, configObject interface{}) {
	// Check for errors
	if configFilePath == "" {
		log.Fatal("Please provide the path to config file.")
	}

	// Initialize
	var configuration interface{}

	// Read config file
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalln("Error while opening configuration file:", err)
	}
	defer common_utils.CloseStream(configFile)

	// Parse config file
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&configuration)
	if err != nil {
		log.Fatalln("Error while parsing the configuration file:", err)
	}

	// Parse generic object to desired object
	err = mapstructure.Decode(configuration, configObject)
	if err != nil {
		errorMsg := "Error while mapping config data onto provided config object. " + url
		log.Fatalln(errorMsg, err)
	}
}
