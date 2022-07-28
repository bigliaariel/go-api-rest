// CRUD note
package models

import (
	"api/config"
	"api/entities"
	"fmt"
)

type NoteModel struct {
}

// FindAll show all notes from table note
func (*NoteModel) FindAll() ([]entities.Note, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select id,title,body,color from note;")

		if err2 != nil {
			return nil, err2
		} else {
			var notes []entities.Note
			for rows.Next() {
				var note entities.Note
				rows.Scan(&note.Id, &note.Title, &note.Body, &note.Color)
				notes = append(notes, note)

			}
			rows.Close()

			db.Close()
			return notes, nil
		}

	}

}

// Add a note
func (*NoteModel) Add(title string, note string, color int) (err error) {
	db, err := config.GetDB()

	if err != nil {
		return err
	} else {
		s := `INSERT INTO note  VALUES (null,'%s','%s',%d)`
		q := fmt.Sprintf(s, title, note, color)
		_, err2 := db.Exec(q)
		defer db.Close()
		if err2 != nil {
			return err
		}
	}
	return nil
}

// Remove delete a note
func (*NoteModel) Remove(id int32) (err error) {
	db, err := config.GetDB()

	if err != nil {
		return err
	} else {
		s := `DELETE FROM note WHERE id=%d`
		q := fmt.Sprintf(s, id)
		_, err2 := db.Exec(q)
		defer db.Close()
		if err2 != nil {
			return err
		}
	}
	return nil
}

// Update the note from id
func (*NoteModel) Update(id int32, title string, note string, color int32) (err error) {
	db, err := config.GetDB()

	if err != nil {
		return err
	} else {
		s := `UPDATE note  SET title='%s',body='%s',color=%d WHERE id=%d`
		q := fmt.Sprintf(s, title, note, color, id)
		_, err2 := db.Exec(q)
		defer db.Close()
		if err2 != nil {
			return err
		}
	}
	return nil
}

// Find a note from id
func (*NoteModel) Find(id int32) (*entities.Note, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		query := fmt.Sprintf("select id,title,body,color from note WHERE id=%d;", id)
		rows, err2 := db.Query(query)

		if err2 != nil {

			return &entities.Note{}, err2
		} else {
			var notes []entities.Note
			for rows.Next() {
				var note entities.Note
				rows.Scan(&note.Id, &note.Title, &note.Body, &note.Color)
				notes = append(notes, note)

			}
			rows.Close()
			db.Close()
			if len(notes) > 0 {
				return &notes[0], nil
			} else {
				return &entities.Note{}, nil
			}
		}

	}
}
