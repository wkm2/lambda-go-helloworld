// main.go
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
        "time"
)

func hello() (string, error) {
	return "Hello ƛ!", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
        time.Sleep(3 * time.Second)
	lambda.Start(hello)
}
