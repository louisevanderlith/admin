package logic

import (
	"github.com/louisevanderlith/mango"
)

type LookupObj struct {
	ID          int64
	Name        string
	Description string
}

func GetCategories(instanceID string) ([]LookupObj, error) {
	return getLookup(instanceID, "category")
}

func GetManufacturers(instanceID string) ([]LookupObj, error) {
	return getLookup(instanceID, "manufacturer")
}

func GetModels(instanceID string) ([]LookupObj, error) {
	return getLookup(instanceID, "model")
}

func GetSubCategories(instanceID string) ([]LookupObj, error) {
	return getLookup(instanceID, "subcategory")
}

func getLookup(instanceID, controller string) ([]LookupObj, error) {
	result := []LookupObj{}
	err := mango.DoGET(&result, instanceID, "Things.API", controller)

	if err != nil {
		return result, err
	}

	return result, nil
}
