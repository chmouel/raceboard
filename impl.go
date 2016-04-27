package raceboard

// DB Impl
import (
	"database/sql"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Impl struct {
	DB         *sql.DB
	DBLocation string
}

func (i *Impl) InitDB() (err error) {
	i.DB, err = sql.Open("sqlite3", i.DBLocation)
	return
}

func (i *Impl) CreateSchema() (err error) {
	// TODO path
	f, err := os.Open("tables.sql")
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	tables := string(b)
	_, err = i.DB.Exec(tables)
	return
}
