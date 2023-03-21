package main

import (
	"context"
	"fmt"

	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
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

	cmtService := comment.NewService(db)
	cmtService.PostComment(context.Background(), comment.Comment{
		ID:     "",
		Slug:   "manual-test",
		Body:   "Hello from comment",
		Author: "salenjhons",
	})

	//fmt.Println(cmtService.GetComment(context.Background(), "9a31bf83-28dc-4d8d-bf70-7d347a24"))

	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }

	//fmt.Println("succesfully connected and pinged db")
	return nil
}

func main() {
	fmt.Println("Go Rest Api Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
