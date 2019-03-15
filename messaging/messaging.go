package messaging

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/heaptracetechnology/microservice-onesignal/result"
	"github.com/tbalthazar/onesignal-go"
)

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

//List applications
func ListApp(responseWriter http.ResponseWriter, request *http.Request) {

	appKey := os.Getenv("APP_KEY")
	userKey := os.Getenv("USER_KEY")
	client := onesignal.NewClient(nil)
	client.AppKey = appKey
	client.UserKey = userKey
	apps, _, err := client.Apps.List()
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}

	bytes, err := json.Marshal(apps)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//Send push notification
func SendMessage(responseWriter http.ResponseWriter, request *http.Request) {

	client := onesignal.NewClient(nil)

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}

	defer request.Body.Close()
	var argumentData onesignal.NotificationRequest

	err = json.Unmarshal(body, &argumentData)

	notificationReq := onesignal.NotificationRequest(argumentData)

	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	_, _, errr := client.Notifications.Create(&notificationReq)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}
	message := Message{"true", "Notification sent", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}
