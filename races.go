package raceboard

//Race
import (
	"database/sql"
	"time"
)

type Race struct {
	ID        int64
	Location  string
	Name      string
	CreatedAt time.Time
}

func (r *Race) Create(db *sql.DB) (err error) {
	row, err := db.Exec(`INSERT INTO Races(Name, location) VALUES(?, ?)`, r.Name, r.Location)
	if err != nil {
		return
	}

	if r.ID, err = row.LastInsertId(); err != nil {
		return
	}

	err = r.GetByID(db)
	return err
}

func (r *Race) GetByName(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT ID, Name, Location, CreatedAt FROM Races WHERE Name=? AND Location=?", r.Name, r.Location).Scan(
		&r.ID, &r.Name, &r.Location, &r.CreatedAt)
	return
}

func (r *Race) GetByID(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT ID, Name, Location, CreatedAt FROM Races WHERE ID=?", r.ID).Scan(
		&r.ID, &r.Name, &r.Location, &r.CreatedAt)
	return
}
