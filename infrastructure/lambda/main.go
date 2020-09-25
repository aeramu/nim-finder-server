package main

import (
	"context"
	"encoding/json"

	"github.com/aeramu/nim-finder-server/implementation/handler/graphql"
	"github.com/aeramu/nim-finder-server/implementation/repository/mongodb"
	"github.com/aeramu/nim-finder-server/usecase/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//convert request body to json
	var parameter struct {
		Query         string
		OperationName string
		Variables     map[string]interface{}
	}
	json.Unmarshal([]byte(request.Body), &parameter)

	repo := mongodb.New()
	interactor := service.New(repo)
	handler := graphql.New(interactor)

	response := handler.Response(ctx, parameter.Query, parameter.OperationName, parameter.Variables)
	responseJSON, _ := json.Marshal(response)

	//response
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type",
		},
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil
}
