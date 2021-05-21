package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Butler Order Payload Struct.
type ButlerOrder struct {
	Summary     string
	Description string
	Assignee    string
}

func butlerHandler(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")

	fmt.Printf("Butler Action is %s", action)
	if action == "create" {
		message := createButlerAction()
		responseWithJSON(w, 200, message)
	}

	responseWithNoJSON(w, 400)
}

func createButlerAction() interface{} {
	message := ButlerOrder{
		Summary:     "Monkey Cert Will Expire Soon",
		Description: "Monkey certificate is going to expire in 10 days so please move your ass now",
		Assignee:    "Squad Infra",
	}

	return message
}

func responseWithNoJSON(w http.ResponseWriter, code int) {

	w.WriteHeader(code)
	w.Write(nil)
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	fmt.Println(string(response))

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)

	return nil
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/HttpJiraButler", butlerHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
