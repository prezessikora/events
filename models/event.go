package models

import (
	"errors"
	"time"

	"github.com/prezessikora/events/db"
)

type Event struct {
	ID          int64
	UserID      int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}

func (event *Event) Delete() error {
	updateSql := "DELETE FROM events WHERE id=?"
	result, err := db.DB.Exec(updateSql, event.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New("expected to affect one row, affected")
	}
	return nil

}

func (e *Event) Update() error {
	updateSql := "UPDATE events SET name=? ,description=?,location=?,dateTime=? WHERE id=?"
	result, err := db.DB.Exec(updateSql, e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New("expected to affect one row, affected")
	}
	return nil

}

func (e *Event) Save() error {

	insertSql := `INSERT INTO events (name,description,location,dateTime,user_id) 
	VALUES (?,?,?,?,?)`
	stmt, err := db.DB.Prepare(insertSql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	e.ID, err = result.LastInsertId()
	return err
}

func GetEventById(id int) (*Event, error) {
	selectSql := `SELECT * FROM events WHERE id=?`
	row := db.DB.QueryRow(selectSql, id)
	event := Event{}
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func GetAll() ([]Event, error) {
	selectSql := `SELECT * FROM events`
	rows, err := db.DB.Query(selectSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return events, nil

}
