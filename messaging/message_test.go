package messaging

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tbalthazar/onesignal-go"
)

var _ = Describe("List Applications", func() {

	os.Setenv("USER_KEY", "NjY1M2ViODktNTIxMi00ODBmLTk4Y2YtMjE3YzE1Y2I0ZmJl")
	os.Setenv("APP_KEY", "MDdjMDJlY2ItNTVhOC00NDI1LTllOTEtOTdkNzgyNWUzYjIz")

	request, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListApp)
	handler.ServeHTTP(recorder, request)

	Describe("List Application", func() {
		Context("ListApp", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Verify http status code with invalid user key", func() {

	os.Setenv("USER_KEY", "//11NjY1M2ViODktNTIxMi00ODBmLTk4Y2YtMjE3YzE1Y2I0ZmJl")
	os.Setenv("APP_KEY", "MDdjMDJlY2ItNTVhOC00NDI1LTllOTEtOTdkNzgyNWUzYjIz")

	request, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListApp)
	handler.ServeHTTP(recorder, request)

	Describe("List Application", func() {
		Context("ListApp", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("OneSignal messaging with valid data", func() {
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = "aebc588d-f5a5-4668-8777-0debf272783a"
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{"7f3ec4c5-9aef-4eb0-8290-ad0d2eaf9c9a", "4b41af62-6bbb-4881-be2d-4f969d5be88d", "c94168ed-8f7f-480a-9a1e-a445f9e0ae42"}
	argumentData.ChromeWebIcon = "https://pbs.twimg.com/profile_images/1008644428964286464/gyBHbzaZ_400x400.jpg"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(argumentData)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusOk", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("OneSignal messaging with empty AppID", func() {
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = ""
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{"7f3ec4c5-9aef-4eb0-8290-ad0d2eaf9c9a", "4b41af62-6bbb-4881-be2d-4f969d5be88d", "c94168ed-8f7f-480a-9a1e-a445f9e0ae42"}
	argumentData.ChromeWebIcon = "https://pbs.twimg.com/profile_images/1008644428964286464/gyBHbzaZ_400x400.jpg"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(argumentData)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("OneSignal messaging with invalid AppID", func() {
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = "123"
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{"7f3ec4c5-9aef-4eb0-8290-ad0d2eaf9c9a", "4b41af62-6bbb-4881-be2d-4f969d5be88d", "c94168ed-8f7f-480a-9a1e-a445f9e0ae42"}
	argumentData.ChromeWebIcon = "https://pbs.twimg.com/profile_images/1008644428964286464/gyBHbzaZ_400x400.jpg"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(argumentData)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("OneSignal messaging with without Heading", func() {
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = "aebc588d-f5a5-4668-8777-0debf272783a"
	argumentData.Headings = map[string]string{"en": ""}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{"7f3ec4c5-9aef-4eb0-8290-ad0d2eaf9c9a", "4b41af62-6bbb-4881-be2d-4f969d5be88d", "c94168ed-8f7f-480a-9a1e-a445f9e0ae42"}
	argumentData.ChromeWebIcon = "https://pbs.twimg.com/profile_images/1008644428964286464/gyBHbzaZ_400x400.jpg"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(argumentData)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("OneSignal messaging without Content", func() {
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = "aebc588d-f5a5-4668-8777-0debf272783a"
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": ""}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{"7f3ec4c5-9aef-4eb0-8290-ad0d2eaf9c9a", "4b41af62-6bbb-4881-be2d-4f969d5be88d", "c94168ed-8f7f-480a-9a1e-a445f9e0ae42"}
	argumentData.ChromeWebIcon = "https://pbs.twimg.com/profile_images/1008644428964286464/gyBHbzaZ_400x400.jpg"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(argumentData)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("OneSignal messaging with invalid playersID", func() {
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = "aebc588d-f5a5-4668-8777-0debf272783a"
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{"c94168//ed-8f7f-480a-9a1e-a445f9e0ae42"}
	argumentData.ChromeWebIcon = "https://pbs.twimg.com/profile_images/1008644428964286464/gyBHbzaZ_400x400.jpg"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(argumentData)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("OneSignal messaging with empty playersID", func() {
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = "aebc588d-f5a5-4668-8777-0debf272783a"
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{""}
	argumentData.ChromeWebIcon = "https://pbs.twimg.com/profile_images/1008644428964286464/gyBHbzaZ_400x400.jpg"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(argumentData)

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})
