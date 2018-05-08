package main
import(
"encoding/json"
    "log"
"fmt"
    "net/http"
    "github.com/gorilla/mux"
)
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person
func GetPeople(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json; charset=utf-8")
w.Header().Set("Access-Control-Allow-Origin", "*")
json.NewEncoder(w).Encode(people)

}
func Handler(w  http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "connected")
}


func main(){
people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
 router := mux.NewRouter()
 router.HandleFunc("/",Handler)
    router.HandleFunc("/people", GetPeople).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}

