package main

import (
	api "example/kitex_gen/api/echo"
	"log"
)

// setup new kitex server
func main() {
	svr := api.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
