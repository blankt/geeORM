package main

import "fmt"

func main() {
	//db, _ := sql.Open("sqlite", "gee.db")
	//defer func() { _ = db.Close() }()
	//_, _ = db.Exec("drop table  if exists User;")
	//_, _ = db.Exec("create  table  User(Name text);")
	//result, err := db.Exec("insert  into User(Name) values (?),(?)", "tlf", "pxy")
	//if err != nil {
	//	affects, _ := result.RowsAffected()
	//	fmt.Println(affects)
	//}
	//var name string
	//row := db.QueryRow("select name from user limit 1")
	//if err = row.Scan(&name); err != nil {
	//	fmt.Println(name)
	//}
	var s string
	for i := 0; i < 10; i++ {
		if s == "" {
			s = "初始化"
			fmt.Println(s)
		}
	}
}
