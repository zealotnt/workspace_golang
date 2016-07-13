package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    "log"
)

func main() {
    db, err := sql.Open("mysql", "golang:1234@tcp(127.0.0.1:3306)/pqDatTest")
    checkErr(err)

    //ping to db
    err = db.Ping()
    checkErr(err)

    var (
        id int
        name string
    )
    rows, err := db.Query("select id, name from users where id = ?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        log.Println(id, name)
        fmt.Println(id, name)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }


    // // insert
    // stmt, err := db.Prepare("INSERT userinfo SET username=golang")
    // checkErr(err)

    // res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
    // checkErr(err)

    // id, err := res.LastInsertId()
    // checkErr(err)

    // fmt.Println(id)
    // // update
    // stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    // checkErr(err)

    // res, err = stmt.Exec("astaxieupdate", id)
    // checkErr(err)

    // affect, err := res.RowsAffected()
    // checkErr(err)

    // fmt.Println(affect)

    // // query
    // rows, err := db.Query("SELECT * FROM userinfo")
    // checkErr(err)

    // for rows.Next() {
    //     var uid int
    //     var username string
    //     var department string
    //     var created string
    //     err = rows.Scan(&uid, &username, &department, &created)
    //     checkErr(err)
    //     fmt.Println(uid)
    //     fmt.Println(username)
    //     fmt.Println(department)
    //     fmt.Println(created)
    // }

    // // delete
    // stmt, err = db.Prepare("delete from userinfo where uid=?")
    // checkErr(err)

    // res, err = stmt.Exec(id)
    // checkErr(err)

    // affect, err = res.RowsAffected()
    // checkErr(err)

    // fmt.Println(affect)

    db.Close()

}

func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}