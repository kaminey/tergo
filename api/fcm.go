package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"tergo/service/fcm"
)

type FCMRequestData struct {
	Tokens []string `json:"tokens`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
}

type FCMResponseData struct {
	Status string `json:"status"`
}

func FCMHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var data FCMRequestData
	json.Unmarshal(body, &data)
	api_key := os.Getenv("FCM_API_KEY")
	fcmService := fcm.FCMService{api_key}
	for _, token := range data.Tokens {
		go fcmService.SendMessage(token, data.Title, data.Body)
	}
	w.Header().Set("Content-Type", "application/json")
	responseData := FCMResponseData{"queued"}
	responseDataJson, _ := json.Marshal(responseData)
	w.Write(responseDataJson)
}
