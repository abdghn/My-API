/*
 * Created on 06/01/22 15.20
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(dataSourceName string) (*sql.DB, error) {
	c, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("[resource.db] open sql connection failed: %v", err)
	}

	err = c.Ping()
	if err != nil {
		return nil, fmt.Errorf("[resource.db] ping db failed: %v", err)
	}

	fmt.Println("Connected!")
	return c, nil
}
