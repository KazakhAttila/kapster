package services


import() 

type ResidentSlugService struct{ 
		Repository repositories.Repository
}


// freak one right now... 

// task is blind sighted -> write services, make assumptions and shit... 
// the same goes for the main... 
// and only then -> make the necessary reads and changes... 
// this is the goal of this session -> service + Main + repetitions where needed... + overall precise holes and questions around there...
// interface -> type -> functions and so on-> the determination of this shit... the determination of the service -> get some info from outside
// and pass the info to the repositories... 
// and do some calculations on the midair! 
// nado naiti chto to che to chto sozdat koe chto zdes'... 
// can't proofread function syntax or write my own or black box them right there... what the fuck? 
// 



func ( r *ResidentSlugService) Get(slug string) (kapster.ResidentSlug, error){ 
	repo := r.Repository()
	slugs, err:= repo.ResidentalSlug.Get(slug)
	return slugs, err


}	

func( r *ResidentSlugService) Refresh(slug string){ 		
	repo := r.Repository()
	slugs, err:= repo.ResidentalSlug.Get(slug)
	return slugs, err

}