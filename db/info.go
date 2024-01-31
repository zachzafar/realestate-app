package db

type Option struct {
	Id   int
	Name string
}

type InfoStore struct {
	Amenities      []Option
	Property_types []Option
	Countries      []Option
	Features       []Option
}

func (d *Database) InitStore() (*InfoStore, error) {
	countries, err := d.GetAllCountries()

	if err != nil {
		return nil, err
	}

	property_types, err := d.GetAllPropertyTypes()

	if err != nil {
		return nil, err
	}

	amenities, err := d.GetAllAmenities()

	if err != nil {
		return nil, err
	}

	features, err := d.GetAllFeatures()

	if err != nil {
		return nil, err
	}

	return &InfoStore{
		Amenities:      amenities,
		Property_types: property_types,
		Countries:      countries,
		Features:       features}, nil

}

func (d *Database) GetAllCountries() ([]Option, error) {
	var countries []Option

	query := `SELECT country_id,name FROM countries`

	rows, err := d.db.Query(query)

	if err != nil {

		return countries, err
	}

	for rows.Next() {
		var country Option

		err := rows.Scan(&country.Id, &country.Name)

		if err != nil {

			return countries, err
		}
		countries = append(countries, country)
	}
	return countries, nil
}

func (d *Database) GetParishes(countryID int) ([]Option, error) {
	var parishes []Option
	var parish Option = Option{Id: 0, Name: "Parish"}
	parishes = append(parishes, parish)

	query := `SELECT parish_id,name FROM parishes WHERE country_id=$1`

	rows, err := d.db.Query(query, countryID)

	if err != nil {

		return parishes, err
	}

	for rows.Next() {
		var parish Option

		err = rows.Scan(&parish.Id, &parish.Name)

		if err != nil {
			return parishes, err
		}
		parishes = append(parishes, parish)
	}

	return parishes, nil
}

func (d *Database) GetAllPropertyTypes() ([]Option, error) {
	var property_types []Option

	query := `SELECT property_type_id,name FROM property_types`

	rows, err := d.db.Query(query)

	if err != nil {
		return property_types, err
	}

	for rows.Next() {
		var property_type Option
		err = rows.Scan(&property_type.Id, &property_type.Name)

		if err != nil {
			return property_types, err
		}

		property_types = append(property_types, property_type)
	}
	return property_types, nil
}

func (d *Database) GetAllAmenities() ([]Option, error) {
	var amenities []Option

	query := `SELECT amenity_id,name FROM amenities`

	rows, err := d.db.Query(query)

	if err != nil {
		return amenities, err
	}

	for rows.Next() {
		var amenity Option

		err = rows.Scan(&amenity.Id, &amenity.Name)

		if err != nil {
			return amenities, err
		}

		amenities = append(amenities, amenity)
	}

	return amenities, nil
}

func (d *Database) GetAllFeatures() ([]Option, error) {
	var features []Option

	query := `SELECT feature_id,name FROM features`

	rows, err := d.db.Query(query)

	if err != nil {
		return features, nil
	}

	for rows.Next() {
		var feature Option

		err := rows.Scan(&feature.Id, &feature.Name)

		if err != nil {
			return features, err
		}

		features = append(features, feature)
	}

	return features, nil
}
