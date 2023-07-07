package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jmoiron/sqlx"
	resident "github.com/kazakhattila/kapster"
)

type ResidentSlugPostgres struct {
	db *sqlx.DB
}

func NewResidentalSlug(db *sqlx.DB) *ResidentSlugPostgres {
	return &ResidentSlugPostgres{
		db: db}

}
func (r *ResidentSlugPostgres) Get(slug string) ([]resident.ResidentSlug, error) {
	var result []resident.ResidentSlug
	query := fmt.Sprintf("SELECT * FROM zhks WHERE slug = %s", slug)
	if err := r.db.Get(&result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ResidentSlugPostgres) Refresh(slug string) error {

	query := fmt.Sprintf(`SELECT FROM %s WHERE slug = %s`, residentSlugTable, slug)
	var slugOld []resident.ResidentSlug

	err := r.db.Get(&slugOld, query)
	if err != nil {

	}

	var Client *http.Client
	endpoint := fmt.Sprintf("https://kapster.kz/api/products/%s", slug)
	resp, err := Client.Get(endpoint)
	if err != nil {

	}
	body, err := ioutil.ReadAll(resp.Body)
	var slugNew []resident.ResidentSlug
	var addlist [10000]int

	for i := 0; i < len(slugNew); i++ {
		addlist[i] = 1
	}
	err = json.Unmarshal(body, &slugNew)
	if err != nil {

	}
	for i := 0; i < len(slugOld); i++ {
		for j := 0; j < len(slugNew); j++ {
			if slugNew[j] == slugOld[i] {
				addlist[j] = 0
			}
		}
	}
	_, columns := getColumns()
	for i := 0; i < len(slugNew); i++ {
		if addlist[i] == 1 {
			vals := slugNew[i]
			valsstring := getFormatted(vals)
			query := fmt.Sprintf(`INSERT INTO %s VALUES %s`, columns, valsstring)
			_, err = r.db.Exec(query)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
