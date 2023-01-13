// main.go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

/* func hello() (string, error) {

	return "Hello λ!", nil
} */

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func hello() (Response, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return Response{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}

/* func hello() (map[string]interface{}, error) {
	return map[string]interface{}{
			"statusCode": 200,
			"headers":    map[string]string{"Content-Type": "application/json"},
			"body":       "Hello World",
		},
		nil
} */

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
