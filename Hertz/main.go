// Code generated by hertz generator.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type RequestBody struct {
	Name string `json:"name"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func main() {
	//initialized with the default server
	h := server.Default()

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		//send a JSON response with an HTTP status code of 200 (OK)
		//response body is a JSON object with a single key-value pair,
		//where the key is "message" and the value is "pong".
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.POST("/greet", func(c context.Context, ctx *app.RequestContext) {
		body, err := ioutil.ReadAll(ctx.Request().Body())
		if err != nil {
			ctx.JSON(consts.StatusBadRequest, ResponseBody{Message: "Invalid request body"})
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			ctx.JSON(consts.StatusBadRequest, ResponseBody{Message: "Failed to decode request body"})
			return
		}

		responseBody := ResponseBody{
			Message: fmt.Sprintf("Hello, %s!", requestBody.Name),
		}

		ctx.JSON(consts.StatusOK, responseBody)
	})

	//start the server and listen for incoming requests.
	//blocks execution of the main goroutine until the server is shut down.
	h.Spin()
}
