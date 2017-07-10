package yelp

// Category describes a business.
type Category struct {
	Alias string `json:"alias"`
	Title string `json:"title"`
}

// Business defines a business returned by the Yelp API.
type Business struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	ImageURL     string      `json:"image_url"`
	IsClaimed    bool        `json:"is_claimed"`
	IsClosed     bool        `json:"is_closed"`
	URL          string      `json:"url"`
	Price        string      `json:"price"`
	Rating       float64     `json:"rating"`
	ReviewCount  int64       `json:"review_count"`
	Phone        string      `json:"phone"`
	Photos       []string    `json:"photos"`
	Categories   []Category  `json:"categories"`
	Coodinates   Coordinates `json:"coordinates"`
	Location     Location    `json:"location"`
	Transactions []string    `json:"transactions"`

	// Only in search result
	DisplayPhone string  `json:"display_phone"`
	Distance     float64 `json:"distance"`
}
