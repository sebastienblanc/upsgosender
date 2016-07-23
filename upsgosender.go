package upsgosender

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Settings needed to configure the sender
type Settings struct {
	URL           string
	ApplicationID string
	MasterSecret  string
}

// Sender ...
type Sender struct {
	Settings Settings
}

// Criteria ...
type Criteria struct {
	Alias      []string `json:"alias,omitempty"`
	DeviceType []string `json:"devices,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Variants   []string `json:"variants,omitempty"`
}

// Message represents the payload to be sent to UPS
type Message struct {
	Alert    string            `json:"alert,omitempty"`
	Sound    string            `json:"sound,omitempty"`
	Badge    string            `json:"badge,omitempty"`
	Priority string            `json:"priority,omitempty"`
	UserData map[string]string `json:"user-data,omitempty"`
	Windows  Windows           `json:"windows,omitempty"`
	Apns     Apns              `json:"apns,omitempty"`
}

// Windows ...s
type Windows struct {
	Type       string   `json:"type,omitempty"`
	Duration   string   `json:"duration,omitempty"`
	Badge      string   `json:"badge,omitempty"`
	TitleType  string   `json:"titleType,omitempty"`
	Images     []string `json:"images,omitempty"`
	TextFields []string `json:"textFields,omitempty"`
}

// Apns ...
type Apns struct {
	Title            string   `json:"title,omitempty"`
	Action           string   `json:"action,omitempty"`
	URLArgs          []string `json:"urlArgs,omitempty"`
	TitleLocKey      string   `json:"titleLocKey,omitempty"`
	TitleLocKeyArgs  []string `json:"titleLocKeyArgs,omitempty"`
	ActionCategory   string   `json:"actionCategory,omitempty"`
	ContentAvailable bool     `json:"contentAvailable,omitempty"`
}

// UnifiedMessage ...
type UnifiedMessage struct {
	Criteria Criteria `json:"criteria,omitempty"`
	Message  Message  `json:"message,omitempty"`
}

func (f UnifiedMessage) toJSON() []byte {
	result, _ := json.Marshal(f)
	return result
}

func (s Sender) send(message UnifiedMessage) *http.Response {
	client := &http.Client{}
	var jsonStr = []byte(string(message.toJSON()))
	req, err := http.NewRequest("POST", s.Settings.URL, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(s.Settings.ApplicationID, s.Settings.MasterSecret)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
