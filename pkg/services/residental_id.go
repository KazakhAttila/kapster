package services


import() 

type ResidentSlugService struct{ 
		Repository repositories.Repository
}

func newResidentSlugService(repo *repositories.Repository) *ResidentSlugService { 
				return &ResidentSlugService{ 
						Repository: repo}
}




func ( r *ResidentSlugService) Get(slug string) ([] kapster.ResidentSlug, error){ 
	return r.Repository.ResidentSlug.Get(slug)



}	

func( r *ResidentSlugService) Refresh(slug string){ 		
		return r.Repository.ResidentSlug.Refresh()
}