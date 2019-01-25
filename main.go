package main

import (
	"fmt"
	"github.com/davswo/BlogWebServices/blog"
	"github.com/davswo/BlogWebServices/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vrischmann/envconfig"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting BlogWebServices...")

	var cfg config.Service
	if err := envconfig.Init(&cfg); err != nil {
		log.Panicf("Error loading main configuration %v\n", err.Error())
	}
	log.Print(cfg)

	if err := startService(cfg.Port); err != nil {
		log.Fatal("Unable to start server", err)
	}
}

func startService(port string) error {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blogs", blog.GetAllBlogs).
		Methods(http.MethodGet)

	router.HandleFunc("/user/login",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("You are logged in.. NOT\n"))
		}).
		Methods(http.MethodPost)

	router.HandleFunc("/user/logout",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("You are logged out.. NOT\n"))
		}).
		Methods(http.MethodPost)

	router.HandleFunc("/user/blogs",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("The Users Blogs will be returned in short form\n"))
		}).
		Methods(http.MethodGet)

	router.HandleFunc("/user/blog/{blogid}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			w.Write([]byte(fmt.Sprintf("Blog with id '%v' will be returned\n", vars["blogid"])))
		}).
		Methods(http.MethodGet)

	router.HandleFunc("/user/blog/{blogId}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Update or create a new Blog\n"))
		}).
		Methods(http.MethodPost)

	log.Printf("Starting server on port %s ", port)

	c := cors.AllowAll()
	return http.ListenAndServe(":"+port, c.Handler(router))
}
