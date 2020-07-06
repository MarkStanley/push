package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// PageData is a convenient struct to pass stuff through to the template
type PageData struct {
	Message string
}

// ShowPushButton displays a template with a Push button on it
func ShowPushButton(w http.ResponseWriter, r *http.Request) {
	var data PageData

	data.Message = "Bite me."

	summaryPage, err := template.ParseFiles("./templates/summary.gohtml")
	if err != nil {
		fmt.Println("Error parsing summary template: ", err)
		return
	}

	summaryPage.ExecuteTemplate(w, "summarypage", data)

}

// PushMessage pushes a message to the web
func PushMessage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "https://myway.thingitude-apps.com")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	//json.NewEncoder(w).Encode(sensors)
}
