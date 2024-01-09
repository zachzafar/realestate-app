package types

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
