.PHONY: build clean deploy

build:
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/line/line-bot-sdk-go/linebot
	go get github.com/PuerkitoBio/goquery
	env GOOS=linux go build -ldflags="-s -w" -o bin/notify-train-delay notify-train-delay/*.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
