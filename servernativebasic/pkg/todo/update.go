package todo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	pb "servernativebasic/gen/protos/todopb"
	"strings"
	"time"
)

func (List TodoList) Update() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		TodoMessage := pb.TodoMessageRequest{}
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || json.Unmarshal(body, &TodoMessage) != nil || strings.TrimSpace(TodoMessage.Task.Value) == "" || strings.TrimSpace(TodoMessage.Id) == "" {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if _, exist := List[TodoMessage.Id]; !exist {
			http.Error(w, "Resource Not Found", http.StatusNotFound)
			return
		}

		List[TodoMessage.Id] = &pb.Task{Value: TodoMessage.Task.Value}

		response, err := json.Marshal(pb.TodoListResponse{List: List})

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseSize, _ := w.Write(response)

		elapsedTime := time.Since(startTime)
		log.Println("Request Time Taken:", elapsedTime, "ns Response Size: ", responseSize, "bytes")
	})
}
