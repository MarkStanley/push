package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// PageData is a convenient struct to pass stuff through to the template
type PageData struct {
	Message string
	Token   string
}

// FCMMessage - defines the messaging struct for Firebase Cloud Messaging
type FCMMessage struct {
	To              string      `json:"to,omitempty"`
	RegistrationIDs []string    `json:"registration_ids,omitempty"`
	Data            interface{} `json:"data,omitempty"`
}

// FCMTokenMessage - the struct for a token from fcm
type FCMTokenMessage struct {
	Token string `json:"token" binding:"required"`
}

// FCMServerURL is the URL for firebase cloud mesasging
var FCMServerURL = "https://fcm.googleapis.com/fcm/send"

// FCMTokenMap - not clear what this is yet
var FCMTokenMap = map[string]bool{}

// IDs is some kind of array of registered users on fcm who I am about to message
var IDs []string

// FCMKey is the server key for Firebase cloud messaging I guess
var FCMKey = "AAAAiC1L8U8:APA91bFg7ZViG6kWiEddLna-SuiDCciF0Yx4mx5x1gH1nWhoOzVR9b3gJUYA-wU3THVv50U_6beSQ5-tGmKeOdh-xEXmbfAfeGo48s0_6z5SzKi43LXtIzymC7Py5nmy8CF3OTDurHcU"

// ShowPushButton displays a template with a Push button on it
func ShowPushButton(w http.ResponseWriter, r *http.Request) {
	var data PageData

	data.Message = "Bite me."
	data.Token = "fomlE-yp9cqsgD6xWUhqKK:APA91bHnvDcDsBkH73FiRa5L1ZvcsYX0ra326pe9Df0ruGB9pdSnEp7xP4-2wn53XU_wTDX0JoR8v07yReThU893ZuLb5D3WQomCV-9aQEzdIGTbse4YEkwAmZC6RS2cqGfed0Kw2ic0"

	page, err := template.ParseFiles("./templates/page.gohtml")
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}

	page.ExecuteTemplate(w, "page", data)

}

// PushMessage pushes a message to the web
func PushMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Button pushed")
	var tm FCMTokenMessage
	tm.Token = r.FormValue("token")

	if _, ok := FCMTokenMap[tm.Token]; !ok {
		FCMTokenMap[tm.Token] = true
	}
	IDs = append(IDs, tm.Token)
	message := r.FormValue("message")
	m := FCMMessage{
		RegistrationIDs: IDs,
		Data:            map[string]string{"message": message},
	}

	jd, err := json.Marshal(&m)
	if err != nil {
		log.Printf("Failed to marshal JSON: %s", err.Error())
		return
	}

	log.Printf("FCM Message: %s", string(jd))
	log.Printf("FCM Token: %s", tm.Token)

	//Lets create the request
	req, err := http.NewRequest("POST", FCMServerURL, bytes.NewReader(jd))
	if err != nil {
		log.Printf("Failed to create new HTTP request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("key=%s", FCMKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	fmt.Printf("And the response was: %v", resp)

	w.Header().Set("Access-Control-Allow-Origin", "https://myway.thingitude-apps.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	//json.NewEncoder(w).Encode(sensors)
}
