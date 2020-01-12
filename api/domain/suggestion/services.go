package suggestion

import (
	"os"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	. "app/infra/aws"
	"fmt"
	"gopkg.in/validator.v2"
	mgo "gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"

)

type Message struct{
	Message string `json:"message" validate:"nonzero"`
	User string `json:"user" validate:"nonzero" ` 
}

type MongoDB struct {
	Server   string
	Database string
}

type Suggestion struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Message       string        `bson:"message" json:"message"`
	User          string        `bson:"user" json:"user"`
	Date          string        `bson:"date" json:"date"`
}

var db *mgo.Database

func CreateSuggestion(c *gin.Context) {

	var m Message

	c.BindJSON(&m)

	validationErr := m.VerifyRequired()

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": validationErr})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ConsulURL := os.Getenv("CONSUL_HOST")
	ConsulSQSURL := os.Getenv("CONSUL_SQS_URL_KV")
	ConsulSQSGROUP := os.Getenv("CONSUL_SQS_GROUP_KV")
		
	result, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	message := string(result)

	SQSSend(message, ConsulURL, ConsulSQSURL, ConsulSQSGROUP)

	c.JSON(http.StatusOK, gin.H{"response": true})
	return
}

func ListSuggestions(c *gin.Context) {

	COLLECTION := os.Getenv("MONGODB_COLLECTION_NAME")
	var suggestion Suggestion

	sug, listErr := suggestion.GetAll(COLLECTION)

	if listErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": listErr})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": sug})
	return
}

func (m Message) VerifyRequired() error {
	if err := validator.Validate(m); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (s *Suggestion) GetAll(COLLECTION string) ([]Suggestion, error) {
	var suggestion []Suggestion
	err := db.C(COLLECTION).Find(bson.M{}).All(&suggestion)
	return suggestion, err
}