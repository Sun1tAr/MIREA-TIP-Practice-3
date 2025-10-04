package main

import (
	"log"
	"net/http"

	"example.com/pz3-http/internal/api"
	"example.com/pz3-http/internal/env"
	"example.com/pz3-http/internal/storage"
)

func main() {
	repository := storage.NewMemoryStore()
	handlers := api.NewHandlers(repository)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Коллекция
	mux.HandleFunc("GET /tasks", handlers.ListTasks)
	mux.HandleFunc("POST /tasks", handlers.CreateTask)
	// Элемент
	mux.HandleFunc("GET /tasks/", handlers.GetTask)

	// Обновление статуса
	mux.HandleFunc("PATCH /tasks/", handlers.PatchTask)

	// Удаление задачи
	mux.HandleFunc("DELETE /tasks/", handlers.DeleteTask)

	// Подключаем логирование
	handler := api.Logging(mux)

	addr := ":" + env.GetPort()
	log.Println("listening on", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
