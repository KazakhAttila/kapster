package repositories

import(
	"github.com/jmoiron/sqlx"


) 
type  Resident struct{ 
	db *sqlx.DB
}

func newResidental(db *sqlx.DB) *Resident{ 
	return &Resident{ 
				db:db}

}
func (r *Resident) get() (Resident, error) { 
		

}

func (r *Resident) refresh() (error) { 


}

