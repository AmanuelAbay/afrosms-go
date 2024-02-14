/**
 * Author: Amanuel Abay
 * File: afro-sms_test.go
 */

package afro

import (
	"testing"
)

func TestGetRetRequest(t *testing.T) {
	expected := Request{BaseURL: "https://api.afromessage.com/api/send"}
	token := "AFRO_SMS_TOKEN"
	endpoint := "/api/send"
	host := "https://api.afromessage.com"
	if got := GetRequest(token, endpoint, host); got.BaseURL != expected.BaseURL {
		t.Errorf("GetRequest(%s, %s, %s) = %s , didn't return %s", token, endpoint, host, got.BaseURL, expected.BaseURL)
	}
}

func TestSender(t *testing.T) {
	request := Request{}
	expected := "AFRO_SMS_SENDER_NAME"
	senderName := "AFRO_SMS_SENDER_NAME"
	request.Sender(senderName)
	if got := request.QueryParams["sender"]; got != expected {
		t.Errorf("request.Sender(%s) = %s, didn't return %s", senderName, got, expected)
	}
}

func TestTo(t *testing.T) {
	request := Request{}
	expected := "+2519xxxxxxxxx"
	phoneNumber := "+2519xxxxxxxxx"
	message := "Afro Test Message"
	request.To(phoneNumber, message)
	if got := request.QueryParams["to"]; got != expected {
		t.Errorf("request.To(%s) = %s, didn't return %s", phoneNumber, got, expected)
	}
}

func TestCallBackUrl(t *testing.T) {
	request := Request{}
	expected := "http://localhost:3000/callback"
	Url := "http://localhost:3000/callback"
	request.CallBackUrl(Url)
	if got := request.QueryParams["callback"]; got != expected {
		t.Errorf("request.CallBackUrl(%s) = %s, didn't return %s", Url, got, expected)
	}
}

func TestApi(t *testing.T) {
	expected := Response{}
	SenderName := "AFRO_SMS_SENDER_NAME"
	PhoneNumber := "+2519xxxxxxxxx"
	Message := "afro-sms test text message"

	request := GetRequest("AFRO_SMS_TOKEN", "/api/send", "https://api.afromessage.com")
	request.Method = "GET"
	request.Sender(SenderName)
	request.To(PhoneNumber, Message)
	if got, _ := Api(request); got != &expected {
		t.Errorf("Api(%s) =, didn't work as expected", request)
	}
}
