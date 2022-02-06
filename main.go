package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var appId string = ""
var apiKey string = ""

func sendNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	m, b := map[string]interface{}{
		"app_id": appId,
		"contents": map[string]string{
			"en": "Test Push SendNotification GOOOOOO",
		},
		"included_segments": []string{"All"},
		"content_available": true,
		"small_icon":        "ic_notification_icon",
		"data": map[string]string{
			"PushTitle": "CUSTOM NOTIFICATİON",
		},
	}, new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)

	url := "https://onesignal.com/api/v1/notifications"

	var basic = "Basic " + apiKey

	req, err := http.NewRequest("POST", url, b)

	// add authorization header to the req
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", basic)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))

}

type Data struct {
	Devices []string
}

func sendNotificationToDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var d Data
	err := json.NewDecoder(r.Body).Decode(&d)
	//var devices []string
	fmt.Println(d.Devices)

	m, b := map[string]interface{}{
		"app_id": appId,
		"contents": map[string]string{
			"en": "Test Push SendNotificationToDevice GOOOOOO",
		},
		"include_player_ids": d.Devices,
		"content_available":  true,
		"small_icon":         "ic_notification_icon",
		"data": map[string]string{
			"PushTitle": "CUSTOM NOTIFICATİON",
		},
	}, new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)

	url := "https://onesignal.com/api/v1/notifications"

	var basic = "Basic " + apiKey

	req, err := http.NewRequest("POST", url, b)

	// add authorization header to the req
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", basic)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/SendNotification", sendNotification).Methods("GET")
	r.HandleFunc("/api/SendNotificationToDevice", sendNotificationToDevice).Methods("POST")
	fmt.Printf("Port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
