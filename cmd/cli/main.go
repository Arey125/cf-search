package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	database "cf-search/internal/db"
	"cf-search/internal/problems"
)

func main() {
    fmt.Println("fetching problems")
	pr, err := fetchProblems()
	if err != nil {
		panic(err)
	}
    fmt.Printf("%d problems fetched\n", len(pr))

    dsn := os.Getenv("DB")
    db := database.Connect(dsn)

    problemModel := problems.NewModel(db)

    err = problemModel.DeleteAll()
    if err != nil {
        panic(err)
    }
    err = problemModel.AddMany(pr)
    if err != nil {
        panic(err)
    }
}
