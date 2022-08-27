package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Message struct {
	EndPoint string   `json:"EndPoint"`
	Msg      []string `json:"Message"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if !check(auth) {
			w.WriteHeader(401)
			return
		}

		log.Println(r.Header.Get("user-agent"))
		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	router.HandleFunc("/tagsCount", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")

		response := Message{EndPoint: "tagsCount", Msg: []string{query, "go home"}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	router.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))

		response := Message{EndPoint: "tags", Msg: []string{query, strconv.Itoa(limit), strconv.Itoa(skip), "go home"}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	router.HandleFunc("/gifs", func(w http.ResponseWriter, r *http.Request) {
		tag := r.URL.Query().Get("tag")
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))

		response := Message{EndPoint: "gifs", Msg: []string{tag, strconv.Itoa(limit), strconv.Itoa(skip), "go home"}}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}
