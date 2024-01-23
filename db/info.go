package db

type Option struct {
	Id   int
	Name string
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
