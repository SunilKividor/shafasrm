package pgrepo

import (
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
)

func (dbClient *PGRepo) StoreNewPhotoKey(userID uuid.UUID, req models.PhotoObject) error {
	client := dbClient.PostgresDBClient

	sql := `INSERT INTO user_photos (user_id,photo_key,is_primary) VALUES($1,$2,$3)`

	_, err := client.Exec(sql, userID, req.Key, req.IsPrimary)
	return err
}

func (dbClient *PGRepo) GetPhotos(userID uuid.UUID) ([]models.PhotoObject, error) {
	client := dbClient.PostgresDBClient
	var photos []models.PhotoObject
	sql := `SELECT photo_key,is_primary FROM user_photos WHERE user_id = $1`

	rows, err := client.Query(sql, userID)
	if err != nil {
		return photos, err
	}

	for rows.Next() {
		var photo models.PhotoObject

		if err := rows.Scan(&photo.Key, &photo.IsPrimary); err != nil {
			return photos, err
		}

		photos = append(photos, photo)
	}

	if err := rows.Err(); err != nil {
		return photos, err
	}
	return photos, err
}
