package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/kazakhattila/kapster"
)

type Resident interface {
	Refresh(error)
	Get() ([]resident.Resident, error)
}

type ResidentSlug interface {
	Refresh(slug string) (error)
	Get(slug string) ([] resident.ResidentSlug, error)
}

type Repository struct {
	Resident
	ResidentSlug
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Resident:     newResidentalPostgres(db),
		ResidentSlug: NewResidentalSlug(db),
	}
}
