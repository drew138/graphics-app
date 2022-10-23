package database

import (
	"graphics-app-backend/src/models"
)

func SelectImage(imageId int, userId string) (*models.Image, error) {
	db, err := createConnection()
	if err != nil {
		return nil, err
	}
	var pk int
	var url string
	err = db.QueryRow(`
		SELECT pk, url 
		FROM public.image 
		WHERE (pk = $1) 
		AND (user_id = $2);`,
		imageId,
		userId).Scan(&pk, &url)
	if err != nil {
		return nil, err
	}
	return &models.Image{Pk: pk, UserId: userId, Url: url}, nil
}

func SelectPaginatedUserImages(userId string, from, to int) ([]*models.Image, error) {

	db, err := createConnection()
	if err != nil {
		return nil, err
	}
	if from < 0 {
		from = 0
	}
	if to < from {
		to = from
	}
	rows, err := db.Query(`
		SELECT pk, url
		FROM public.image 
		WHERE (user_id = $1) LIMIT $2 
		OFFSET $3;`,
		userId, to-from, from)
	if err != nil {
		return nil, err
	}
	var images []*models.Image
	var pk int
	var url string
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&pk, &url); err != nil {
			return nil, err
		}
		images = append(images, &models.Image{Pk: pk, UserId: userId, Url: url})
	}
	return images, nil
}
