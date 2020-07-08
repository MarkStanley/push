package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

// PageData is a convenient struct to pass stuff through to the template
type PageData struct {
	Message string
	Token   string
}

// FCMServerURL is the URL for firebase cloud mesasging
var FCMServerURL = "https://fcm.googleapis.com/fcm/send"

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

	body := strings.NewReader(`
		{
			"notification": {
				"title": "DILU",
				"body": "DILU is awesome",
				"click_action": "http://localhost:3000/",
				"icon": "http://url-to-an-icon/icon.png"
			},
			"to": "fomlE-yp9cqsgD6xWUhqKK:APA91bHnvDcDsBkH73FiRa5L1ZvcsYX0ra326pe9Df0ruGB9pdSnEp7xP4-2wn53XU_wTDX0JoR8v07yReThU893ZuLb5D3WQomCV-9aQEzdIGTbse4YEkwAmZC6RS2cqGfed0Kw2ic0"
		}
		`)
	req, err := http.NewRequest("POST", FCMServerURL, body)
	if err != nil {
		log.Printf("Failed to create new HTTP request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("key=%s", FCMKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	log.Printf("And the response was: %v\n", resp.Body)

}
