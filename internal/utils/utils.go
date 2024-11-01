package utils

import "twitter-clone/internal/models"

func CreateResponse(status int, message string) models.CreateResponse {
	return models.CreateResponse{
		Status:  status,
		Message: message,
	}
}
