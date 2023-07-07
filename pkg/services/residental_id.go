package services


import(
	"github.com/kazakhattila/kapster/pkg/repositories"
	"github.com/kazakhattila/kapster"
) 

type ResidentSlugService struct{ 
		Repository repositories.ResidentSlug
}

func newResidentSlugService(repo repositories.ResidentSlug) *ResidentSlugService { 
				return &ResidentSlugService{ 
						Repository: repo}
}




func ( r *ResidentSlugService) Get(slug string) ([] resident.ResidentSlug, error){ 
	return r.Repository.Get(slug)



}	

func( r *ResidentSlugService) Refresh(slug string) (error){ 		
		return r.Repository.Refresh(slug)
}