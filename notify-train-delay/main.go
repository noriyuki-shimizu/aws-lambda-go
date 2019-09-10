package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Handler return (events.APIGatewayProxyResponse error)
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	line := Line{}
	err := line.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}
	eve, err := ParseRequest(line.ChannelSecret, request)
	if err != nil {
		status := 200
		if err == linebot.ErrInvalidSignature {
			status = 400
		} else {
			status = 500
		}
		return events.APIGatewayProxyResponse{StatusCode: status}, errors.New("Bat Request")
	}
	line.EventRouter(eve)
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
