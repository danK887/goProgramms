package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	SiteRate uint8  `json:"siterate"`
	Content  string `json:"content"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/feetback")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connected to the database!")

	// Create a new user data
	// insert, err := db.Query("INSERT INTO `usersdata` (name, email, siteRate, content) VALUES (?,?,?,?)", "Admin Dan2", "admin_dan2@example.com", 3, "")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	//выборка данных
	res, err := db.Query("SELECT `name`,`email`,`siteRate`,`content` FROM `usersdata`")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Email, &user.SiteRate, &user.Content)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Name: %s Email: %s feetback: %d Content: %s\n", user.Name, user.Email, user.SiteRate, user.Content)
	}

}
