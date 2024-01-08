package db

import (
	"openlettings.com/types"
)

func (d *Database) CreateProperty(property *types.Property) error {
	query := `INSERT INTO properties (title,description,property_type,address,city,size,bedrooms,bathrooms,year_built,price,user_id) VALUES ($1, $2, $3, $4, $5, $6,$7,$8, $9,$10,$11)`
	_, err := d.db.Exec(query, property.Title, property.Description, property.PropertyType, property.Address, property.City, property.Size, property.Bedrooms, property.Bathrooms, property.YearBuilt, property.Price, property.UserID)

	return err
}

func (d *Database) GetProperties(filter *types.PropertyFilter, page int, pageSize int) ([]types.PropertySummary, error) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	query := `SELECT property_id,title,description,price,address FROM properties`
	filterString, parameters := filter.GenerateQueryString()
	parameters = append(parameters, pageSize, (pageSize * page))
	if filterString != "" {
		query = query + ` WHERE ` + filterString + " LIMIT ? OFFSET ?"
	}

	rows, err := d.db.Query(query, parameters...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []types.PropertySummary

	for rows.Next() {
		var property types.PropertySummary

		err := rows.Scan(&property.PropertyId, &property.Title, &property.Description, &property.Price, &property.Address)

		if err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}

	return properties, nil
}

func (d *Database) GetPropertyDetails(propertyId int) (*types.Property, error) {
	query := `SELECT title,description,property_type,address,city,size,bedrooms,bathrooms,year_built,price,user_id FROM properties WHERE property_id=?`
	var property types.Property

	err := d.db.QueryRow(query, propertyId).Scan(&property.Title, &property.Description, &property.PropertyType, &property.Address, &property.City, &property.Size, &property.Bedrooms, &property.Bathrooms, &property.YearBuilt, &property.Price, &property.UserID)

	if err != nil {
		return nil, err
	}

	return &property, nil
}
