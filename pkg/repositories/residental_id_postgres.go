package repositories

import(
	"github.com/jmoiron/sqlx"


) 
type  ResidentSlug struct{ 
	db *sqlx.DB
}

func NewResidentalSlug(db *sqlx.DB) *ResidentSlug{ 
	return &ResidentSlug{ 
				db:db}

}
func (r *ResidentSlug) getSlug (slug string) (kapster.residental_individual, error){ 
	db := r.db	
	query:= "SELECT * FROM table"
	db.Exec(query)
}

func refreshSlug(slug string) (kapster.residental_individual, slug string){ 
		db := r.db
		query:= ""	
		db.Exec()
}

