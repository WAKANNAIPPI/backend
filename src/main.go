package main

import controller "backend/Controller"

func main() {
	r := controller.GetRouter()

	r.Run()
}
