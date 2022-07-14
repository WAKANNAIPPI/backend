package main

import (
	controller "backend/Controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}
