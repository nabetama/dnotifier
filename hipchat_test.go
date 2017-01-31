package dnotifier

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	authToken = "AuthToken"
	roomID    = "111111"
	mux       *http.ServeMux
	server    *httptest.Server
	hc        *HipChat
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	hc = NewHipChat("111111", "AuthToken")
}

func tearDown() {
	server.Close()
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %s want %s", header, got, want)
	}
}

func TestNewHipChat(t *testing.T) {
	hc := NewHipChat(roomID, authToken)
	if hc.Token != authToken {
		t.Errorf("NewClient Token %s, want %s", hc.Token, authToken)
	}
}

func TestNewMessageData(t *testing.T) {
	messageFmt := "text"
	message := "testmessage"
	notify := true
	md := NewMessageData(messageFmt, message, notify)

	if md.Color != "purple" {
		t.Errorf("MessageData Color %s, want %s", md.Color, "purple")
	}

	if md.MessageFormat != messageFmt {
		t.Errorf("MessageData MessageFormat %s, want %s", md.MessageFormat, messageFmt)
	}

	if md.Message != message {
		t.Errorf("MessageData Message %s, want %s", md.Message, message)
	}

	if md.Notification != notify {
		t.Errorf("MessageData Notification %s, want %t", md.Notification, notify)
	}
}

func TestHookURL(t *testing.T) {
	hc := NewHipChat(roomID, authToken)
	hookURL := "https://api.hipchat.com/v2/room/%s/notification?auth_token=%s"
	validHookURL := fmt.Sprintf(hookURL, roomID, authToken)
	if hc.HookURL() != validHookURL {
		t.Errorf("NewClient hookURL %s, want %s", hc.HookURL, validHookURL)
	}
}
