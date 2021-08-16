package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type User struct {
	Name string
	ID string
	Password string
}

type Toggle struct {
	Name string
}

var Users []User
var features []string
var features2Toggles map[string]bool

func main() {
	fmt.Println("Starting Toggler")
	features = append(features, "addUser")
	features = append(features, "addFeatureToggle")
	features = append(features, "toggleFeature")
	features2Toggles = make(map[string]bool)
	handleRequests()
}

func handleRequests () {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/add_user", addUserHandler).Methods("POST")
	router.HandleFunc("/add_feature_toggle", addFeatureToggleHandler).Methods("POST")
	router.HandleFunc("/toggle_feature", toggleFeatureHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

/////
// HANDLERS
/////

func addUserHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: add user")
	toggle, ok := features2Toggles["addUser"]
	if ok == true && toggle == false {
		fmt.Println("Feature by the name of addUser is currently disabled")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newUser User
	json.Unmarshal(reqBody, &newUser)
	addUser(newUser)
}

func addFeatureToggleHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: add toggle")
	if features2Toggles["addFeatureToggle"] && features2Toggles["addFeatureToggle"] == false {
		fmt.Println("Feature by the name of addFeatureToggle is currently disabled")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newToggle Toggle
	json.Unmarshal(reqBody, &newToggle)
	addFeatureToggle(newToggle)
}

func toggleFeatureHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: toggle feature")
	if features2Toggles["toggleFeatureFeature"] && features2Toggles["toggleFeatureFeature"] == false {
		fmt.Println("Feature by the name of toggleFeatureFeature is currently disabled")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var toggle Toggle
	json.Unmarshal(reqBody, &toggle)
	toggleFeature(toggle)
}

/////
// FEATURES
/////

func addUser (newUser User) {
	Users = append(Users, newUser)
	fmt.Println (Users)
}

func addFeatureToggle (newToggle Toggle) {
	for _, f := range features {
		if f == newToggle.Name {
			fmt.Println("Adding new toggle for feature by the name of " + newToggle.Name)
			features2Toggles[f] = true
			return
		}
	}
	fmt.Println("No feature by the name of " + newToggle.Name)
}

func toggleFeature (toggle Toggle) {

	_, ok := features2Toggles[toggle.Name]

	if ok == true {

		if features2Toggles[toggle.Name] == true {
			fmt.Println(toggle.Name + " has been disabled")
			features2Toggles[toggle.Name] = false
		} else {
			fmt.Println(toggle.Name + " has been enabled")
			features2Toggles[toggle.Name] = true
		}

	} else {
		fmt.Println("No toggle for feature by the name of " + toggle.Name)
	}
	
}
