package services

import (
	"main.go/models"
)

// CreateRequest : Function to return a object of Request type.
func CreateRequest(url string, items_count int32) *models.Request {
	return &models.Request{
		Url:         url,
		Items_Count: items_count,
	}
}
