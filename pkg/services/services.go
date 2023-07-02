package services


import( 

)
// netu i nevozmozhno nahooi bez horoshego servica pryamo was reshenia delat'? 
/*
reps: 
1. services services -> business logic, captures repositories, uses golang, used by handlers... -> 
services?? the picture or just the closed already patterns? uses closed patterns and doesn't know any shit... 
2. what is the form of services.
3. how the task can be divided.
4. what is independent from each other? 
5. new golang knowledge.
6. how many independent creatures are there??? 
7. what banal stuff must be repeated? 
objects OR links OR functions.
structures of them + interfaces.
data structures and information.
libraries and packages, external rules.
data flow endless + time of the flow!
memory spaces and places everyone is taking.
*/

// business logic -> new interfaces... 

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