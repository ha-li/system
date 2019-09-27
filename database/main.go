package main

import (
	"database/sql"
	"fmt"
	//_ "gopkg.in/goracle.v2"
	_ "github.com/mattn/go-oci8"
)

func main () {
	fmt.Println("Hello Database")
	//connect := "IHUB_OWNER/IHUB_OWNER@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=localhost)(PORT=49161)))(CONNECT_DATA=(SERVICE_NAME=xe)))"
	connect := "IHUB_OWNER/IHUB_OWNER@localhost:49161/xe"
	db, err := sql.Open("goracle", connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows,err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error during query")
		fmt.Println(err)
		return
	}

	defer rows.Close()

	var thedate string
	for rows.Next() {
		rows.Scan(&thedate)
	}
	fmt.Printf("The date is %s\n", thedate)
}
