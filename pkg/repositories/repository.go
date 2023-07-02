package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Residental interface {
	refresh()
	get()
}

type ResidentalSlug interface {
	refresh()
	get()
}

type Repository struct {
	Residental
	ResidentalSlug
}

func newRepository(db *sqlx.DB) *Repository { 

		return &Repository{ 
				Residental: NewResidental(db),
				ResidentalSlug: NewResidentalSlug(db),
		}
}
