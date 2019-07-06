package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Bot is the main goal of this project
type Bot struct {
	questionDict [][]string
}

func (b Bot) resolveReleaseID(q *string) string {
	pos := strings.LastIndex(*q, "\n")

	if pos < 0 {
		pos = strings.LastIndex(*q, " ")
	}

	if pos < 0 {
		return ""
	}

	result := *q

	return strings.Trim(result[pos+1:], " ?")
}

func (b Bot) answer(q string) string {
	for _, dic := range b.questionDict {
		i := 0
		l := len(dic)

		for _, word := range dic {
			if strings.Contains(q, word) {
				i++
			}
		}

		if i == l {
			return b.resolveReleaseID(&q)
		}
	}

	return "No te entiendo"
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", DoHealthCheck).Methods("GET")
	router.HandleFunc("/", Chat).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

type SlackRequest struct {
	Token      string
	TeamI_id   string
	Api_app_id string
	Type       string
	Challenge  string
	Event      struct {
		Text string
		Type string
	}
}

// Return a 200 with a verification token
func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Chat(w http.ResponseWriter, r *http.Request) {
	handleEvent(w, r)
	w.WriteHeader(http.StatusOK)
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	var req SlackRequest

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	err := json.Unmarshal([]byte(buf.String()), &req)
	if err != nil {
		panic(err)
	}

	switch req.Type {
	case "event_callback":
		b := Bot{
			questionDict: [][]string{
				[]string{"podéis decir", "tenant"},
				[]string{"tell us", "tenant"},
				[]string{"dime", "tenant"},
			},
		}

		fmt.Fprintf(w, b.answer("hey\n me podéis decir el tenant?\n 1234"))
		break
	case "url_verification":
		fmt.Fprintf(w, "%s", req.Challenge)
	}
}
