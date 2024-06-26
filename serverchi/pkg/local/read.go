package local

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (List TodoList) Read() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		response, err := json.Marshal(List)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseSize, _ := w.Write(response)

		elapsedTime := time.Since(startTime)
		log.Printf("Request Time Taken: %s, Response Size: %d bytes", elapsedTime, responseSize)
	})
}
