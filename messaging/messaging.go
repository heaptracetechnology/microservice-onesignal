package messaging

import (
	"github.com/tbalthazar/onesignal-go"
	"os"
	"github.com/heaptracetechnology/microservice-onesignal/result"
	"net/http"
	"encoding/json"
	"io/ioutil"
)
type Message struct {
    Success 			string 			`json:"success"`
    Message 			string 			`json:"message"`
	  StatusCode 	  int 	 			`json:"statuscode"`
}
func ListApp(responseWriter http.ResponseWriter, request *http.Request){

	appKey :=os.Getenv("APP_KEY")
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


func SendMessage(responseWriter http.ResponseWriter, request *http.Request){

	client := onesignal.NewClient(nil)

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	defer request.Body.Close()
	var argumentData onesignal.NotificationRequest

	err = json.Unmarshal(body, &argumentData)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	notificationReq := onesignal.NotificationRequest(argumentData)

	_, _ , errr := client.Notifications.Create(&notificationReq)
	if err != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}
	message := Message{"true","Notification sent",http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}