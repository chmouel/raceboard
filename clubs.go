package raceboard

//Club
import (
	"database/sql"
	"time"
)

type Club struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

func (c *Club) Create(db *sql.DB) (err error) {
	row, err := db.Exec(`INSERT INTO Clubs(Name) VALUES(?)`, c.Name)
	if err != nil {
		return err
	}
	c.ID, err = row.LastInsertId()
	if err != nil {
		return err
	}
	err = c.GetByID(db)
	return err
}

func (c *Club) GetByName(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT ID, Name, CreatedAt FROM Clubs WHERE Name=?", c.Name).Scan(
		&c.ID, &c.Name, &c.CreatedAt)
	return
}

func (c *Club) GetByID(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT ID, Name, CreatedAt FROM Clubs WHERE ID=?", c.ID).Scan(
		&c.ID, &c.Name, &c.CreatedAt)
	return
}
