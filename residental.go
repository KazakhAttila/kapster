package main

type Resident struct {
	ID                   int     `json:"id"`
	Uid                  string  `json:"uid"`
	Title                string  `json:"title"`
	Slug                 string  `json:"slug"`
	Category_id          string  `json:"category_id"`
	Subcategory_id       string  `json:"subcategory_id"`
	City_id              int     `json:"city_id"`
	Deadline_year        string  `json:"deadline_year"`
	Floors_min           int     `json:"floors_min"`
	Floors_max           int     `json:"floors_max"`
	Ceiling_min          float64 `json:"ceiling_min"`
	Ceiling_max          float64 `json:"ceiling_max"`
	Rooms_min            int     `json:"rooms_min"`
	Rooms_max            int     `json:"rooms_max"`
	Has_free_layout      bool    `json:"has_free_layout"`
	Has_hgf              bool    `json:"has_hgf"`
	Housing_class        string  `json:"housing_class"`
	Price_per_square_min int     `json:"price_per_square_min"`
	Price_min            int     `json:"price_min"`
	Kapster_verified     bool    `json:"kapster_verified"`
	Status               string  `json:"status"`
	Me_bookmarked        bool    `json:"me_bookmarked"`
	Image                string  `json:"image"`
	Published_at         string  `json:"published_at"`
	//The_rest map[string]interface{}

}
