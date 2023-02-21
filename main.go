package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Start server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(map[string]string{
			"time": time.Now().String(),
		})
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
