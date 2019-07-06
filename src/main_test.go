package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

func TestUnMarshalSlackVerification(t *testing.T) {
	var verificationR SlackRequest
	r := ioutil.NopCloser(strings.NewReader("{\"token\":\"1234\", \"challenge\":\"qwerty\", \"type\":\"url_verification\", \"team_id\": \"777\"}"))

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	r.Close()

	err := json.Unmarshal([]byte(buf.String()), &verificationR)
	if err != nil {
		t.Errorf("Can't unmarshal json %s", err)
	}
}
