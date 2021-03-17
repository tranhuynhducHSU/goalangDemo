package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

type Student struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	IsMale bool   `json:"IsMale"`
}

func decodeJson(req *http.Request, newStudent interface{}) interface{} {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newStudent)
	if err != nil {
		panic(err)
	}
	return newStudent
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is listening.....")
	fmt.Println("Endpoint Hit: homePage")
}
func returnAllStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nEndpoint Hit: returnAllStudent")
	json.NewEncoder(w).Encode(studens)
}
func createNewStudent(res http.ResponseWriter, req *http.Request) {
	var jBody interface{}
	jBody = decodeJson(req, jBody)
	newStudent := Student{}
	mapstructure.Decode(jBody, &newStudent)

	log.Printf("\nCreate user [%v] success.", newStudent.Name)
	jNewStudent, _ := json.Marshal(newStudent)
	fmt.Fprint(res, string(jNewStudent))
}
func updateStudent(res http.ResponseWriter, req *http.Request) {

}
func delateStudent(res http.ResponseWriter, req *http.Request) {

}
func findStudent(res http.ResponseWriter, req *http.Request) {

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/student", createNewStudent).Methods("POST")
	myRouter.HandleFunc("/student", returnAllStudent)

	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

var studens []Student

func main() {
	studens = append(studens, Student{
		ID:     "2162398",
		Name:   "Trần Huỳnh Đức",
		IsMale: true,
	})
	studens = append(studens, Student{
		ID:     "2171289",
		Name:   "Nguyễn Hoàng Phong Phú",
		IsMale: true,
	})
	fmt.Print("Server is listening on: 3000")
	handleRequests()
}
