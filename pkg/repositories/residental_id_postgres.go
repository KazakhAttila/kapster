package repositories

import(
	"github.com/jmoiron/sqlx"
	"fmt"
	"github.com/kazakhattila/kapster"
) 
type  ResidentSlugPostgres struct{ 
	db *sqlx.DB
}

func NewResidentalSlug(db *sqlx.DB) *ResidentSlugPostgres{ 
	return &ResidentSlugPostgres{ 
				db:db}

}
func (r *ResidentSlugPostgres) Get(slug string) ([] resident.ResidentSlug, error){ 
	var result []resident.ResidentSlug
	query:= fmt.Sprintf("SELECT * FROM zhks WHERE slug = %s", slug) 
	if err:=r.db.Get(&result, query); err!=nil{
		return nil, err
	}
	
	
	return result, nil
}

func (r *ResidentSlugPostgres) Refresh(slug string) (resident.ResidentSlug, error){ 
		db := r.db
		query:= ""		
		db.Exec()
}

/*
time out suka: 
right now task is -> this or other random approahes way -> just to get to the main ...
get the services, repos and handlers work together in real time, just as the fucking main... 

then get to the FUCKING ZAPROSI AND FETCHING THEM REAL TIME!!! 
also CHANNELS AND DOCKER! 
tomorrow ASIK SOZVON!!!! 

current plan and what was done -> little cosmetic changes and some slurs or some micro mistakes out there... pzdc
pohui. prostim i zabudem? 
pohui dal'we same tempo + some new microinfos

koroche was -> VSyu chernuyu rabotu. 
dazhe channels nahooi dolzhen bit' zdes' -> status of the request, pohui nahui! 
repos ideal'no. 
dal'we -> services ideal'no. 
dal'we -> handlers tozhe ideal'no suka. 

*/