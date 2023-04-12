/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"github.com/YoniArkin/GolangAPITemplate/internal/db/mongo"
	"github.com/YoniArkin/GolangAPITemplate/internal/models"
)

func main() {
	//cmd.Execute()
	dao := mongo.GreetingDAOImpl{}
	greet1 := models.Greet{
		Message: "Hello World",
	}
	err := dao.AddGreet(context.Background(), &greet1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Greet1 added with id: %s", greet1.ID)
	fmt.Println()

	greet, err := dao.GetGreetsByPage(context.Background(), 1, 10)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, greet := range greet {
		fmt.Printf("Greet: %s. ID = %s", greet.Message, greet.ID)
		fmt.Println()

	}

	count, _ := dao.GetGreetsCount(context.Background())
	fmt.Printf("Count: %d", count)
}
