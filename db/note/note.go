package note

import (
	"database/sql"
	"time"

	"github.com/5anthosh/nottu/db"
	"github.com/5anthosh/nottu/db/status"

	"github.com/5anthosh/mint"
)

//Create creates new note with title and content
func Create(database *sql.DB, title string, content string) (string, string, int, error) {
	if result := proccessInputs(title, content); result != ok {
		return emptyString, result, status.BadRequest, nil
	}
	id := db.GenerateUniqueID()
	created := time.Now()
	query := "INSERT INTO NOTE(NOTE_ID, TITLE, CONTENT, CREATED) VALUES(?, ?, ?, ?)"
	tx, err := database.Begin()
	if err == nil {
		_, err = tx.Exec(query, id, title, content, created)
		if err == nil {
			tx.Commit()
			return id, creationSuccess, status.CreatedSuccess, err
		}

	}
	tx.Rollback()
	return emptyString, creationErr, status.ServerError, mint.Traceable(err)

}

//Get gets list of notes
func Get(database *sql.DB) ([]Note, string, int, error) {
	query := "SELECT * FROM Note"
	rows, err := database.Query(query)
	if err != nil {
		return nil, retrieveErr, status.ServerError, mint.Traceable(err)
	}
	defer rows.Close()
	var notes = make([]Note, 0, 16)
	for rows.Next() {
		note := new(Note)
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.Created)
		if err != nil {
			return nil, retrieveErr, status.ServerError, mint.Traceable(err)
		}
		notes = append(notes, *note)
	}
	return notes, success, status.OK, err
}

//ByID gets a note by its id
func ByID(database *sql.DB, id string) (Note, string, int, error) {
	var note Note
	query := "SELECT * FROM Note WHERE NOTE_ID = ?"
	note = Note{}
	err := database.QueryRow(query, id).Scan(&note.ID, &note.Title, &note.Content, &note.Created)
	if err != nil {
		if err == sql.ErrNoRows {
			return note, notFound, status.NotFound, nil
		}
		return note, retrieveErrSingular, status.ServerError, mint.Traceable(err)
	}
	return note, success, status.OK, err
}

//Delete deletes the note by its id
func Delete(database *sql.DB, id string) (string, int, error) {
	_, result, code, err := ByID(database, id)
	if code != status.OK {
		return result, code, err
	}
	query := "DELETE FROM Note WHERE NOTE_ID = ?"
	_, err = database.Exec(query, id)
	if err != nil {
		return deletionErr, status.ServerError, mint.Traceable(err)
	}
	return success, status.DeletedSuccess, err
}

//Update updates the note with title or content or both
func Update(database *sql.DB, id string, title *string, content *string) (string, int, error) {
	_, result, code, err := ByID(database, id)
	if code != status.OK {
		return result, code, err
	}
	var query string
	if title != nil && content != nil {
		query = "UPDATE Note SET TITLE = ?, CONTENT = ? WHERE NOTE_ID = ?"
		_, err = database.Exec(query, *title, *content, id)
		if err != nil {
			return updationErr, status.ServerError, mint.Traceable(err)
		}
	} else if title != nil {
		query = "UPDATE Note SET TITLE = ? WHERE NOTE_ID = ?"
		_, err = database.Exec(query, *title, id)
		if err != nil {
			return updationErr, status.ServerError, mint.Traceable(err)
		}
	} else if content != nil {
		query = "UPDATE Note SET CONTENT=? WHERE NOTE_ID = ?"
		_, err = database.Exec(query, *content, id)
		if err != nil {
			return updationErr, status.ServerError, mint.Traceable(err)
		}
	}
	return updationSuccess, status.OK, err
}
