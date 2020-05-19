package slackutil

import (
	"bytes"
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"os"
)

// Post given text to a slack web-hook url. Also adds some process parameters to the message. Also alert if set true @here is added to slack message.
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
		log.Println("error in posting message to slack", err)
	}

	res, err := http.Post(webHookUrl, "application/json", bytes.NewReader(structBytes))
	if err != nil || res.StatusCode != 200 {
		log.Println("Error in posting message to slack", err)
	}
}
