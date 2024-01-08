package entity

type JobDefinition struct {
	ID               string   `bson:"id"`
	Service          string   `bson:"service"`
	Source           string   `bson:"source"`
	Context          string   `bson:"context"`
	BusinessOwner    []string `bson:"business_owner"`
	EventInputMethod string   `bson:"event_input_method"`
}
