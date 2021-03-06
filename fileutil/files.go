package fileutil

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/zapr-oss/go-utils/common"
	"log"
	"os"
)


// Loads the config file into an object. On error, the program will exit. Returns the configuration object.s
func LoadConfiguration(configFilePath string, configObject interface{}) error {
	// Check for errors
	if configFilePath == "" {
		log.Println("Please provide the path to config file.")
		return errors.New("ConfigPathNotProvided")
	}

	// Initialize
	var configuration interface{}

	// Read config file
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Println("Error while opening configuration file:", err)
		return err
	}
	defer commonutil.CloseStream(configFile)

	// Parse config file
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&configuration)
	if err != nil {
		log.Println("Error while parsing the configuration file:", err)
		return err
	}

	// Parse generic object to desired object
	err = mapstructure.Decode(configuration, configObject)
	if err != nil {
		log.Println("Error while mapping config data onto provided config object.", err)
		return err
	}

	return nil
}
