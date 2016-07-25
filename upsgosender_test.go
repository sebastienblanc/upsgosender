package upsgosender

import (
	"bufio"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

//Test the basis marshalling of
func TestToJSON(t *testing.T) {
	dat, err := os.Open("test.txt")
	if err != nil {
		//panic(err)
	}
	//criteria := &Criteria{Aliases: []string{"seb", "bob"}}
	message := &Message{Alert: "hello"}
	unifiedMessage := &UnifiedMessage{Message: *message}

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		line := scanner.Text()
		if line != string(unifiedMessage.toJSON()) {
			t.Log(string(unifiedMessage.toJSON()))
			t.Fail()
		}
	}
}

func TestSend(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://aerogearpushpush-sblanc.rhcloud.com/ag-push/rest/sender",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, nil)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
	criteria := &Criteria{Alias: []string{"seb", "bob"}}
	message := &Message{Alert: "hello from #golang sender"}
	unifiedMessage := &UnifiedMessage{Message: *message, Criteria: *criteria}
	settings := &Settings{
		URL:           "https://aerogearpushpush-sblanc.rhcloud.com/ag-push",
		ApplicationID: "58f87fb7-829c-4c6f-a0eb-326d3017a94c",
		MasterSecret:  "3366736b-d52c-4115-87d3-c08095e87955"}
	sender := NewSender(*settings)
	if sender.send(*unifiedMessage).StatusCode != 200 {
		t.Fail()
	}
}
