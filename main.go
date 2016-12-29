package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var index = template.Must(template.ParseFiles(
	"templates/index.html",
))

func init() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/v1/distance", distanceHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	index.ExecuteTemplate(w, "index.html", nil)
}

func distanceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if valid, dr := checkPostValidityAndParseForm(r); !valid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(marshalResponse(dr))
		log.Println("Invalid request")
		return
	}

	//todo:  MaxBytesReader

}

func checkPostValidityAndParseForm(r *http.Request) (bool, DistanceResponse) {

	if r.Method != "POST" {
		dr := getUnsuccessfulResponse("Only HTTP POST is supported")
		return false, dr
	}

	if r.Body == nil {
		dr := getUnsuccessfulResponse("HTTP POST body must contain a JSON object with keys 'source' and 'target'")
		return false, dr
	}

	decoder := json.NewDecoder(r.Body)

	var post DistancePost
	err := decoder.Decode(&post)

	if err != nil {
		message := fmt.Sprintf("Error decoding JSON http post, expected valid JSON body: %v", err)
		dr := getUnsuccessfulResponse(message)
		return false, dr
	}

	if post.Source == "" {
		dr := getUnsuccessfulResponse(`HTTP POST must have JSON object with key called 'source' containing a comma delimited list of source strings to be mapped`)
		return false, dr

	}

	if post.Target == "" {
		dr := getUnsuccessfulResponse(`HTTP POST must have JSON object with key called 'target' containing a comma delimited list of target strings to be mapped to`)
		return false, dr
	}

	return true, DistanceResponse{}

}

func marshalResponse(dr DistanceResponse) []byte {

	jData, err := json.Marshal(dr)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
	}

	return jData

}

func getUnsuccessfulResponse(message string) DistanceResponse {
	return DistanceResponse{Success: false, Message: message}
}
