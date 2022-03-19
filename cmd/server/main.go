package main

import (
	"fmt"

	"github.com/anfelo/comments-api-v2/internal/comment"
	"github.com/anfelo/comments-api-v2/internal/db"
	transportHttp "github.com/anfelo/comments-api-v2/internal/transport/http"
)

// Run - Responsible for the instantiation
// and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
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
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
