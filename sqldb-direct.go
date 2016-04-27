// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2016-04-26 19:49:37
//
// run snippet with Ctl-Return

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

//Club
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
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c *Club) GetByID(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT ID, Name, CreatedAt FROM Clubs WHERE ID=?", c.ID).Scan(
		&c.ID, &c.Name, &c.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	return
}

//Race
type Race struct {
	ID        int64
	Location  string
	Name      string
	CreatedAt time.Time
}

func (r *Race) Create(db *sql.DB) (err error) {
	row, err := db.Exec(`INSERT INTO Races(Name, location) VALUES(?, ?)`, r.Name, r.Location)
	if err != nil {
		log.Fatal(err)
	}

	r.ID, err = row.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	err = r.GetByID(db)
	return err
}

func (r *Race) GetByName(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT ID, Name, Location, CreatedAt FROM Races WHERE Name=?", r.Name).Scan(
		&r.ID, &r.Name, &r.Location, &r.CreatedAt)
	return
}

func (r *Race) GetByID(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT ID, Name, Location, CreatedAt FROM Races WHERE ID=?", r.ID).Scan(
		&r.ID, &r.Name, &r.Location, &r.CreatedAt)
	return
}

//RaceClubAssociations
type RaceClubAssociations struct {
	ID     int64
	RaceID int64
	ClubID int64
}

// Associate race to club by ID
func (rca *RaceClubAssociations) Associate(db *sql.DB) (err error) {
	_, err = db.Exec(`INSERT INTO RaceClubAssociations(clubID, raceID) VALUES(?, ?)`, rca.ClubID, rca.RaceID)
	return
}

// GetClubsAssociatedToRace get all clubs associated to race via the Race ID
func (rca *RaceClubAssociations) GetClubsAssociatedToRace(db *sql.DB) (ret []Club, err error) {
	rows, err := db.Query("SELECT c.ID, c.Name FROM RaceClubAssociations rc, Clubs c WHERE rc.raceID = ? AND rc.clubID = c.ID", rca.RaceID)
	if err != nil {
		return
	}
	for rows.Next() {
		club := Club{}
		err = rows.Scan(&club.ID, &club.Name)
		if err != nil {
			return
		}
		ret = append(ret, club)
	}
	return
}

// DB Impl
type Impl struct {
	DB         *sql.DB
	DBLocation string
}

func (i *Impl) InitDB() {
	var err error
	i.DB, err = sql.Open("sqlite3", i.DBLocation)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalf("Got an error when connecting to database, the error is '%v'", err)
	}
}

func (i *Impl) CreateSchema() {
	f, err := os.Open("tables.sql")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	tables := string(b)
	_, err = i.DB.Exec(tables)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i := Impl{DBLocation: "/tmp/screw.db"}

	cmd := exec.Command("rm", "-f", i.DBLocation)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	i.InitDB()
	i.CreateSchema()

	jauresClub := Club{Name: "Jaures"}
	if err = jauresClub.Create(i.DB); err != nil {
		log.Fatal(err)
	}

	abesseClub := Club{Name: "Abesses"}
	if err := abesseClub.Create(i.DB); err != nil {
		log.Fatal(err)
	}

	marathonRace := Race{Name: "Marathon", Location: "Paris"}
	if err := marathonRace.Create(i.DB); err != nil {
		log.Fatal(err)
	}

	associaTion := RaceClubAssociations{ClubID: jauresClub.ID, RaceID: marathonRace.ID}
	if err := associaTion.Associate(i.DB); err != nil {
		log.Fatal(err)
	}

	associaTion = RaceClubAssociations{ClubID: abesseClub.ID, RaceID: marathonRace.ID}
	if err := associaTion.Associate(i.DB); err != nil {
		log.Fatal(err)
	}

	var ret = []Club{}
	fmt.Println("Getting all clubs associated to a race: " + marathonRace.Name)
	associaTions := RaceClubAssociations{RaceID: marathonRace.ID}
	if ret, err = associaTions.GetClubsAssociatedToRace(i.DB); err != nil {
		log.Fatal(err)
	}
	for _, v := range ret {
		fmt.Println("\t" + v.Name)
	}

	jaures := Club{Name: "Jaures"}
	if err = jaures.GetByName(i.DB); err != nil {
		log.Fatal(err)
	}

	var r = Race{Name: "Marathon", Location: "Paris"}
	if err = r.GetByName(i.DB); err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Name, r.ID, r.CreatedAt)
}
