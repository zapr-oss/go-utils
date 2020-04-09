package slack_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Post(webHookUrl, text string, alert bool) {
	hostName, _ := os.Hostname()
	message := fmt.Sprintf("*Host*: %v,\nPID: %v,\n*Message*: %v\n", hostName, os.Getpid(), text)
	log.Println(message)

	if alert {
		message = "<!here>\n" + message
	}
	data := struct {
		Text string `json:"text"`
	}{message}

	structBytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Error in posting message to slack")
	}

	res, err := http.Post(webHookUrl, "application/json", bytes.NewReader(structBytes))
	if err != nil || res.StatusCode != 200 {
		log.Println("Error in posting message to slack")
	}
}
