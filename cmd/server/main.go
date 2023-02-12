package main

import (
	"fmt"
	"github.com/sajjadanwar0/go-rest-api/internal/comment"
	db2 "github.com/sajjadanwar0/go-rest-api/internal/db"
	transportHttp "github.com/sajjadanwar0/go-rest-api/internal/transport/http"
)

// Run -- Going to be responsible for  the instantiation and startup of our Go application
func Run() error {
	fmt.Println("Starting up our application")
	db, err := db2.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Migrate(); err != nil {
		fmt.Println("failed to migrate the database")
		return err
	}

	cmtService := comment.NewService(db)
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
