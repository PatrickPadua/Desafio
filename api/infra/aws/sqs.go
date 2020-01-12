package aws

import (
	"log"
	"fmt"
	lib "app/lib"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SQSSend(message, ConsulURL, ConsulEndpointURL, ConsulEndpointGROUP string){

	sess := session.New(&aws.Config{
		MaxRetries:  aws.Int(5),
	})

	svc := sqs.New(sess)

	queueURL      := lib.ConsulRequest(ConsulURL, ConsulEndpointURL)
	queueGroupID  := lib.ConsulRequest(ConsulURL, ConsulEndpointGROUP)

	sendParams := &sqs.SendMessageInput{
		MessageDeduplicationId: aws.String(lib.GetMD5Hash(message)),
		MessageGroupId: aws.String(queueGroupID),
		MessageBody:  aws.String(message),
		QueueUrl:     aws.String(queueURL),
		DelaySeconds: aws.Int64(0),
	}

	sendResp, err := svc.SendMessage(sendParams)

	if err != nil {
		fmt.Println(err)
	}
	log.Printf("[API][INFO][SEND] Message ID: %s\n", *sendResp.MessageId)
}