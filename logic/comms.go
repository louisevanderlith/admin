package logic

import (
	"github.com/louisevanderlith/mango"
)

type CommsObject struct {
	ID    int64
	Name  string
	Email string
	Phone string
	Body  string
}

func GetCommsMessages(instanceID string) ([]CommsObject, error) {
	resp, err := mango.GETMessage(instanceID, "Communication.API", "message")

	if err != nil {
		return []CommsObject{}, err
	}

	if resp.Failed() {
		return []CommsObject{}, resp
	}

	result := resp.Data.([]CommsObject)

	return result, nil
}