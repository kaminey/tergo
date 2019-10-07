package fcm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type FCMService struct {
	ApiKey string
}

type MessagePayload struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

type NotificationPayload struct {
	To           string         `json:"to"`
	Notification MessagePayload `json:"notification"`
}

var client http.Client

func generateJSONPayload(token, title, body string) []byte {
	payload, _ := json.Marshal(NotificationPayload{
		To: token,
		Notification: MessagePayload{
			Body:  body,
			Title: title,
		},
	})
	return payload
}

func (service *FCMService) SendMessage(token, title, body string) {
	payload := generateJSONPayload(token, title, body)
	url := os.Getenv("FCM_URL")
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Authorization", "key="+service.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)
}
