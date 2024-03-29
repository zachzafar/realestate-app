package types

import (
	"fmt"
	"net/http"
	"strconv"
)

type PropertyFilter struct {
	PriceRange    Range `db:"price"`
	Property_type int   `db:"property_type_id"`
	Country       int   `db:"country_id"`
	Parish        int   `db:"parish_id"`
	Beds          int   `db:"bedrooms"`
	Bathrooms     int   `db:"bathrooms"`
	UserId        int   `db:"user_id"`
}

func (f PropertyFilter) GetRelationForFilterName() string {
	return "properties"
}

func NewPropertyFilter(property_type int, country int, parish int, priceRange Range, beds int, bathrooms int, userId int) *PropertyFilter {

	return &PropertyFilter{
		Beds:          beds,
		Bathrooms:     bathrooms,
		Property_type: property_type,
		Country:       country,
		Parish:        parish,
		PriceRange:    priceRange,
		UserId:        userId,
	}
}

type PropertySummary struct {
	PropertyId  int
	Title       string
	Description string
	Price       float64
	Country     string
	Parish      string
	Address     string
	Url         string
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

type Property struct {
	Title          string  `db:"title"`
	Description    string  `db:"description"`
	PropertyTypeID int     `db:"property_type_id"`
	Address        string  `db:"address"`
	CountryID      int     `db:"country_id"`
	ParishID       int     `db:"parish_id"`
	Neighborhood   string  `db:"neighborhood"`
	Size           int     `db:"size"`
	Bedrooms       int     `db:"bedrooms"`
	Bathrooms      int     `db:"bathrooms"`
	YearBuilt      int     `db:"year_built"`
	FlooringType   string  `db:"flooring_type"`
	Price          float64 `db:"price"`
	Currency       string  `db:"currency"`
	PaymentTerms   string  `db:"payment_terms"`
	ContactName    string  `db:"contact_name"`
	ContactEmail   string  `db:"contact_email"`
	ContactPhone   string  `db:"contact_phone"`
	Availability   bool    `db:"availability"`
	UserID         int     `db:"user_id"`
}

func (p Property) GetRelationName() string {
	return "properties"
}

func (p Property) GetPrimaryKeyName() string {
	return "property_id"
}

func ParseListingParams(r *http.Request) (*PropertyFilter, int) {
	userId := 0

	if id, err := strconv.Atoi(r.URL.Query().Get("user-id")); err == nil {
		userId = id
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize == 0 {
		pageSize = 10
	}
	propertyType, _ := strconv.Atoi(r.URL.Query().Get("property_type"))
	country, _ := strconv.Atoi(r.URL.Query().Get("country"))
	parish, _ := strconv.Atoi(r.URL.Query().Get("parish"))
	beds, _ := strconv.Atoi(r.URL.Query().Get("bedrooms"))
	bathrooms, _ := strconv.Atoi(r.URL.Query().Get("bathrooms"))
	lowerPriceRange, _ := strconv.Atoi(r.URL.Query().Get("maxprice"))
	upperPriceRange, _ := strconv.Atoi(r.URL.Query().Get("minprice"))

	priceRange := NewRange(upperPriceRange, lowerPriceRange)

	propertyFilter := NewPropertyFilter(propertyType, country, parish, *priceRange, beds, bathrooms, userId)

	return propertyFilter, page
}

func ParsePropertyBody(r *http.Request) (*Property, error) {
	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		return nil, err
	}

	sessionData, ok := r.Context().Value("user-id").(*SessionData)
	if !ok {
		return nil, fmt.Errorf("No user")
	}

	countryID, _ := strconv.Atoi(r.PostFormValue("country"))
	parishID, _ := strconv.Atoi(r.PostFormValue("parish"))
	property_type, _ := strconv.Atoi(r.PostFormValue("type"))
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	address := r.PostFormValue("address")
	size, _ := strconv.Atoi(r.PostFormValue("size"))
	bedrooms, _ := strconv.Atoi(r.PostFormValue("bedrooms"))
	bathrooms, _ := strconv.Atoi(r.PostFormValue("bathrooms"))
	year, _ := strconv.Atoi(r.PostFormValue("year"))
	availability, err := strconv.ParseBool(r.PostFormValue("availability"))
	if err != nil {
		availability = false
	}

	property := &Property{
		Title:          r.PostFormValue("title"),
		Description:    r.PostFormValue("description"),
		PropertyTypeID: property_type,
		Address:        address,
		CountryID:      countryID,
		ParishID:       parishID,
		YearBuilt:      year,
		Bathrooms:      bathrooms,
		Size:           size,
		Bedrooms:       bedrooms,
		UserID:         sessionData.UserId,
		Price:          price,

		Neighborhood: r.PostFormValue("neighborhood"),
		FlooringType: r.PostFormValue("flooring_type"),
		Currency:     r.PostFormValue("currency"),
		PaymentTerms: r.PostFormValue("payment_terms"),
		ContactName:  r.PostFormValue("contact_name"),
		ContactEmail: r.PostFormValue("contact_email"),
		ContactPhone: r.PostFormValue("contact_phone"),
		Availability: availability,
	}

	return property, nil
}
