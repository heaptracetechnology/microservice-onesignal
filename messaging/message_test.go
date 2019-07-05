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
	"log"
)

var (
	userKey          = os.Getenv("ONESIGNAL_USER_KEY")
	appKey           = os.Getenv("ONESIGNAL_APP_KEY")
	appID            = os.Getenv("ONESIGNAL_APP_ID")
	includePlayerID1 = os.Getenv("ONESIGNAL_PLAYER_ID1")
	includePlayerID2 = os.Getenv("ONESIGNAL_PLAYER_ID2")
)

var playerIdList []string

var _ = Describe("List Applications", func() {

	request, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
		log.Fatal(err)
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

var _ = Describe("Verify http status code with invalid user key", func() {

	os.Setenv("USER_KEY", ",mockuserkey")

	request, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
		log.Fatal(err)
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

	if includePlayerID1 != "" {
		playerIdList = append(playerIdList, includePlayerID1)
	}
	if includePlayerID2 != "" {
		playerIdList = append(playerIdList, includePlayerID2)
	}

	var argumentData onesignal.NotificationRequest
	argumentData.AppID = appID
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = playerIdList

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
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

	if includePlayerID1 != "" {
		playerIdList = append(playerIdList, includePlayerID1)
	}
	if includePlayerID2 != "" {
		playerIdList = append(playerIdList, includePlayerID2)
	}
	var argumentData onesignal.NotificationRequest
	argumentData.AppID = ""
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = playerIdList

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
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
	argumentData.IncludePlayerIDs = playerIdList

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
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
	argumentData.AppID = appID
	argumentData.Headings = map[string]string{"en": ""}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = playerIdList

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
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
	argumentData.AppID = appID
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": ""}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = playerIdList

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
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
	argumentData.AppID = appID
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{"mockplayerid"}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
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
	argumentData.AppID = appID
	argumentData.Headings = map[string]string{"en": "Test Notification"}
	argumentData.Contents = map[string]string{"en": "Demo testing"}
	argumentData.IsAnyWeb = true
	argumentData.IncludePlayerIDs = []string{""}

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
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

var _ = Describe("OneSignal messaging with Invalid data", func() {

	argumentData := []byte(`{"appid":false}`)

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
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
