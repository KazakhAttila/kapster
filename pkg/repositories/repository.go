package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Resident interface {
	Refresh()
	Get()
}

type ResidentSlug interface {
	Refresh()
	Get()
}

type Repository struct {
	Resident
	ResidentSlug
}

func newRepository(db *sqlx.DB) *Repository { 

		return &Repository{ 
				Resident: NewResidental(db),
				ResidentSlug: NewResidentalSlug(db),
		}
}
