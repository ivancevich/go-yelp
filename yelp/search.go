package yelp

import "net/url"

// SearchOptions contains the available parameters for the Search API.
type SearchOptions struct {
	Term        *string
	Location    *string
	Coordinates *Coordinates
	Radius      *int64
	Categories  *string
	Locale      *string
	Limit       *int64
	Offset      *int64
	SortBy      *string
	Price       *string
	OpenNow     *bool
	OpenAt      *int64
	Attributes  *string
}

// SearchResults reflects the JSON returned by the Search API.
type SearchResults struct {
	Total      int64      `json:"total"`
	Businesses []Business `json:"businesses"`
	Region     Region     `json:"region"`
}

// IsValid returns true when either Location or Coordinates is set and OpenNow
// and OpenAt are not both set.
func (so SearchOptions) IsValid() bool {
	return ((so.Location != nil) != (so.Coordinates != nil)) &&
		!(so.OpenNow != nil && so.OpenAt != nil)
}

// URLValues returns SearchOptions as url.Values.
func (so SearchOptions) URLValues() url.Values {
	vals := url.Values{}
	if so.Coordinates != nil {
		vals = so.Coordinates.URLValues()
	} else if so.Location != nil {
		vals.Add("location", *so.Location)
	}

	if so.Term != nil {
		vals.Add("term", *so.Term)
	}
	if so.Radius != nil {
		vals.Add("radius", IntString(*so.Radius))
	}
	if so.Categories != nil {
		vals.Add("categories", *so.Categories)
	}
	if so.Locale != nil {
		vals.Add("locale", *so.Locale)
	}
	if so.Limit != nil {
		vals.Add("limit", IntString(*so.Limit))
	}
	if so.Offset != nil {
		vals.Add("offset", IntString(*so.Offset))
	}
	if so.SortBy != nil {
		vals.Add("sort_by", *so.SortBy)
	}
	if so.Price != nil {
		vals.Add("price", *so.Price)
	}
	if so.OpenNow != nil {
		vals.Add("open_now", BoolString(*so.OpenNow))
	} else if so.OpenAt != nil {
		vals.Add("open_at", IntString(*so.OpenAt))
	}
	if so.Attributes != nil {
		vals.Add("attributes", *so.Attributes)
	}
	return vals
}
