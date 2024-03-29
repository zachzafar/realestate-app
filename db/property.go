package db

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"openlettings.com/types"
	"openlettings.com/utils"
)

func (d *Database) CreateProperty(property *types.Property) (int, error) {
	query, values := types.GeneratInsertQuery(*property)
	query = query + " RETURNING property_id"
	var id int
	err := d.db.QueryRow(query, values...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *Database) UpdateProperty(property *types.Property, propertyId int) error {
	query, values := types.GenerateUpdateQuery(property)
	values = append(values, propertyId)
	_, err := d.db.Exec(query, values...)
	fmt.Println(query)
	fmt.Println(values...)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteProperty(propertyId int) error {
	query := `DELETE FROM properties WHERE property_id=$1`
	_, err := d.db.Exec(query, propertyId)

	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UploadPropertyPhotos(files []*multipart.FileHeader, id int) error {
	query := `INSERT INTO media (property_id,url) VALUES ($1,$2)`
	id_string := strconv.FormatInt(int64(id), 10)

	if err := os.MkdirAll("./media/properties/"+id_string, 0777); err != nil {
		return utils.CustomError{Message: "failed to create directory"}
	}

	for _, file := range files {

		uploadedFile, err := file.Open()
		if err != nil {
			return utils.CustomError{Message: "problem opening file open"}
		}
		defer uploadedFile.Close()
		f, err := os.OpenFile("./media/properties/"+id_string+"/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {

			return utils.CustomError{Message: "problem saving file"}
		}
		defer f.Close()

		_, err = io.Copy(f, uploadedFile)
		if err != nil {
			return utils.CustomError{Message: "problem copying file"}
		}

		_, err = d.db.Exec(query, id, id_string+"/"+file.Filename)
		if err != nil {
			return err
		}

	}
	return nil
}

func (d *Database) GetImages(id string) ([]string, error) {
	query := `SELECT url FROM media where property_id=$1`
	rows, err := d.db.Query(query, id)

	if err != nil {
		return nil, err
	}
	urls := make([]string, 0)

	for rows.Next() {
		var image string
		rows.Scan(&image)
		urls = append(urls, image)
	}

	return urls, nil
}

func (d *Database) GetProperties(filter *types.PropertyFilter, page int, pageSize int) ([]types.PropertySummary, error) {

	query := `
		SELECT p.property_id,p.title,p.description,p.price,p.address, COALESCE(m.url, '') AS url 
		FROM properties p
		LEFT JOIN media m ON p.property_id = m.property_id
		`
	filterString, parameters := types.GenerateFilterQueryString(*filter)
	paramLength := len(parameters)
	parameters = append(parameters, pageSize, (pageSize * (page - 1)))

	limit := fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramLength+1, paramLength+2)

	if filterString != "" {
		query = query + ` WHERE ` + filterString + limit
	} else {
		query = query + " ORDER BY p.property_id" + limit
	}

	rows, err := d.db.Query(query, parameters...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []types.PropertySummary

	for rows.Next() {
		var property types.PropertySummary

		err := rows.Scan(&property.PropertyId, &property.Title, &property.Description, &property.Price, &property.Address, &property.Url)

		if err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}

	return properties, nil
}

func (d *Database) GetPropertyCount(queryFilter *types.PropertyFilter) (int, error) {

	filterString, params := types.GenerateFilterQueryString(*queryFilter)
	query := `SELECT COUNT(*) AS row_count FROM properties`
	if filterString != "" {
		query = query + " WHERE " + filterString
	}

	row, err := d.db.Query(query, params...)
	if err != nil {

	}
	var count int
	if row.Next() {
		err = row.Scan(&count)
	}

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d *Database) GetPropertyDetails(propertyId int) (*types.Property, error) {
	property := &types.Property{}

	query, pointers := types.GenerateQueryByIDString(property)

	err := d.db.QueryRow(query, propertyId).Scan(pointers...)

	if err != nil {
		return nil, err
	}

	return property, nil
}
