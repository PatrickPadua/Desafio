module app

go 1.12

require (
	github.com/aws/aws-sdk-go v1.20.19
	github.com/gin-contrib/gzip v0.0.1 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/h2ik/go-sqs-poller/v2 v2.0.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/validator.v2 v2.0.0-20180514200540-135c24b11c19
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
