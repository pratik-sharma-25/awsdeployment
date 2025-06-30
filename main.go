package main

import "net/http"

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", greetUser)

	http.ListenAndServe(":8060", router)
}

func greetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
