package database

import (
	"graphics-app-backend/src/models"
)

func InsertImage(userId string, url string) (*models.Image, error) {
	db, err := createConnection()
	if err != nil {
		return nil, err
	}
	var pk int
	err = db.QueryRow("INSERT INTO public.image (user_id, url) VALUES ($1, $2) RETURNING pk;", userId, url).Scan(&pk)
	if err != nil {
		return nil, err
	}
	return &models.Image{Pk: pk, UserId: userId, Url: url}, nil
}
