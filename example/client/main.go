package main

import (
	"context"
	"log"

	"example/kitex_gen/api"
	"example/kitex_gen/api/echo"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	req := &api.Request{Message: "nabei"}
	resp, err := client.Echo(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
