package main

import "echo-rest/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
