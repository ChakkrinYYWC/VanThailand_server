package main

import (
	"context"
	"fmt"
	"net/http"
	"van_thailand_server/controller"
	"van_thailand_server/database"
)

func main() {
	ctx := context.Background()
	client := database.ConnectDB(ctx)
	controller.HandleRequest(ctx)
	controller.HandleAuth(ctx)

	http.ListenAndServe(":8080", nil)

	defer client.Disconnect(ctx)
	defer fmt.Println("MongoDB disconnected")
}
