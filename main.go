package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"stringdistance/distance"
)

var index = template.Must(template.ParseFiles(
	"templates/index.html", "templates/api.html",
))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/v1/distance", distanceHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/sitemap.xml", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	index.ExecuteTemplate(w, "index.html", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	index.ExecuteTemplate(w, "api.html", nil)
}

func distanceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if err := isRequestValid(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		dr := getUnsuccessfulResponse(err.Error())
		w.Write(marshalResponse(dr))
		log.Println("Invalid request")
		return
	}

	distanceRequest, err := decodeDistanceRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		dr := getUnsuccessfulResponse(err.Error())
		w.Write(marshalResponse(dr))
		log.Println("Invalid request")
		return
	}

	distanceResponse := distance.GetDistanceResponse(distanceRequest)

	w.WriteHeader(http.StatusOK)
	w.Write(marshalResponse(distanceResponse))

}

func isRequestValid(r *http.Request) error {

	if r.Method != "POST" {
		return fmt.Errorf("Only HTTP POST is supported")
	}

	if r.Body == nil {
		return fmt.Errorf("HTTP POST body must contain a JSON object with keys 'source' and 'target'")
	}

	return nil

}

func decodeDistanceRequest(r *http.Request) (distance.DistanceRequest, error) {

	decoder := json.NewDecoder(r.Body)

	var distanceRequest distance.DistanceRequest
	err := decoder.Decode(&distanceRequest)

	if err != nil {
		return distanceRequest, fmt.Errorf("Error decoding JSON http post, expected valid JSON body: %v", err)
	}

	if distanceRequest.Source == "" {
		return distanceRequest, fmt.Errorf(`HTTP POST must have JSON object with key called 'source' containing a comma delimited list of source strings to be mapped`)
	}

	if distanceRequest.Target == "" {
		return distanceRequest, fmt.Errorf(`HTTP POST must have JSON object with key called 'target' containing a comma delimited list of target strings to be mapped to`)
	}

	return distanceRequest, nil
}

func marshalResponse(dr distance.DistanceResponse) []byte {

	jData, err := json.Marshal(dr)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
	}

	return jData

}

func getUnsuccessfulResponse(message string) distance.DistanceResponse {
	return distance.DistanceResponse{Success: false, Message: message}
}
