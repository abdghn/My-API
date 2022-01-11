/*
 * Created on 08/01/22 02.36
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package model

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
