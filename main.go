package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FProject struct {
	ProjectId string   `json:"projectid"`
	Title     string   `json:"title"`
	Branch    string   `json:"branch"`
	Students  *Student `json:"student"`
}

type Student struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

var fproject []FProject

// getallProjects Info
func allProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fproject)
}

// get Single Candidate info
func getSingleProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, project := range fproject {
		if project.ProjectId == params["id"] {
			json.NewEncoder(w).Encode(project)
			return
		}
	}

	json.NewEncoder(w).Encode("No Project Found with Requested Id")
	return
}

// add new project
func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")

	var project FProject
	json.NewDecoder(r.Body).Decode(&project)

	project.ProjectId = strconv.Itoa(rand.Intn(100))

	fproject = append(fproject, project)
	json.NewEncoder(w).Encode(fproject)
	return

}

func updateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, project := range fproject {
		if project.ProjectId == params["id"] {
			fproject = append(fproject[:index], fproject[index+1:]...)

			var project FProject
			json.NewDecoder(r.Body).Decode(&project)
			project.ProjectId = params["id"]

			fproject = append(fproject, project)
			json.NewEncoder(w).Encode(project)
			return
		}
	}
}

func removeProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, project := range fproject {
		if project.ProjectId == params["id"] {
			fproject = append(fproject[:index], fproject[index+1:]...)
			json.NewEncoder(w).Encode(fproject)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	fproject = append(fproject, FProject{ProjectId: "1", Title: "AI Bot", Branch: "CS", Students: &Student{FirstName: "Pranav", LastName: "Thorve"}})

	r.HandleFunc("/", allProjects).Methods("GET")
	r.HandleFunc("/getcandidate/{id}", getSingleProject).Methods("GET")
	r.HandleFunc("/addcandidate", addProject).Methods("POST")
	r.HandleFunc("/updateInfo/{id}", updateProject).Methods("PUT")
	r.HandleFunc("/removecandidate/{id}", removeProject).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}
