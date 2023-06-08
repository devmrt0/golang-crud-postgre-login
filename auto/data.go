package auto

import "cloudgobackend/api/models"

var users = []models.User{
	models.User{Name: "steven victor", Email: "steven@gmail.com", Password: "password"},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title",
		Content: "Hello world",
	},
}
