package aws

import (
	"log"
	"time"
	lib "app/lib"
	"encoding/json"
	s "app/domain/suggestion"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SQSReceiveAndDrop(svc *sqs.SQS, ConsulURL, ConsulEndpoint, COLLECTION string){

	queueURL  := lib.ConsulRequest(ConsulURL, ConsulEndpoint)
	// Receive message
	receiveParams := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(1), 
	}
	receiveResp, err := svc.ReceiveMessage(receiveParams)

	if err != nil {
		log.Println(err)
	}
	
	if len(receiveResp.Messages) == 0{
		return
	}

	JobProcess, err := processMessage(receiveResp, COLLECTION)
	if err != nil {
		log.Printf("[WORKER][ERROR][PROCESS] Error: %s, Jobs count: %b \n", err, JobProcess)
		return
	}
	log.Printf("[WORKER][INFO][PROCESS][JOB] Processed: %b\n", JobProcess)


	JobDelete, err := deleteMessage(svc, queueURL, receiveResp)
	if err != nil {
		log.Printf("[WORKER][ERROR][DELETE] Error: %s, Jobs count: %b \n", err, JobDelete)
		return
	}
	log.Printf("[WORKER][INFO][DELETE][JOB] Processed: %b\n", JobDelete)

	return
}

func processMessage(m *sqs.ReceiveMessageOutput, COLLECTION string) (int64, error){
	var count int64 = 0
	for _, message := range m.Messages {
		var s s.Suggestion
		// Realiza JSON decode do campo message.Body
		decodeErr := json.Unmarshal([]byte(*message.Body), &s)
		if decodeErr != nil {
			log.Println(decodeErr)
		}

		currentTime := time.Now()
		s.Date = string(currentTime.Format("2006-01-02 15:04:05"))
		createErr := s.Create(COLLECTION)
		if createErr != nil {
			log.Println(createErr)
		}
		
		log.Printf("[WORKER][INFO][PROCESS] Message ID: %s\n", *message.MessageId)
		count++
	}

	return count, nil
}

func deleteMessage(svc *sqs.SQS, queueURL string, m *sqs.ReceiveMessageOutput) (int64, error){
	var count int64 = 0
	for _, message := range m.Messages {
		deleteParams := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),  // Required
			ReceiptHandle: message.ReceiptHandle, // Required
		}
		_, err := svc.DeleteMessage(deleteParams) // No response returned when successed.

		if err != nil {
			return count, err
		}
		count++
		log.Printf("[WORKER][INFO][DELETE] Message ID: %s\n", *message.MessageId)
	}

	return count, nil
}