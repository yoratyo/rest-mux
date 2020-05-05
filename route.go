package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{
		{Id: 1, Title: "Post 1", Content: "Content of Post 1"},
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Post Data")

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error wrong format id post data"}`))
		return
	}

	if len(posts) < id {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error post data not found"}`))
		return
	}

	post := posts[id-1]
	result, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling post data"}`))
		return
	}

	w.Write(result)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Posts Data")

	w.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling posts data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Post Data")

	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling post data"}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	w.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling post data"}`))
		return
	}

	w.Write(result)
}
