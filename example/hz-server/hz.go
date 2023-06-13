package main

// imports
import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"example/kitex_gen/api"
	"example/kitex_gen/api/echo"

	"github.com/cloudwego/kitex/client"
)

// Person struct, for demo purposes. Consists of name and email address
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	// initiates new hertz server at port 8080
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	// Get request with JSON body. To test routing and hertz server. 
	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	//json format to access := `{"name": "ArRay", "email": "ArRay@example.com"}`
	//endpoint page: parses JSON body, calls Kitex server's echo function with the name received
	h.POST("/endpoint", func(c context.Context, ctx *app.RequestContext) {

		// gets request body
		requestBody, err := ctx.Body()
		if err != nil {
			// Handle error
			fmt.Println("Error")
			return
		}

		// initates new person object using reqest body 
		var person Person
		err = json.Unmarshal(requestBody, &person)
		if err != nil {
			// Handle error
			fmt.Println("Error")
		}

		// prints sender's name and email
		fmt.Println("Name:", person.Name)
		fmt.Println("Email:", person.Email)

		// initates new kitex client
		client, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
		if err != nil {
			log.Fatal(err)
		}

		// request to be sent; consists of name of the person
		req := &api.Request{Message: person.Name}

		// sends request, accepts response. Response should be in the form of "Hello, [NAME]"
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		// prints response in terminal
		log.Println(resp)

		// creates json response object encompassing the kitex response
		jsonResponse := utils.H{
			"kitexResponse": resp,
		}

		// Convert the JSON response to bytes
		responseBytes, err := json.Marshal(jsonResponse)
		if err != nil {
			// Handle error
			fmt.Println("Error marshaling JSON response:", err)
			return
		}

		// Return the JSON response to the HTTP request sender
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetBody(responseBytes)

	})

	//start the server and listen for incoming requests.
	//blocks execution of the main goroutine until the server is shut down.
	h.Spin()
}
