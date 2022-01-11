/*
 * Created on 08/01/22 14.09
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package service

import (
	"database/sql"
	"fmt"
	"github.com/abdghn/My-API/internal/model"
)

type Service interface {
	GetTasks() (*[]model.Task, error)
	GetTaskById(id int) (*model.Task, error)
	SaveTask(todo *model.Task) error
	DeleteTaskByID(id int) error
	DeleteAllTasks() error
	UpdateTask(todo *model.Task) error
	CompleteTask(id int) error
}

type service struct {
	DB *sql.DB
}

func NewService(DB *sql.DB) Service {
	return &service{
		DB: DB,
	}
}

func (s *service) GetTasks() (*[]model.Task, error) {
	var todos []model.Task
	rows, err := s.DB.Query("SELECT id, name, status FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("[resource.db.persistent] query get todos error: %v", err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var todo model.Task
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Status); err != nil {
			return nil, fmt.Errorf("[resource.db.persistent] query get todos error: %v", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[resource.db.persistent] query get todos error: %v", err)
	}

	return &todos, nil
}

func (s *service) GetTaskById(id int) (*model.Task, error) {
	var todo model.Task

	row := s.DB.QueryRow("SELECT id, name, status FROM tasks WHERE id = ?", id)
	if err := row.Scan(&todo.ID, &todo.Name, &todo.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("[resource.db.persistent] query get todos error: %v", err)
		}
		return nil, fmt.Errorf("[resource.db.persistent] query get todos error: %v", err)
	}

	return &todo, nil
}

func (s *service) SaveTask(todo *model.Task) error {
	query := `
		INSERT INTO tasks (name, status) VALUES (?, ?)
		`

	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("[resource.db.saveTodo] begin tx failed: %v", err)
	}
	defer tx.Rollback() // if committed, will be no-op

	_, err = tx.Exec(query, todo.Name, todo.Status)
	if err != nil {
		return fmt.Errorf("[resource.db.saveTodo] insert data failed: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("[resource.db.saveTodo] commit tx failed: %v", err)
	}

	return nil

}

func (s *service) DeleteTaskByID(id int) error {
	query := `
		DELETE FROM tasks WHERE id = ? 
		`
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("[resource.db.deleteTodo] begin tx failed: %v", err)
	}
	defer tx.Rollback() // if committed, will be no-op

	_, err = tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("[resource.db.deleteTodo] delete data failed: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("[resource.db.deleteTodo] commit tx failed: %v", err)
	}
	return nil
}

func (s *service) DeleteAllTasks() error {
	query := `
		DELETE FROM tasks 
		`
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("[resource.db.DeleteAllTasks] begin tx failed: %v", err)
	}
	defer tx.Rollback() // if committed, will be no-op

	_, err = tx.Exec(query)
	if err != nil {
		return fmt.Errorf("[resource.db.DeleteAllTasks] delete data: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("[resource.db.DeleteAllTasks] commit tx failed: %v", err)
	}
	return nil
}

func (s *service) UpdateTask(todo *model.Task) error {
	query := `
		UPDATE tasks SET name = ?, status = ? WHERE id = ?  
		`
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("[resource.db.UpdateTask] begin tx failed: %v", err)
	}
	defer tx.Rollback() // if committed, will be no-op

	_, err = tx.Exec(query, todo.Name, todo.Status, todo.ID)
	if err != nil {
		return fmt.Errorf("[resource.db.UpdateTask] update data: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("[resource.db.UpdateTask] commit tx failed: %v", err)
	}
	return nil
}

func (s *service) CompleteTask(id int) error {
	query := `
		UPDATE tasks SET  status = true WHERE id = ?  
		`
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("[resource.db.CompleteTask] begin tx failed: %v", err)
	}
	defer tx.Rollback() // if committed, will be no-op

	_, err = tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("[resource.db.CompleteTask] complete task: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("[resource.db.CompleteTask] commit tx failed: %v", err)
	}
	return nil
}
