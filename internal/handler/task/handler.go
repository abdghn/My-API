/*
 * Created on 06/01/22 15.33
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package handler

import (
	json "encoding/json"
	"github.com/abdghn/My-API/internal/request"
	usecase "github.com/abdghn/My-API/internal/usecase/task"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Handler interface {
	HandleGetTasks(w http.ResponseWriter, r *http.Request)
	HandleGetTaskByID(w http.ResponseWriter, r *http.Request)
	HandleHealthCheck(w http.ResponseWriter, r *http.Request)
	HandleSaveTask(w http.ResponseWriter, r *http.Request)
	HandleDeleteTaskByID(w http.ResponseWriter, r *http.Request)
	HandleDeleteAllTask(w http.ResponseWriter, r *http.Request)
	HandleUpdateTask(w http.ResponseWriter, r *http.Request)
	HandleCompleteTask(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	uc usecase.Usecase
}

// New creates a new handler
func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

func (h *handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h *handler) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	todos, err := h.uc.GetTasks()
	if err != nil {
		log.Printf("[handler.HandleGetTasks] unable to find todos: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{"data": todos})
	if err != nil {
		log.Printf("[handler.HandleGetTasks] unable to find todos: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}
}

func (h *handler) HandleGetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("[handler.HandleGetTaskByID] convert string to int failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	todo, err := h.uc.GetTaskByID(id)
	if err != nil {
		log.Printf("[handler.HandleGetTasks] unable to find todos: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{"data": &todo})
	if err != nil {
		log.Printf("[handler.HandleGetTasks] unable to find todos: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}
}

func (h *handler) HandleSaveTask(w http.ResponseWriter, r *http.Request) {
	req := request.CreateUpdateTask{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[handler.HandleGetTasks] error read request data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("[handler.HandleGetTasks] error read request data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	err = h.uc.SaveTodo(&req)
	if err != nil {
		log.Printf("[HandleCreateUser] error create user: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte("error create user"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "ok"}`))
}

func (h *handler) HandleDeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("[handler.HandleDeleteTaskByID] convert string to int failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	err = h.uc.DeleteTask(id)
	if err != nil {
		log.Printf("[handler.HandleDeleteTaskByID] unable to delete task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "ok"}`))

}

func (h *handler) HandleDeleteAllTask(w http.ResponseWriter, r *http.Request) {
	err := h.uc.DeleteAllTask()
	if err != nil {
		log.Printf("[handler.HandleDeleteAllTask] unable to delete all task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "ok"}`))
}

func (h *handler) HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	req := request.CreateUpdateTask{}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("[handler.HandleUpdateTask] convert string to int failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[handler.HandleUpdateTask] error read request data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("[handler.HandleUpdateTask] error read request data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	err = h.uc.UpdateTask(id, &req)
	if err != nil {
		log.Printf("[HandleCreateUser] error update task: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte("error create user"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "ok"}`))

}

func (h *handler) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("[handler.HandleCompleteTask] convert string to int failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	err = h.uc.CompleteTask(id)
	if err != nil {
		log.Printf("[handler.HandleCompleteTask] unable to complete task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "ok"}`))

}
