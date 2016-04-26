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

type Club struct {
	ID        uint
	Name      string
	CreatedAt time.Time
}

type Race struct {
	ID        uint
	Location  string
	Name      string
	CreatedAt time.Time
}

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

func (i *Impl) CreateClub(name string) (int64, error) {
	result, err := i.DB.Exec(`INSERT INTO Clubs(Name) VALUES(?)`, name)
	if err != nil {
		log.Fatal(err)
	}
	return result.LastInsertId()
}

func (i *Impl) CreateRace(name, location string) (int64, error) {
	result, err := i.DB.Exec(`INSERT INTO Races(Name, location) VALUES(?, ?)`, name, location)
	if err != nil {
		log.Fatal(err)
	}
	return result.LastInsertId()
}

func (i *Impl) AssociateClubRace(clubID, raceID int64) (int64, error) {
	result, err := i.DB.Exec(`INSERT INTO RaceClubAssociations(clubID, raceID) VALUES(?, ?)`, clubID, raceID)
	if err != nil {
		log.Fatal(err)
	}
	return result.LastInsertId()
}

// GetClubByName
func (i *Impl) GetClubByName(name string) (club Club) {
	err := i.DB.QueryRow("SELECT ID, Name, CreatedAt FROM Clubs WHERE Name=?", name).Scan(
		&club.ID, &club.Name, &club.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// GetRaceByName
func (i *Impl) GetRaceByName(name string) (race Race) {
	err := i.DB.QueryRow("SELECT ID, Name, Location, CreatedAt FROM Races WHERE Name=?", name).Scan(
		&race.ID, &race.Name, &race.Location, &race.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (i *Impl) GetClubsAssociatedToRace(raceID int64) (ret []Club) {
	rows, err := i.DB.Query("SELECT c.ID, c.Name FROM RaceClubAssociations rc, Clubs c WHERE rc.raceID = ? AND rc.clubID = c.ID", raceID)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		club := Club{}
		err := rows.Scan(&club.ID, &club.Name)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, club)
	}
	return
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

	newclubID, err := i.CreateClub("Jaures")
	if err != nil {
		log.Fatal(err)
	}
	secondClubID, err := i.CreateClub("Abesses")
	if err != nil {
		log.Fatal(err)
	}

	newRaceID, err := i.CreateRace("Marathon", "Paris")
	if err != nil {
		log.Fatal(err)
	}

	_, err = i.AssociateClubRace(newclubID, newRaceID)
	if err != nil {
		log.Fatal(err)
	}
	_, err = i.AssociateClubRace(secondClubID, newRaceID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Getting all clubs associated to a race")
	ret := i.GetClubsAssociatedToRace(newRaceID)
	for _, v := range ret {
		fmt.Println(v.Name)
	}

	fmt.Println(i.GetClubByName("Jaures"))
	fmt.Println(i.GetRaceByName("Marathon"))

}
