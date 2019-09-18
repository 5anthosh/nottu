package note

import (
	"database/sql"
	"time"

	"github.com/5anthosh/nottu/core/database"
	"github.com/5anthosh/nottu/core/status"
	"github.com/5anthosh/oops"
)

//Create creates new note with title and content
func Create(DB *sql.DB, title string, content string) (string, string, int, error) {
	if result := proccessInputs(title, content); result != ok {
		return emptyString, result, status.BadRequest, nil
	}
	id := database.GenerateUniqueID()
	created := time.Now()
	query := "INSERT INTO NOTE(NOTE_ID, TITLE, CONTENT, CREATED, UPDATED) VALUES(?, ?, ?, ?, ?)"
	tx, err := DB.Begin()
	if err == nil {
		_, err = tx.Exec(query, id, title, content, created, created)
		if err == nil {
			tx.Commit()
			return id, creationSuccess, status.CreatedSuccess, err
		}

	}
	tx.Rollback()
	return emptyString, creationErr, status.ServerError, oops.T(err)

}

//Get gets list of notes
func Get(DB *sql.DB) ([]Note, string, int, error) {
	query := "SELECT * FROM Note"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, retrieveErr, status.ServerError, oops.T(err)
	}
	defer rows.Close()
	var notes = make([]Note, 0, 16)
	for rows.Next() {
		note := new(Note)
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.Created, &note.Updated)
		if err != nil {
			return nil, retrieveErr, status.ServerError, oops.T(err)
		}
		notes = append(notes, *note)
	}
	return notes, success, status.OK, err
}

//ByID gets a note by its id
func ByID(DB *sql.DB, id string) (Note, string, int, error) {
	var note Note
	query := "SELECT * FROM Note WHERE NOTE_ID = ?"
	note = Note{}
	err := DB.QueryRow(query, id).Scan(&note.ID, &note.Title, &note.Content, &note.Created, &note.Updated)
	if err != nil {
		if err == sql.ErrNoRows {
			return note, notFound, status.NotFound, nil
		}
		return note, retrieveErrSingular, status.ServerError, oops.T(err)
	}
	return note, success, status.OK, err
}

//Delete deletes the note by its id
func Delete(DB *sql.DB, id string) (string, int, error) {
	_, result, code, err := ByID(DB, id)
	if code != status.OK {
		return result, code, err
	}
	query := "DELETE FROM Note WHERE NOTE_ID = ?"
	_, err = DB.Exec(query, id)
	if err != nil {
		return deletionErr, status.ServerError, oops.T(err)
	}
	return success, status.DeletedSuccess, err
}

//Update updates the note with title or content or both
func Update(DB *sql.DB, id string, title *string, content *string) (string, int, error) {
	_, result, code, err := ByID(DB, id)
	if code != status.OK {
		return result, code, err
	}
	var query string
	now := time.Now()
	if title != nil && content != nil {
		query = "UPDATE Note SET TITLE = ?, CONTENT = ?, UPDATED = ? WHERE NOTE_ID = ?"
		_, err = DB.Exec(query, *title, *content, now, id)
		if err != nil {
			return updationErr, status.ServerError, oops.T(err)
		}
	} else if title != nil {
		query = "UPDATE Note SET TITLE = ?, UPDATED = ? WHERE NOTE_ID = ?"
		_, err = DB.Exec(query, *title, now, id)
		if err != nil {
			return updationErr, status.ServerError, oops.T(err)
		}
	} else if content != nil {
		query = "UPDATE Note SET CONTENT = ?, UPDATED = ? WHERE NOTE_ID = ?"
		_, err = DB.Exec(query, *content, now, id)
		if err != nil {
			return updationErr, status.ServerError, oops.T(err)
		}
	}
	return updationSuccess, status.OK, err
}
