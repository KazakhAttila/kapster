package services


import()

type ResidentService struct{ 

			Repository repositories.Resident
}

func newResidentService(repo *repositories.Repository) *ResidentService { 
	return &ResidentService{ 
			Repository: repo}
}


func (r *ResidentService) Get() ([] kapster.Resident) { 
		return r.Repository.Resident.Get()


}

func (r *ResidentService) Refresh(){ 
		return r.Repository.Resident.Refresh()
}