package repositories

import(
	"github.com/jmoiron/sqlx"
	"fmt"
	"github.com/kazakhattila/kapster"
) 
type  ResidentSlugPostgres struct{ 
	db *sqlx.DB
}

func NewResidentalSlug(db *sqlx.DB) *ResidentSlugPostgres{ 
	return &ResidentSlugPostgres{ 
				db:db}

}
func (r *ResidentSlugPostgres) Get(slug string) ([] resident.ResidentSlug, error){ 
	var result []resident.ResidentSlug
	query:= fmt.Sprintf("SELECT * FROM zhks WHERE slug = %s", slug) 
	if err:=r.db.Get(&result, query); err!=nil{
		return nil, err
	}
	
	
	return result, nil
}

func (r *ResidentSlugPostgres) Refresh(slug string) (resident.ResidentSlug, error){ 
		db := r.db
		query:= ""		
		db.Exec()
}

/*
*/
