package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	resident "github.com/kazakhattila/kapster"
)

type ResidentPostgres struct {
	db *sqlx.DB
}

func newResidentalPostgres(db *sqlx.DB) *ResidentPostgres {
	return &ResidentPostgres{db: db}

}

func (r *ResidentPostgres) Get() ([]resident.Resident, error) {
	db := r.db
	var returning []resident.Resident
	query := "SELECT * FROM dataas"
	err := db.Select(&returning, query)
	if err != nil {
		return nil, err
	}
	return returning, nil

}

// refresh function to send to the kapster.kz

func (r *ResidentPostgres) Refresh() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	c := http.Client{Timeout: 20 * time.Second}

	for i := 0; i < 2; i++ {
		endpoint := fmt.Sprintf("https://kapster.kz/api/products?city_id=%s&per_page=100", strconv.Itoa(i))
		resp, err := c.Get(endpoint)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var dev resident.Type1
		err = json.Unmarshal(body, &dev)

		datt := dev.Resident

		if err != nil {
			return (err)
		}

		residentColumns, _ := getColumns()
		query := fmt.Sprintf(`INSERT INTO %s %s VALUES %s;`, residentTable, residentColumns, getFormatted(datt))
		_, err = tx.Exec(query)
		if err != nil {
			fmt.Println(datt)
			return err
		}

		tx.Commit()

	}
	return nil

}
