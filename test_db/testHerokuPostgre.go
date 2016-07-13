package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
)

func main() {
	queryPostgres := true
	


	fmt.Println("1. Open postgres db")
	db, err := sql.Open("postgres", "postgres://ukwklxchohiwrt:m_9V5QFtURhM6JKjlkBAVyNvvm@ec2-54-83-202-64.compute-1.amazonaws.com:5432/dbbpbsb92frcn6")
	checkErr(err)

	fmt.Println("2. Ping to postgres")
	err = db.Ping()
	checkErr(err)

	if (queryPostgres == true) {
		var (
			id int
			city string
		)

		fmt.Println("3. Query postgres")
		rows, err := db.Query("select id, city from weather where id = 1")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		fmt.Println("4. Scan result")
		for rows.Next() {
			err := rows.Scan(&id, &city)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, city)
			fmt.Println(id, city)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		db.Close()
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}