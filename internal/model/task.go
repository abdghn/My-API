/*
 * Created on 06/01/22 15.36
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package model

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
