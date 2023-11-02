package event

import "time"

type InputStatusUpdated struct {
	Name    string
	Payload interface{}
}

func NewInputStatusUpdated() *InputStatusUpdated {
	return &InputStatusUpdated{
		Name: "InputStatusUpdated",
	}
}

func (e *InputStatusUpdated) GetName() string {
	return e.Name
}

func (e *InputStatusUpdated) GetPayload() interface{} {
	return e.Payload
}

func (e *InputStatusUpdated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *InputStatusUpdated) GetDateTime() time.Time {
	return time.Now()
}
