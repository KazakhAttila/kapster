package repositories

import(
	"github.com/jmoiron/sqlx"
	"github.com/kazakhattila/kapster"
) 
type  Resident struct{ 
	db *sqlx.DB
}

func newResidental(db *sqlx.DB) *Resident{ 
	return &Resident{ 
				db:db}

}
func (r *Resident) Get() ([] kapster.Resident, error) { 
	db := r.db	
	var returning kapster.Resident
	query:= "SELECT * FROM dataas"
	err:=db.Get(&returning, query)
	if err!=nil{ 
			return nil, err
	}
	return returning, nil			

}

func (r *Resident) Refresh() (error) { 
			// zapros posilat'... 


}

