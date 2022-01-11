/*
 * Created on 06/01/22 15.34
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package usecase

import (
	"github.com/abdghn/My-API/internal/model"
	"github.com/abdghn/My-API/internal/request"
	service "github.com/abdghn/My-API/internal/service/task"
)

type Usecase interface {
	GetTasks() (*[]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	SaveTodo(req *request.CreateUpdateTask) error
	DeleteTask(id int) error
	DeleteAllTask() error
	UpdateTask(id int, req *request.CreateUpdateTask) error
	CompleteTask(id int) error
}

type usecase struct {
	service service.Service
}

func New(service service.Service) Usecase {
	return &usecase{service: service}
}

func (u *usecase) GetTasks() (*[]model.Task, error) {
	return u.service.GetTasks()
}
func (u *usecase) GetTaskByID(id int) (*model.Task, error) {
	return u.service.GetTaskById(id)
}

func (u *usecase) SaveTodo(req *request.CreateUpdateTask) error {
	return u.service.SaveTask(&model.Task{
		Name:   req.Name,
		Status: req.Status,
	})
}

func (u *usecase) DeleteTask(id int) error {
	return u.service.DeleteTaskByID(id)
}

func (u *usecase) DeleteAllTask() error {
	return u.service.DeleteAllTasks()
}

func (u *usecase) UpdateTask(id int, req *request.CreateUpdateTask) error {
	task, err := u.service.GetTaskById(id)
	if err != nil {
		return err
	}
	task.Name = req.Name
	task.Status = req.Status

	return u.service.UpdateTask(task)
}

func (u *usecase) CompleteTask(id int) error {
	return u.service.CompleteTask(id)
}
