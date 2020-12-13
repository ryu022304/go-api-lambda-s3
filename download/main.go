package main

import (
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Response ...
type Response events.APIGatewayProxyResponse

// Request ...
type Request events.APIGatewayProxyRequest

// Handler ...
func Handler(request Request) (Response, error) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	))
	svc := s3.New(sess)

	fileName, _ := url.PathUnescape(request.QueryStringParameters["fileName"])

	resp, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(fileName),
	})

	url, _ := resp.Presign(5 * time.Minute)

	res := Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/html; charset=utf-8",
		},
		Body: url,
	}

	return res, nil
}

func main() {
	lambda.Start(Handler)
}
