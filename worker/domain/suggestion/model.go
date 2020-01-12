package suggestion

type Suggestion struct {
	Message       string        `bson:"message" json:"message"`
	User          string        `bson:"user" json:"user"`
	Date          string        `bson:"date" json:"date"`
}