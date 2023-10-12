 package handler

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Post struct{
  DB *pgxpool.Pool
}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
  

}

func (p *Post) List(w http.ResponseWriter, r *http.Request) {
  fmt.Println("List all posts")
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request){
  fmt.Println("Get post by ID")
}

func (p *Post) UpdateByID(w http.ResponseWriter, r *http.Request){
  fmt.Println("Update a post by ID")
}

func (p *Post) DeleteByID(w http.ResponseWriter, r *http.Request){
  fmt.Println("Delete an order by ID")
}
