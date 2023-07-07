package repositories

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"http/json"
	"github.com/jmoiron/sqlx"
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

/*
to check -> agree to make a map and banalize a lot of shit... 
trade up between this and freedom? what about old escape routes??? 

sorevnovanie vs freedom

primernaya formula -> one flashy interesting move and 5 banal moves... 

all 3 flashy VS freedom? how to do this trade???

*/

func (r *ResidentSlugPostgres) Refresh(slug string) (error){ 
		// get data from the db... 
		// slugs get fetched... 
		// add the slug fetched data into -> import into zhks () VALUES ()
		// make the abstract decision VS understand the fucking object rules VS identify all the points VS eliminate them one by one VS make a map(thus becoming 
		//bitch) VS understand the MACRO state and price of hesitance VS all the punishable paths VS revenge VS ugly scenes VS freedom and randomness
		// VS issledovaniye vseh owibok i pochemu oni owibki... VS unikal'nie bezduhovnie views... 
		/*
				1. Issledovatel' 
				2. Macro instruments... 
				3. speed and shit... 
				4. freedom and random and pohuism...
				5. one zero shot or many zero shots -> to check the capacity of the system's feedback giving...

				testing possibilities and shit -> why it bothers me so much? 
				maybe freedom is better? 
		*/
		
		
		query:=fmt.Sprintf(`SELECT FROM %s WHERE slug = %s`, residentSlugTable, slug)
		var slugOld[] resident.ResidentSlug

		err:= r.db.Get(&slugOld, query)
		if err!=nil{

		}
		// current trade up: 
		/*
				1. totally finishing this shit to send it to the test
				2. testing other shit
				3. 
		*/
	
		/*
				1. object from the resp -> http and so on -> objects and data structures cound
				2. connections count...
				3. this is a game for many more connections to happen...(???)
				4. relaxed vs insane vs freedom and relaxation trade... 

		*/
		var Client *http.Client
		endpoint:=fmt.Sprintf("https://kapster.kz/api/products/%s", slug)
		resp, err:= Client.Get(endpoint)
		if err!= nil{ 

		}
		body,err:= ioutil.ReadAll(resp.Body)
		var slugNew[] resident.ResidentSlug
		var addlist [10000] int

		for i:=0; i<len(slugNew);i++{
			addlist[i]=1
		}
		err = json.Unmarshal(body, &slug2)
		if err!=nil{ 

		}
		for i:=0;i<len(slugOld);i++{ 
			for j:=0;j<len(slugNew);j++{ 
					if slugNew[j]==slugOld[i]{ 
						addlist[j] = 0
					}
			} 
		}
		// exec function
		/*
				next trade is going out with the ustnii test + the black box... 
				with every shit here i am going to go out with the black box.
				write everything till the minor details...
				and go out with the black box... 

				after finishing -> test everything... 

				ALSO -> establish UNIT testing and other testing strats like zhashkevych... 
				dal'we -> channels and shit -> write the code x5

				tomorrow and so on-> we will train to bring one artery to the testing stage as fast as possible, even with defects.
				deffective trading... 

				oba varianta kogda imeesh' -> oba ostavit' na teh zhe urovnyah na kotorih stoyat... 
				ne nado userdstvovat'. nado rasslabit'sya... 


		*/
		for i:=0; i<len(slugNew);i++{
			if(addlist[i]==1){
				vals:=slugNew[i]
				vals = f
				query:=fmt.Sprintf(`INSERT INTO %s VALUES `, )
				r.db.Exec()
			}
		}
		if (err!=nil) { 
				fmt.Println(err.Error())
		}

		// trade up between banality VS establishing and noticing some difference vs making a bet vs memorizing some sequence vs freedom and random
		// banality or freedom?


		// first time -> creating code myself, creating data structures, testing 2-3 out of them, just like with mebel -> engineering and shit.
		// second time -> 
		// third time -> 
		// fourth time -> 
		// 
		// 
		// 
		// 
		//		


		/*
	c := &http.Client{Timeout: 1000 * time.Second}
	cc := NewClient(c)

	for ss := range dataa {
		Zhk := ZHK{}
		endpoint := fmt.Sprintf("https://kapster.kz/api/products/%s", dataa[ss].Slug)
		resp, err := cc.http.Get(endpoint)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
		}

		err = json.Unmarshal(body, &Zhk)
		if err != nil {
			panic(err)
		}
		DB.Create(&Zhk)


		*/

}

