package main

import "test-code-mst/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
