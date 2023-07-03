package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"github.com/kazakhattila/kapster"
	"github.com/jmoiron/sqlx"
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
	err := db.Get(&returning, query)
	if err != nil {
		return nil, err
	}
	return returning, nil

}

func (r *ResidentPostgres) Refresh() error {
	tx, err := r.db.Begin()
	if err!=nil{ 
		return err
	}

	c := &http.Client{Timeout: 10 * time.Second}

	for i := 0; i < 19; i++ {
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


		// dev.Dat = []Dataa{}
		// dev.Link = []Links{}
		// dev.Meta = []Meta{}
		var dev resident.Type1
		err = json.Unmarshal(body, &dev)
		// Надо сломать эту хуйню
		var datt []resident.Resident = dev.Resident

		if err != nil {
			return (err)
		}
		
		fmt.Println(dev)


		return nil

	}
}
