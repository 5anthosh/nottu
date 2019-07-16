package note

import (
	"database/sql"
	"nottu/db"
	"nottu/db/status"
	"time"
)

//Create creates new note with title and content
func Create(database *sql.DB, title *string, content *string) (string, string, int, error) {
	if result := proccessInputs(title, content); result != ok {
		return emptyString, result, status.BadRequest, nil
	}
	id := db.GenerateUniqueID()
	created := time.Now()
	query := "INSERT INTO NOTE(NOTE_ID, TITLE, CONTENT, CREATED) VALUES(?, ?, ?, ?)"
	tx, err := database.Begin()
	if err == nil {
		_, err = tx.Exec(query, id, *title, *content, created)
		if err == nil {
			tx.Commit()
			return id, creationSuccess, status.CreatedSuccess, err
		}

	}
	tx.Rollback()
	return emptyString, creationErr, status.ServerError, err

}
