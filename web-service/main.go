package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/ravjotsingh9/discussionForum-Web-Service/db"
)

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", getWelcomeHandler).
		Methods("GET")
	router.HandleFunc("/welcome/{name}", getWelcomeHandler).
		Methods("GET")
	router.HandleFunc("/getTopic/{id}", getTopicHandler).
		Methods("GET")
	router.HandleFunc("/topic/", createCommentHandler).
		Methods("POST")
	return
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	/*
		//connect to db
		addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
		db, err := sql.Open("postgres", addr)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer db.Close()
	*/

	addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	repo, err := db.NewPostgres(addr)
	if err != nil {
		log.Println(err)
		return
	}
	db.SetRepository(repo)

	defer db.Close()

	// Run HTTP server

	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

}
