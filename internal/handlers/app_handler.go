package handlers

import (
	"Todo-list/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func CreateTaskHandler(taskModel *models.TaskModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title    string `json:"title"`
			ActiveAt string `json:"activeAt"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		t := models.NewTask(0, "active", input.Title, input.ActiveAt)

		if err != nil {
			http.Error(w, "Error reading body: "+err.Error(), http.StatusBadRequest)
			return
		}
		if t.Title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}
		if t.ActiveAt.IsZero() {
			http.Error(w, "ActiveAt is required", http.StatusBadRequest)
			return

		}
		if len([]rune(t.Title)) > 200 {
			http.Error(w, "Title is too long", http.StatusBadRequest)
			return
		}
		//TODO:
		//проверка уникальность таска по названию и дате
		//Если запись существует, возвращается код 404.
		fmt.Println(t)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(strconv.Itoa(t.ID)))

	}
}
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func MarkTaskDone(w http.ResponseWriter, r *http.Request) {

}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {

}
