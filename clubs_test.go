package raceboard

import "testing"

func init() {
	FixtureInitDB()
}

func TestClub(t *testing.T) {
	name := "Club"
	assignedid := int64(1)

	c := Club{Name: name}
	if err := c.Create(impl.DB); err != nil {
		t.Fail()
	}

	c = Club{ID: assignedid}
	if err := c.GetByID(impl.DB); err != nil {
		t.Fail()
	}
	if c.ID != assignedid || c.Name != name {
		t.Fail()
	}

	c = Club{Name: name}
	if err := c.GetByName(impl.DB); err != nil {
		t.Fail()
	}
	if c.ID != assignedid || c.Name != name {
		t.Fail()
	}

}
