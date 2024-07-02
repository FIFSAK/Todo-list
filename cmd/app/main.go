package main

import (
	"Todo-list/internal/handlers"
	"Todo-list/internal/models"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	db, err := initializeDB()
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	defer db.Close()

	taskModel := models.TaskModel{DB: db}

	router := mux.NewRouter()

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/api/todo-list/tasks", handlers.CreateTaskHandler(&taskModel)).Methods(http.MethodPost)
	router.HandleFunc("/api/todo-list/tasks/{id:[0-9]+}", handlers.UpdateTaskHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/todo-list/tasks/{id:[0-9]+}", handlers.DeleteTaskHandler).Methods(http.MethodDelete)
	router.HandleFunc("/api/todo-list/tasks/{id:[0-9]+}/done", handlers.MarkTaskDone).Methods(http.MethodPut)
	router.HandleFunc("/api/todo-list/tasks", handlers.GetTasksHandler).Methods(http.MethodGet)
	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// graceful shutdown
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		<-signals

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Graceful shutdown failed: %v\n", err)
		}
	}()

	log.Printf("Server is starting on port %s\n", port)
	// start server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed: %v\n", err)
	}

	log.Println("Server gracefully stopped")

}
