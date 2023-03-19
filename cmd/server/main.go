package main

import (
	"context"
	"fmt"

	"github.com/TutorialEdge/go-rest-api-course/internal/db"
)

func Run() error {
	fmt.Println("Starting our application")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the db")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest Api Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}