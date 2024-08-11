package utils

import "golang-chapter-42/model"

type Response struct {
	Status  int
	Message string
	Data    model.User
	Error   string
}
