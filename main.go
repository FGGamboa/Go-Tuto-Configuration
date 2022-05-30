package main

import (
	"context"
	"fmt"
	"tuto-configuration/configuration"
	"tuto-configuration/database"
)

func main() {
	var id int
	var err error
	ctx := context.Background()

	config, err := configuration.Load("config.yml")
	if err != nil {
		panic(err)
	}

	//create connection
	db := database.CreateConnection(ctx, config)
	defer db.Close()

	err = db.QueryRowContext(ctx, "SELECT id FROM books").Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	fmt.Println("All good")
}
