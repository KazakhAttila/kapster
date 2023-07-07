package services

import (
	"github.com/kazakhattila/kapster"
	"github.com/kazakhattila/kapster/pkg/repositories"
)

type ResidentService struct {
	Repository repositories.Resident
}

func newResidentService(repo repositories.Resident) *ResidentService {
	return &ResidentService{
		Repository: repo}
}

func (r *ResidentService) Get() ([]resident.Resident, error) {
	return r.Repository.Get()

}

func (r *ResidentService) Refresh() (error) {
	return r.Repository.Refresh()
}
