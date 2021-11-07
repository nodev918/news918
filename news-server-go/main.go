package main

import (
 	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"fmt"
//	"database/sql"

	_ "github.com/lib/pq"
)
const (
	host	= "localhost"
	port	= 5432
	user	= "user"
	password= "mysecretpassword"
	dbname	= "test"
)

func main() {
/*
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)

	db, err := sql.Open("postgres",psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("Successfully connected!")
*/



	//sqlStatement := `INSERT INTO news (title,author,content) values ($1,$2,$3) returning id`
	//sqlStatement1 := `insert into "news" (title,author,content) values ("2","3","4");`
	//id := 0
	//s2 := `select * from news`

	//fmt.Println(db.Query(s2))
/*
	err = db.QueryRow(sqlStatement1).Scan(&id)
	if err != nil {
	panic(err)
	}
	fmt.Println("New record ID is:", id)
*/
	/*	
	id := 0
	err = db.QueryRow(sqlStatement, "first title","yale","this is an article content").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	*/











	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./dist",true)))
	/*
	r.GET("/web",func(c *gin.Context){
		fmt.Printf("in /web918")
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"Main website",
		})
	})
*/
	r.GET("/ping", func(c *gin.Context) {
		fmt.Printf("http connect")
		c.JSON(200, gin.H{
		

			"message": "pong",
		})
	})
	r.Run(":5000")
}
