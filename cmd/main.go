package main

import (
	"fmt"
	"log"

	db "github.com/Childebrand94/micro-reddit/pkg/database"
	m "github.com/Childebrand94/micro-reddit/pkg/mock"
	"github.com/Childebrand94/micro-reddit/pkg/models"
)

func main () {
  // Connect and build Database
	pool := db.BuildDb()
	defer pool.Close()
	
	err := db.AddUser(pool, m.User1.First_name, m.User1.Last_name, m.User1.Email)
	if err != nil{
		log.Fatalf("Failed to execute AddUser: %v", err)
	}

	err = db.AddPostByUser(pool, m.User1.ID, m.Post1.URL, m.Post1.Title) 
	if err != nil{
		log.Fatalf("Failed to execute AddPostByUser: %v", err)
	}
	
	var post *models.Post

	post ,err = db.GetPostById(pool, m.Post1.ID)
	if err != nil{
		log.Fatalf("Failed to execute GetPost: %v", err)
	}

    fmt.Println("Author ID:", post.Author_ID)
    fmt.Println("Title:", post.Title)
    fmt.Println("URL:", post.URL)

}
