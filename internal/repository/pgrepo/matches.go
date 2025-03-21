package pgrepo

import (
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
)

func (dbClient *PGRepo) CreateNewMatch(user_id uuid.UUID, match models.Match) error {
	db := dbClient.PostgresDBClient

	smt := `INSERT INTO matches (user_id_1,user_id_2) VALUES ($1,$2)`
	_, err := db.Exec(smt, user_id, match.UserID)
	return err
}

func (dbClient *PGRepo) GetMatches(user_id uuid.UUID) (models.Matches, error) {
	db := dbClient.PostgresDBClient
	var matches models.Matches
	smt := `SELECT 
    			CASE 
        			WHEN user_id_1 = $1 THEN user_id_2 
        			ELSE user_id_1 
    			END 
    			AS user_ids
			FROM matches
			WHERE user_id_1 = $1
    			OR user_id_2 = $1 
			ORDER BY created_at DESC
		`
	rows, err := db.Query(smt, user_id)
	if err != nil {
		return matches, err
	}
	defer rows.Close()

	for rows.Next() {
		var matchID uuid.UUID
		if err := rows.Scan(&matchID); err != nil {
			return matches, err
		}
		matches.UserIDs = append(matches.UserIDs, matchID)
	}

	if err := rows.Err(); err != nil {
		return matches, err
	}
	return matches, err
}
