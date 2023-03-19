package main

import (
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	}

	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }

	fmt.Println("succesfully connected and pinged db")

	return nil
}

func main() {
	fmt.Println("Go Rest Api Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
