package services


import( 

)

type Resident interface {
		Get() Resident
		Refresh() 
}


type ResidentSlug interface {
	 	Get(slug string) ResidentSlug
		Refresh(slug string) 
}


type Service struct {
		Resident
		ResidentSlug
}

func newService(repos *repositories.Repository) *Service{ 
			return &Service{ 	
					Res: newResidentalService(repos),
					ResSlug: newResidentalSlugService(repos),
			}

}