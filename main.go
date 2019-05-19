package main

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	serviceLambda "github.com/aws/aws-sdk-go/service/lambda"
	"log"
	"os"
)

func proxy(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {

	result, err := serviceLambda.New(getNewSession()).Invoke(createLambdaInvokeInput(getRequestBytes(request)))
	if err != nil {
		log.Printf("Cannot proxy request %+v. Error: %+v.", request, err)
		return events.APIGatewayProxyResponse{}, errors.New(err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode: int(*result.StatusCode),
		Body:       string(result.Payload),
	}, nil

}

func getRequestBytes(request events.APIGatewayProxyRequest) []byte {
	bytes, err := json.Marshal(request)
	if err != nil {
		log.Panicf("Cannot obtain bytes from request %+v", request)
	}
	return bytes
}

func getNewSession() *session.Session {
	region := os.Getenv("region")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Panicf("Cannot create new session with region %v", region)
	}
	return sess
}

func createLambdaInvokeInput(bytes []byte) *serviceLambda.InvokeInput {
	input := &serviceLambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("functionName")),
		InvocationType: aws.String(os.Getenv("invocationType")),
		LogType:        aws.String(os.Getenv("logType")),
		Payload:        bytes,
	}
	if qualifier := os.Getenv("qualifier"); qualifier != "" {
		input.Qualifier = &qualifier
	}
	if clientContext := os.Getenv("clientContext"); clientContext != "" {
		input.ClientContext = &clientContext
	}
	return input
}

func main() {
	lambda.Start(proxy)
}
