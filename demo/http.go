package main

import (
	"log"
	"net/http"
)

func handlerTest(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)
	if r.Method == "GET" {
		raw := `{"name" : "Snooker","age":28}`
		w.Write([]byte(raw))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func main() {
	http.HandleFunc("/test", handlerTest)

	log.Println("Start server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
