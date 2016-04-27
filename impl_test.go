package raceboard

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	impl     Impl
	tempFile *os.File
)

const totalTables int = 3

func init() {
	FixtureInitDB()
}

func FixtureInitDB() {
	var err error
	tempFile, err = ioutil.TempFile(os.TempDir(), ".tmp.db.XXXXX")
	if err != nil {
		log.Fatal(err)
	}

	impl = Impl{DBLocation: tempFile.Name()}
	if err := impl.InitDB(); err != nil {
		log.Fatal(err)
	}
	if err := impl.CreateSchema(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	exitcode := m.Run()
	if err := tempFile.Close(); err != nil {
		log.Fatal(err)
	}
	if err := os.Remove(tempFile.Name()); err != nil {
		log.Fatal(err)
	}
	os.Exit(exitcode)
}

func TestCreatedDB(t *testing.T) {
	if _, err := os.Stat(impl.DBLocation); os.IsNotExist(err) {
		t.Errorf("DB File %s didn't get created", impl.DBLocation)
	}
}

func TestCreatedSchema(t *testing.T) {
	var numberTables int
	var sql string = `SELECT count(*) FROM sqlite_master WHERE type = 'table' AND name != 'android_metadata' AND name != 'sqlite_sequence'`

	if err := impl.DB.QueryRow(sql).Scan(&numberTables); err != nil {
		t.Error(err)
	}
	if numberTables != totalTables {
		t.Errorf("Not enough created schemas: %d != %d", numberTables, totalTables)
	}
}
