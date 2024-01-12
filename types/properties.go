package types

import (
	"net/http"
	"strconv"
)

// type Property struct {
// 	PropertyID   int      `json:"property_id"`
// 	Title        string   `json:"title" validate:"required"`
// 	Description  *string  `json:"description"`
// 	PropertyType string   `json:"property_type" validate:"required"`
// 	Address      string   `json:"address" validate:"required"`
// 	City         string   `json:"city" validate:"required"`
// 	State        string   `json:"state" validate:"required"`
// 	ZipCode      string   `json:"zip_code" validate:"required"`
// 	Neighborhood *string  `json:"neighborhood"`
// 	Size         *int     `json:"size"`
// 	Bedrooms     *int     `json:"bedrooms"`
// 	Bathrooms    *int     `json:"bathrooms"`
// 	YearBuilt    *int     `json:"year_built"`
// 	FlooringType *string  `json:"flooring_type"`
// 	Appliances   []string `json:"appliances"`
// 	Price        *float64 `json:"price"`
// 	Currency     *string  `json:"currency"`
// 	PaymentTerms *string  `json:"payment_terms"`
// 	Amenities    []string `json:"amenities"`
// 	Features     []string `json:"features"`
// 	Availability bool     `json:"availability"`
// 	Media        []string `json:"media"`
// 	UserID       int      `json:"user_id"`
// }

type Property struct {
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Bedrooms     int
	Bathrooms    int
	Size         int
	Address      string `json:"address"`
	City         string `json:"city"`
	PropertyType string `json:"property_type"`
	YearBuilt    string
	UserID       int
}

type PropertySummary struct {
	PropertyId  int     `json:"property_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Address     string  `json:"address"`
	Url         string  `json:"url"`
}

func NewProperty(title string, desc string, PropType string, address string, city string, year string, bathrooms int, size int, bedrooms int, userid int, price float64) *Property {
	return &Property{
		Title:        title,
		Description:  desc,
		PropertyType: PropType,
		Price:        price,
		Bedrooms:     bedrooms,
		Bathrooms:    bathrooms,
		City:         city,
		UserID:       userid,
		Address:      address,
		YearBuilt:    year,
		Size:         size,
	}
}

func NewPropertySummary(propertyId int, title string, desc string, price float64, address string, url string) *PropertySummary {
	return &PropertySummary{
		PropertyId:  propertyId,
		Title:       title,
		Description: desc,
		Price:       price,
		Address:     address,
		Url:         url,
	}
}

func ParseListingParams(r *http.Request) (*PropertyFilter, int) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize == 0 {
		pageSize = 10
	}
	propertyType := r.URL.Query().Get("type")
	address := r.URL.Query().Get("address")
	upperBedCount, _ := strconv.Atoi(r.URL.Query().Get("upperBedCount"))
	lowerBedcount, _ := strconv.Atoi(r.URL.Query().Get("lowerBedCount"))
	lowerPriceRange, _ := strconv.Atoi(r.URL.Query().Get("lowerPriceRange"))
	upperPriceRange, _ := strconv.Atoi(r.URL.Query().Get("upperPriceRange"))

	priceRange := NewRange(upperPriceRange, lowerPriceRange)
	bedRange := NewRange(upperBedCount, lowerBedcount)

	propertyFilter := NewPropertyFilter(propertyType, address, *priceRange, *bedRange, 0)

	return propertyFilter, page
}

func ParsePropertyBody(r *http.Request) (*Property, error) {
	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		return nil, err
	}

	userId := r.Context().Value("user-id").(int)
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	property_type := r.PostFormValue("type")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	address := r.PostFormValue("address")
	size, _ := strconv.Atoi(r.PostFormValue("size"))
	bedrooms, _ := strconv.Atoi(r.PostFormValue("bedrooms"))
	bathrooms, _ := strconv.Atoi(r.PostFormValue("bathrooms"))
	year := r.PostFormValue("year")
	city := r.PostFormValue("city")

	property := NewProperty(title, description, property_type, address, city, year, bathrooms, size, bedrooms, userId, price)

	return property, nil
}
