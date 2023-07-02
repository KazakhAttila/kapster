package main

import (
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

// var DB *gorm.DB

// var err error

/*

id  uid  title  slug  category_id  subcategory_id  city_id  deadline_year
 deadline_quarter  location  floors_min  floors_max  ceiling_min  ceiling_max
  rooms_min  rooms_max  has_free_layout  purchase_methods/0  has_hgf  housing_class
   price_per_square_min  price_min  kapster_verified  status  me_bookmarked  image
   published_at  user_id  user/id  user/name  user/username  user/is_company  user/profile_verified
   user/avatar  images/0  images/1  images/2  images/3  images/4  images/5  plans/0/id  plans/0/rooms
     plans/0/area  plans/0/price  plans/1/id  plans/1/rooms  plans/1/area  plans/1/price  plans/2/id
	 plans/2/rooms  plans/2/area  plans/2/price  plans/3/id  plans/3/rooms  plans/3/area  plans/3/price
	  plans/4/id  plans/4/rooms  plans/4/area  plans/4/price  plans/5/id  plans/5/rooms  plans/5/area
	  plans/5/price  plans/6/id  plans/6/rooms  plans/6/area  plans/6/price  purchase_methods/1  plans/7/id
	  plans/7/rooms  plans/7/area  plans/7/price  purchase_methods/2  purchase_methods/3  images/6  plans/8/id
	    plans/8/rooms  plans/8/area  plans/8/price  plans/9/id  plans/9/rooms  plans/9/area  plans/9/price  plans/10/id
		 plans/10/rooms  plans/10/area  plans/10/price  plans/11/id  plans/11/rooms  plans/11/area  plans/11/price  plans/12/id
		  plans/12/rooms  plans/12/area  plans/12/price  plans/13/id  plans/13/rooms  plans/13/area  plans/13/price
		   plans/14/id  plans/14/rooms  plans/14/area  plans/14/price  plans/15/id  plans/15/rooms  plans/15/area
		    plans/15/price  plans/16/id  plans/16/rooms  plans/16/area  plans/16/price  plans/17/id  plans/17/rooms
			 plans/17/area  plans/17/price  plans/18/id  plans/18/rooms  plans/18/area  plans/18/price  plans/19/id
			  plans/19/rooms  plans/19/area  plans/19/price  plans/20/id  plans/20/rooms  plans/20/area  plans/20/price
			  plans/21/id  plans/21/rooms  plans/21/area  plans/21/price  plans/22/id  plans/22/rooms  plans/22/area
			    plans/22/price  images/7  images/8  images/9  images/10  images/11  images/12  images/13
				 purchase_methods/4  images/14  images/15  images/16

*/

// type Dataa struct {
// 	ID                   int     `json:"id"`
// 	Uid                  string  `json:"uid"`
// 	Title                string  `json:"title"`
// 	Slug                 string  `json:"slug"`
// 	Category_id          string  `json:"category_id"`
// 	Subcategory_id       string  `json:"subcategory_id"`
// 	City_id              int     `json:"city_id"`
// 	Deadline_year        string  `json:"deadline_year"`
// 	Floors_min           int     `json:"floors_min"`
// 	Floors_max           int     `json:"floors_max"`
// 	Ceiling_min          float64 `json:"ceiling_min"`
// 	Ceiling_max          float64 `json:"ceiling_max"`
// 	Rooms_min            int     `json:"rooms_min"`
// 	Rooms_max            int     `json:"rooms_max"`
// 	Has_free_layout      bool    `json:"has_free_layout"`
// 	Has_hgf              bool    `json:"has_hgf"`
// 	Housing_class        string  `json:"housing_class"`
// 	Price_per_square_min int     `json:"price_per_square_min"`
// 	Price_min            int     `json:"price_min"`
// 	Kapster_verified     bool    `json:"kapster_verified"`
// 	Status               string  `json:"status"`
// 	Me_bookmarked        bool    `json:"me_bookmarked"`
// 	Image                string  `json:"image"`
// 	Published_at         string  `json:"published_at"`
// 	//The_rest map[string]interface{}

// }
// type Type1 struct{
// 	Dat[] Dataa
// }
// type ZHK struct {
// 	Id               int     `json:"id"`
// 	Type1            string  `json:"_type"`
// 	Uid              string  `json:"uid"`
// 	Title            string  `json:"title"`
// 	Slug             string  `json:"slug"`
// 	Short            string  `json:"short"`
// 	Category_id      int     `json:"category_id"`
// 	Subcategory_id   int     `json:"subcategory_id"`
// 	Status           string  `json:"status"`
// 	Country_id       int     `json:"country_id"`
// 	City_id          int     `json:"city_id"`
// 	Published_at     string  `json:"published_at"`
// 	Edited_at        string  `json:"edited_at"`
// 	User_id          int     `json:"user_id"`
// 	Build_status     string  `json:"build_status"`
// 	Deadline_year    string  `json:"deadline_year"`
// 	Deadline_quarter int     `json:"deadline_quarter"`
// 	Latitude         float64 `json:"latitude"`
// 	Longitude        float64 `json:"longitude"`
// 	Floors_min       int     `json:"floors_min"`
// 	Floors_max       int     `json:"floors_max"`
// 	Ceiling_min      float64 `json:"ceiling_min"`
// 	Ceiling_max      float64 `json:"ceiling_max"`
// 	Rooms_min        int     `json:"rooms_min"`
// 	Rooms_max        int     `json:"rooms_max"`
// 	Has_free_layout  bool    `json:"has_free_layout"`
// }

// type Links struct {
// 	First string `json:"first"`
// 	Next  string `json:"next"`
// 	Last  string `json:"last"`
// 	Prev  string `json:"prev"`
// }
// type Meta struct {
// 	Current_page string `json:"current_page"`
// }

// type Client struct {
// 	http *http.Client
// }

// type type2 struct{

// }

// func NewClient(http *http.Client) *Client {

// 	return &Client{http}
// }

// func initialMigration() {
// 	dsn := "host=localhost user=postgres password=123 dbname=kapster port=5432 sslmode=disable"
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("kotakbas")
// 	}
// 	DB.AutoMigrate(&Dataa{})

// }
// func initialMigration2() {
// 	dsn := "host=localhost user=postgres password=123 dbname=kapster port=5432 sslmode=disable"
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("kotakbas")
// 	}
// 	DB.AutoMigrate(&ZHK{})
// }

// func initialMigration2() {
// 	dsn := "host=localhost user=postgres password=123 dbname=kapster port=5432 sslmode=disable"
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("kotakbas")
// 	}
// 	DB.AutoMigrate(&type2{})

// }
func (c *Client) FetchEverything1(city string) (*[]Dataa, error) {
	endpoint := fmt.Sprintf("https://kapster.kz/api/products?city_id=%s&per_page=100", city)
	resp, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	// dev.Dat = []Dataa{}
	// dev.Link = []Links{}
	// dev.Meta = []Meta{}
	var dev Type1
	err = json.Unmarshal(body, &dev)
	// Надо сломать эту хуйню
	var datt []Dataa = dev.Dat

	if err != nil {
		panic(err)
	}
	fmt.Println(dev)

	DB.Create(datt)
	return &datt, err
}

// return slugs available
func handle_first() {
	initialMigration()

	c := &http.Client{Timeout: 10 * time.Second}
	cc := NewClient(c)
	for i := 1; i < 19; i++ {
		qq, error := cc.FetchEverything1(strconv.Itoa(i))
		if error != nil && qq == nil {
			fmt.Println(error.Error())
		}

	}
	// var vals[] interface{}
}

// use the data structure provided from the first or what?
// zapros v db!!!
//
//
//

/*
goals:
	1. second make it possible -> any way -> use the database, use any way
	2. in the second -> make the new data structure and new request same shit...
	3. make the map of agreements and count them
	4. make the map of objects here and all the other layers suka...

*/



// automatize the shit!!
// then print koroche!

/*
 */
type Slug struct {
	Slug string
}

func handle_second() {

	initialMigration2()
	var dataa []Slug

	// my speech -> everything is banal as fuck, initial formula was torn apart and wasn't still achieved. look at this shit. everyone is weak.
	// na urovne rechi and macro -> reaching this(with additional automata theories and shit),
	// just to do some remigrations and adding some additional funcs -> macro and useful, even pictures and so on-> making an API!!!
	// first thing -> extract out of DB -> giving too much informaiton...

	DB.Table("dataas").Select("slug").Find(&dataa)
	fmt.Println()
	fmt.Println(dataa)

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

	}
}

// }

func main() {
	//handle_first()
	// dat := handle_first()
	// for i := range (*dat){
	// 		qq := *dat
	// 		x:=  qq[i]
	// 		for()
	// }

	// starts the program -> shit isn't enough here... what to do?
	// program is highly precise and there are too many choices.
	// goal is main + service + okonchatel'naya schema vezde... i potom zakonchit' hod i povtorit' svoi hod i svoi pobedi...
	// another mini goal -> spam of multiple independent yet objects, but possible to connect through right data types and data structures...
	// another mini goal -> leave the micro absurd calculations here -> why they are absurd without the connection?
	//

	// repo obyavlyaetsya! \
	// handlers obyavlayutsya
	// services also obyavlayutsya nahooi!
	// otsyuda berem Config nahooi i perevodim na yazik configov...
	// po suti dohuya nezavisimih drug ot druga veshei, nado ih vseh prosto snachala nazvat'.
	// main, services, handlers, repos, urls???, hanler, services, repos and their functions..
	// the goal is event driven shit!
	// pochemu pole zakonchilos'?

	repos := repositories.NewRepository()
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)

	server.run()
	handler.initRoutes()

	// 

	//handle_second()
}
