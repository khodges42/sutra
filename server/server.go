package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sutra/models"
	"sutra/userstore"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/message/{messageHash}", GetMessage).Methods("GET")
	router.HandleFunc("/messages/{userHash}", GetMessages).Methods("GET")
	router.HandleFunc("/user/{userHash}", GetUser).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")

	router.HandleFunc("/message", SendMessage).Methods("POST")
	router.HandleFunc("/user", CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":6969", router))
}

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	users := userstore.NewUserStore()
	json.NewEncoder(w).Encode(&users)
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	// read the user input
	// slap user input into db and write to file.

	var jsonBlob = []byte(`
        {
	  "MessageHash":"foo",
      "rootHash":"foo",
      "toHash":"bar",
      "fromHash":"baz",
      "sentTime":12345,
      "content":"fizz"
   }`)
	message := models.Message{}
	err := json.Unmarshal(jsonBlob, &message)
	if err != nil {
		log.Fatal("there was some fuckshit")
	}

	rankingsJson, _ := json.Marshal(message)
	err = ioutil.WriteFile("F:\\trash\\messages.json", rankingsJson, 0644)

	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	dat, _ := ioutil.ReadFile("F:\\trash\\messages.json")

	json.NewEncoder(w).Encode(string(dat))
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	dat, _ := ioutil.ReadFile("F:\\trash\\messages.json")

	json.NewEncoder(w).Encode(string(dat))
}



func GetUser(w http.ResponseWriter, r *http.Request) {
	users := userstore.NewUserStore()
	json.NewEncoder(w).Encode(&users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	users := userstore.NewUserStore()
	json.NewEncoder(w).Encode(&users)
}


// endpoint to post data
//     POST http://server/sutra/api/message
//			{
//   			"message":{
//     		 		"rootHash":"foo",
//      	 		"toHash":"bar",
//      	 		"fromHash":"baz",
//      	 		"sentTime":12345,
//      	 		"content":"fizz"
//   			}
//			}

// endpoint to query messages
//     GET http://server/sutra/api/messages/9a3raaaaaaaaaa

// figure out how to delete messages

// our main function