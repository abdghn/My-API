/*
 * Created on 07/01/22 20.17
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package request

type CreateUpdateTask struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type CreateUpdateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
