package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = 8000

func main() {
	log.Printf("Start server on port %d", port)

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

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
