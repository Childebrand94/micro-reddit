package main

import (
	"context"
	"fmt"
	"log"
	"os"

	db "github.com/Childebrand94/micro-reddit/pkg/database"
	m "github.com/Childebrand94/micro-reddit/pkg/mock"
	"github.com/Childebrand94/micro-reddit/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// type App struct{
//   Pool *pgxpool.Pool
// }
//


func main () {
  // Connect and build Database
// Load in .env file 
  err := godotenv.Load(".env")
  if err != nil {
      log.Fatal("Error loading .env file")
  } 

//Set up database connection
  	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
 
	err = db.AddUser(pool, m.User1.First_name, m.User1.Last_name, m.User1.Email)
	if err != nil {
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

	// Starting Server 
	// fmt.Println("Starting Server....")
	// r := chi.NewRouter()
	// r.Get("/", getAllPosts)
	// r.Post("/user", createNewUser)

	
	//  log.Fatal(http.ListenAndServe(":8010", r))
}

// // helper function form article 
// func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
//
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
//
// Post method to create a new user
// func createNewUser(w http.ResponseWriter, r *http.Request, pool pgxpool.Pool){
//   var user models.User
//   json.NewDecoder(r.Body).Decode(&user)
//
//   err:= db.AddUser(pool, user.First_name, user.Last_name, user.Email)
//   if err != nil {
//     log.Fatalf("Failed to execute AddUser: %v", err)
// }
//
//   respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully created"})
// }
//

