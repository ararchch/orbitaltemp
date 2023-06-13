// Code generated by hertz generator.

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

//for testing 
type Person struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}


func main() {
	//http://127.0.0.1:8888/

	h := server.Default()

	//Get request with Json body
	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pongmain1"})
	})

	h.POST("/post", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "post"})
	})

	h.POST("/endpoint", func(c context.Context, ctx *app.RequestContext) {
		
		requestBody, err := ctx.Body()
		if err != nil {
			// Handle error
			fmt.Println("Error")
			return
		}
	
		// Process the requestBody as needed
	
		// Send the received data as a JSON response
		ctx.JSON(consts.StatusOK, requestBody)

		//jsonStr := `{"name": "ArRay", "email": "ArRay@example.com"}`

		var person Person
		err = json.Unmarshal(requestBody, &person)
		if err != nil {
			// Handle error
			fmt.Println("Error")
		}
	
		fmt.Println("Name:", person.Name)
		fmt.Println("Email:", person.Email)
	})

	//start the server and listen for incoming requests.
	//blocks execution of the main goroutine until the server is shut down.
	h.Spin()
}
