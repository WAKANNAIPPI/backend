package main

import controller "src/Controller"

func main() {
	r := controller.GetRouter()

	r.Run()
}
