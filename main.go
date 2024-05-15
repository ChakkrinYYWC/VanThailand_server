package main

import (
	"context"
	"fmt"
	"van_thailand_server/controller"
	"van_thailand_server/database"
)

func main() {
	ctx := context.Background()
	client := database.ConnectDB(ctx)
	controller.HandleRequest(ctx)

	defer client.Disconnect(ctx)
	defer fmt.Println("MongoDB disconnected")
}
