package main

import (
	"fmt"
	"test-code-mst/routers"
)

func main() {
	router, err := routers.Init()
	if err != nil {
		fmt.Printf("Error initializing application: %s \n", err)
		return
	}

	router.Run(":8080")
}
