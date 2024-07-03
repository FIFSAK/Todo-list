package handlers

import (
	"Todo-list/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

var (
	id int
	db sync.Map
)

func generateID() int {
	id++
	return id
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string `json:"title"`
		ActiveAt string `json:"activeAt"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	t := models.NewTask(0, "active", input.Title, input.ActiveAt) //return a pointer to a new task

	if err != nil {
		http.Error(w, "Error reading body: "+err.Error(), http.StatusBadRequest)
		return
	}
	//data validation
	err = dataValidation(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//duplicate check
	var duplicateFound bool
	db.Range(func(_, value interface{}) bool {
		existingTask := value.(*models.Task)
		if existingTask.Title == t.Title && existingTask.ActiveAt == t.ActiveAt {
			duplicateFound = true
			return false
		}
		return true
	})

	if duplicateFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	t.ID = generateID()
	db.Store(t.ID, t)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(t.ID)))

}
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	intItemId, _ := strconv.Atoi(itemID)

	t, ok := db.Load(intItemId)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	old_task := t.(*models.Task)
	var input struct {
		Title    string `json:"title"`
		ActiveAt string `json:"activeAt"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	new_task := models.NewTask(generateID(), "active", input.Title, input.ActiveAt)

	old_task.Title = new_task.Title
	old_task.ActiveAt = new_task.ActiveAt
	//not necessary to store the new task in the db because we work with pointer
	w.WriteHeader(http.StatusNoContent)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	intItemId, _ := strconv.Atoi(itemID)

	_, ok := db.Load(intItemId)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	db.Delete(intItemId)
}

func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	intItemId, _ := strconv.Atoi(itemID)

	t, ok := db.Load(intItemId)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	task := t.(*models.Task)
	task.Status = "done"
	w.WriteHeader(http.StatusNoContent)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		status = "active"
	}
	if status != "active" && status != "done" {
		http.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}
	data := make([]*models.Task, 0)
	if status == "done" {
		db.Range(func(_, value interface{}) bool {
			task := value.(*models.Task)
			if task.Status == "done" {
				data = append(data, task)
			}
			return true
		})
	}
	if status == "active" {
		db.Range(func(_, value interface{}) bool {
			task := value.(*models.Task)
			if task.Status == "active" {
				activeAt, err := time.Parse("2006-01-02", task.ActiveAt)
				if err != nil {
					http.Error(w, "Error parsing date: "+err.Error(), http.StatusBadRequest)
					return false
				}
				if activeAt.Before(time.Now()) || activeAt.Equal(time.Now()) {
					data = append(data, task)
				}
			}
			return true
		})
	}
	sort.Sort(models.TaskComparator(data))
	for _, task := range data {
		activeAt, err := time.Parse("2006-01-02", task.ActiveAt)
		if err != nil {
			http.Error(w, "Error parsing date: "+err.Error(), http.StatusBadRequest)
			return
		}
		if activeAt.Weekday() == time.Saturday || activeAt.Weekday() == time.Sunday {
			task.Title = "ВЫХОДНОЙ - " + task.Title
		}
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding data: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func dataValidation(t *models.Task) error {
	if t.Title == "" {
		return fmt.Errorf("Title is required")
	}
	if t.ActiveAt == "" {
		return fmt.Errorf("ActiveAt is required")
	}
	//if t.ActiveAt.IsZero() {
	//	http.Error(w, "ActiveAt is required", http.StatusBadRequest)
	//	return
	//
	//}
	if len([]rune(t.Title)) > 200 {
		return fmt.Errorf("Title is too long")
	}
	return nil
}
