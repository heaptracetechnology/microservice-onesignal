package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"github.com/tbalthazar/onesignal-go"
	"net/http/httptest"
	"os"
)

var _ = Describe("List Applications", func() {


	os.Setenv("USER_KEY", "NjY1M2ViODktNTIxMi00ODBmLTk4Y2YtMjE3YzE1Y2I0ZmJl")
	os.Setenv("APP_KEY", "MDdjMDJlY2ItNTVhOC00NDI1LTllOTEtOTdkNzgyNWUzYjIz")

	request, err := http.NewRequest("GET", "/send",nil)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListApp)
	handler.ServeHTTP(recorder, request)

	Describe("List Application", func() {
		Context("ListApp", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})



var _ = Describe("OneSignal messaging", func() {
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = "aebc588d-f5a5-4668-8777-0debf272783a"
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents =  map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb =  true
	argumentData.IncludePlayerIDs = []string{"7f3ec4c5-9aef-4eb0-8290-ad0d2eaf9c9a","4b41af62-6bbb-4881-be2d-4f969d5be88d","c94168ed-8f7f-480a-9a1e-a445f9e0ae42"}
	argumentData.SmallIcon = "/my-icon.png"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(argumentData)

	// os.Setenv("USER_KEY", "NjY1M2ViODktNTIxMi00ODBmLTk4Y2YtMjE3YzE1Y2I0ZmJl")
	// os.Setenv("APP_KEY", "MDdjMDJlY2ItNTVhOC00NDI1LTllOTEtOTdkNzgyNWUzYjIz")

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})




