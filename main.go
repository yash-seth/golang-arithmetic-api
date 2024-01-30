package main

import (
	"fmt"
	"go-api/router"
)

func main() {
	fmt.Println("Hello World!")
	router.InitRouter().Run() // runs on port 8080 by default
}
