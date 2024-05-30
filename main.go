package main

import (
	"context"
	"fmt"
	"net/http"
	"van_thailand_server/controller"
	"van_thailand_server/database"
	"van_thailand_server/storage"
)

func main() {
	ctx := context.Background()
	mongoDB := database.ConnectDB(ctx)
	storage.Init(ctx)

	controller.HandleRequest(ctx)
	controller.HandleAuth(ctx)

	http.ListenAndServe(":8080", nil)

	defer mongoDB.Disconnect(ctx)
	defer fmt.Println("MongoDB disconnected")
}
