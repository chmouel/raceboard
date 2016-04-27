package raceboard

import "testing"

func init() {
	FixtureInitDB()
}

func TestRace(t *testing.T) {
	name := "Race"
	location := "Location"
	assignedid := int64(1)

	r := Race{Name: name, Location: location}
	if err := r.Create(impl.DB); err != nil {
		t.Error(err)
	}

	r = Race{ID: assignedid}
	if err := r.GetByID(impl.DB); err != nil {
		t.Error(err)
	}
	if r.ID != assignedid || r.Location != location || r.Name != name {
		t.Fail()
	}

	r = Race{Name: name, Location: location}
	if err := r.GetByName(impl.DB); err != nil {
		t.Error(err)
	}
	if r.ID != assignedid || r.Location != location || r.Name != name {
		t.Fail()
	}

}
