package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	_ "github.com/lib/pq"
)

const (
	host         = "127.0.0.1"
	port         = 5432
	user         = "postgres"
	password     = "password1"
	dbname       = "files_data"
	goroutineNum = 3
)

//var db *sql.DB

func main() {
	//PArams
	pathStr := flag.String("path", ".", "Path to the traking folder")
	param := flag.String("p", "", "D-Drop table")
	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//Connect to DB
	//var err error
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//Check the ping to DB
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//Drop table if param D
	if *param == "D" {
		_, err := db.Query("DROP TABLE files_Data;")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Table was deleted")
		}

	} else { //add files from dir if param is default
		var checkDatabase string
		db.QueryRow("SELECT to_regclass('files_Data')").Scan(&checkDatabase)
		if err != nil {
			fmt.Print(err)
		}
		//if table dose not exist then create one to use for this example
		if checkDatabase == "" {
			fmt.Println("Database Created")
			createSQL := "CREATE TABLE files_data (id SERIAL PRIMARY KEY,name VARCHAR ,fsize INT,fdate VARCHAR);"
			db.Query(createSQL)
		}

		//sql to insert employee information
		sqlStatement := "INSERT INTO files_data(name,fsize,fdate) VALUES($1,$2,$3)"
		//prepare statement for sql
		stmt, err := db.Prepare(sqlStatement)
		if err != nil {
			fmt.Print(err)
		}
		defer stmt.Close()
		files, err := ioutil.ReadDir(*pathStr)
		if err != nil {
			log.Fatal(err)
		}
		inputFile := make(chan os.FileInfo, 1)
		for i := 0; i < goroutineNum; i++ {
			go SaveToDb(db, sqlStatement, inputFile)
		}
		for _, file := range files {
			if !file.IsDir() {
				inputFile <- file
			} else if err != nil {
				fmt.Print(err)
			}

		}
		close(inputFile)
		//time.Sleep(time.Milliseconds)
		rows, err := db.Query("SELECT * FROM files_data")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		fmt.Println("---------------------------------------------------------------------")
		for rows.Next() {
			var id int
			var name string
			var fsize int64
			var fdata string
			err := rows.Scan(&id, &name, &fsize, &fdata)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Printf("%s \n", name)
		} //end of for loop
	}
}

func SaveToDb(db *sql.DB, sqlStatement string, in <-chan os.FileInfo) {
	for file := range in {

		row := db.QueryRow("SELECT name FROM files_data WHERE name=$1", file.Name())
		//fmt.Println(file.Name())
		id := 0
		err := row.Scan(&id)
		if err == sql.ErrNoRows {
			_, err = db.Exec(sqlStatement, file.Name(), file.Size(), file.ModTime())
			if err != nil {
				panic(err)
			}
		}
		runtime.Gosched()
	}

}
