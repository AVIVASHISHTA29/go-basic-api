package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

type Comment struct {
    PostID int    `json:"postId"`
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Body   string `json:"body"`
}

var posts = []Post{
    {UserID: 1, ID: 1, Title: "Post 1", Body: "This is the first post."},
    {UserID: 1, ID: 2, Title: "Post 2", Body: "This is the second post."},
    {UserID: 2, ID: 3, Title: "Post 3", Body: "This is the third post."},
}

var users = []User{
    {ID: 1, Name: "John Doe", Username: "johndoe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Smith", Username: "janesmith", Email: "jane@example.com"},
}

var comments = []Comment{
    {PostID: 1, ID: 1, Name: "Commenter 1", Email: "commenter1@example.com", Body: "This is the first comment."},
    {PostID: 1, ID: 2, Name: "Commenter 2", Email: "commenter2@example.com", Body: "This is the second comment."},
    {PostID: 2, ID: 3, Name: "Commenter 3", Email: "commenter3@example.com", Body: "This is the third comment."},
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
 	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Welcome to the Go API Server")
}


func postsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func commentsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comments)
}

func main() {
	http.HandleFunc("/", helloWorld)
    http.HandleFunc("/posts", postsHandler)
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/comments", commentsHandler)

    fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
