package services

import "example.com/urlibre/models"

type UrlService interface {
	CreateUrl(*models.Url) error
	GetUser(*string) (*models.Url, error)
	DeleteUrl(*string) error
}
