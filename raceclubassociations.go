package raceboard

//RaceClubAssociations
import "database/sql"

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
