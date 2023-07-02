package repositories

import(
	"github.com/jmoiron/sqlx"
	"fmt"

) 
type  ResidentSlug struct{ 
	db *sqlx.DB
}

func NewResidentalSlug(db *sqlx.DB) *ResidentSlug{ 
	return &ResidentSlug{ 
				db:db}

}
func (r *ResidentSlug) getSlug (slug string) (kapster.ResidentSlug, error){ 
	db := r.db	
	query:= fmt.Sprintf("SELECT * FROM zhks WHERE slug = %s", slug) 
	result, err:=db.Get(query)
	if err!=nil{ 
			return nil, err
	}
	return result, err
}

func refreshSlug(slug string) (kapster.ResidentSlug, slug string){ 
		db := r.db
		query:= ""		
		db.Exec()
}

