/*
 * Created on 06/01/22 15.14
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package main

import (
	handler "github.com/abdghn/My-API/internal/handler/task"
	"github.com/abdghn/My-API/internal/resource/db"
	service "github.com/abdghn/My-API/internal/service/task"
	usecase "github.com/abdghn/My-API/internal/usecase/task"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow!\n"))
}

func main() {

	DB, err := db.DbConnect("root:belajar@tcp(127.0.0.1:3306)/belajar?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	srv := service.NewService(DB)
	uc := usecase.New(srv)
	h := handler.New(uc)

	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/health", h.HandleHealthCheck).Methods(http.MethodGet, http.MethodHead)

	s := r.PathPrefix("/task").Subrouter()
	s.HandleFunc("/", h.HandleGetTasks).Methods(http.MethodGet)
	s.HandleFunc("/", h.HandleSaveTask).Methods(http.MethodPost)
	s.HandleFunc("/", h.HandleDeleteTaskByID).Methods(http.MethodDelete)
	s.HandleFunc("/", h.HandleUpdateTask).Methods(http.MethodPut)
	s.HandleFunc("/complete", h.HandleCompleteTask).Methods(http.MethodPut)
	s.HandleFunc("/all", h.HandleDeleteAllTask).Methods(http.MethodDelete)
	s.HandleFunc("/detail", h.HandleGetTaskByID).Methods(http.MethodGet)

	p := r.PathPrefix("/product").Subrouter()
	p.HandleFunc("/", h.HandleHealthCheck).Methods(http.MethodGet)
	log.Printf("server is running at port: 8000")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
