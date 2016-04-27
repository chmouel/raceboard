package raceboard

import "testing"

func init() {
	FixtureInitDB()
}

func TestRaceClubAssociations(t *testing.T) {
	raceName := "Race"
	raceLocation := "Location"
	clubName := "RaceClubAssociation"

	//rc = RaceClubAssociations{}

	r := Race{Name: raceName, Location: raceLocation}
	if err := r.Create(impl.DB); err != nil {
		t.Error(err)
		t.Fail()
	}

	c := Club{Name: clubName}
	if err := c.Create(impl.DB); err != nil {
		t.Error(err)
		t.Fail()
	}

	rc := RaceClubAssociations{RaceID: r.ID, ClubID: c.ID}
	if err := rc.Associate(impl.DB); err != nil {
		t.Error(err)
		t.Fail()
	}

	rc = RaceClubAssociations{RaceID: r.ID}
	clubs, err := rc.GetClubsAssociatedToRace(impl.DB)
	if err != nil {
		t.Error(err)
	}
	if clubs[0].Name != c.Name || clubs[0].ID != c.ID {
		t.Error("Cannot match")
	}
}
