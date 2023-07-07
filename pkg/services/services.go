package services


import( 
	"github.com/kazakhattila/kapster"
	"github.com/kazakhattila/kapster/pkg/repositories"

)

type Resident interface {
		Get() ([] resident.Resident, error)
		Refresh() (error) 
}


type ResidentSlug interface {
	 	Get(slug string) ([] resident.ResidentSlug, error)
		Refresh(slug string) (error)
}


type Service struct {
		Resident
		ResidentSlug
}

func newService(repos *repositories.Repository) *Service{ 
			return &Service{ 	
				Resident: newResidentService(repos.Resident),
				ResidentSlug: newResidentSlugService(repos.ResidentSlug),
			}

}