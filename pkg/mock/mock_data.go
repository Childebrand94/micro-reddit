package mock

import (
	m "github.com/Childebrand94/micro-reddit/pkg/models"
)

var User1 = m.User{
	ID : 1,
	First_name: "Chris",
	Last_name: "Hildebrand",
	Email: "hildebrandc94@gmail.com",
}

var Post1 = m.Post{
	Author_ID: 1,
	URL: "www.amazing_example.com",
	Title: "My first amazing post",

}